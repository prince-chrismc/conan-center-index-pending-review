package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"syscall"

	"github.com/google/go-github/github"
	"github.com/google/go-querystring/query"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

// Model
type Package struct {
	FullName        string
	Description     string
	StarsCount      int
	ForksCount      int
	LastUpdatedBy   string
	OpenIssuesCount int
}

type PullRequest struct {
	Number    int
	Comments  int
	ReviewUrl string
}

type Review struct {
	ID                *int         `json:"id,omitempty"`
	NodeID            *string      `json:"node_id,omitempty"`
	User              *github.User `json:"user,omitempty"`
	State             *string      `json:"state,omitempty"`
	AuthorAssociation *string      `json:"author_association,omitempty"`
	CommitID          *string      `json:"commit_id,omitempty"`
}

func (r *Review) GetState() string {
	if r == nil || r.State == nil {
		return ""
	}
	return *r.State
}
func (r *Review) GetAuthorAssociation() string {
	if r == nil || r.AuthorAssociation == nil {
		return ""
	}
	return *r.AuthorAssociation
}
func (r *Review) GetCommitID() string {
	if r == nil || r.CommitID == nil {
		return ""
	}
	return *r.CommitID
}

type ReviewsResponse []*Review

var timestampType = reflect.TypeOf(github.Timestamp{})

func Stringify(message interface{}) string {
	var buf bytes.Buffer
	v := reflect.ValueOf(message)
	stringifyValue(&buf, v)
	return buf.String()
}

func stringifyValue(w io.Writer, val reflect.Value) {
	if val.Kind() == reflect.Ptr && val.IsNil() {
		w.Write([]byte("<nil>"))
		return
	}

	v := reflect.Indirect(val)

	switch v.Kind() {
	case reflect.String:
		fmt.Fprintf(w, `"%s"`, v)
	case reflect.Slice:
		w.Write([]byte{'['})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte{' '})
			}

			stringifyValue(w, v.Index(i))
		}

		w.Write([]byte{']'})
		return
	case reflect.Struct:
		if v.Type().Name() != "" {
			w.Write([]byte(v.Type().String()))
		}

		// special handling of Timestamp values
		if v.Type() == timestampType {
			fmt.Fprintf(w, "{%s}", v.Interface())
			return
		}

		w.Write([]byte{'{'})

		var sep bool
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			if fv.Kind() == reflect.Ptr && fv.IsNil() {
				continue
			}
			if fv.Kind() == reflect.Slice && fv.IsNil() {
				continue
			}

			if sep {
				w.Write([]byte(", "))
			} else {
				sep = true
			}

			w.Write([]byte(v.Type().Field(i).Name))
			w.Write([]byte{':'})
			stringifyValue(w, fv)
		}

		w.Write([]byte{'}'})
	default:
		if v.CanInterface() {
			fmt.Fprint(w, v.Interface())
		}
	}
}

func (u ReviewsResponse) String() string {
	return Stringify(u)
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func main() {
	context := context.Background()

	var httpClient *http.Client

	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if exists {
		tokenService := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient = oauth2.NewClient(context, tokenService)
	} else {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("GitHub Username: ")
		username, _ := r.ReadString('\n')

		fmt.Print("GitHub Password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		password := string(bytePassword)

		tp := github.BasicAuthTransport{
			Username: strings.TrimSpace(username),
			Password: strings.TrimSpace(password),
		}

		httpClient = tp.Client()
	}

	client := github.NewClient(httpClient)

	// Get Rate limit information
	rateLimit, _, err := client.RateLimits(context)
	if err != nil {
		fmt.Printf("Problem in getting rate limit information %v\n", err)
		return
	}

	fmt.Printf("Limit: %d \nRemaining %d \n", rateLimit.Core.Limit, rateLimit.Core.Remaining)

	repo, _, err := client.Repositories.Get(context, "conan-io", "conan-center-index")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &Package{
		FullName:        *repo.FullName,
		Description:     *repo.Description,
		ForksCount:      *repo.ForksCount,
		StarsCount:      *repo.StargazersCount,
		OpenIssuesCount: *repo.OpenIssuesCount,
	}

	fmt.Printf("%+v\n", pack)

	pulls, _, err := client.PullRequests.List(context, "conan-io", "conan-center-index", &github.PullRequestListOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 3,
		},
	})
	for _, pr := range pulls {
		p := PullRequest{Number: pr.GetNumber(), Comments: pr.GetComments(), ReviewUrl: pr.GetReviewCommentsURL()}
		fmt.Printf("%+v\n", p)

		u, err := addOptions(p.ReviewUrl, github.ListOptions{
			Page:    0,
			PerPage: 10,
		})
		if err != nil {
			fmt.Printf("Problem format reviews request url %v\n", err)
			return
		}

		req, err := client.NewRequest("GET", u, nil)
		if err != nil {
			fmt.Printf("Problem making reviews request %v\n", err)
			return
		}

		reviews := new(ReviewsResponse)
		// resp, err := client.Do(context, req, reviews)
		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Printf("Problem executing reviews request %v\n", err)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		ioutil.WriteFile("dump.json", body, 0600)
		buff := string(body)
		fmt.Println(buff)

		json.Unmarshal([]byte(buff), reviews)

		for _, review := range *reviews {
			// r := *review
			// fmt.Printf("%+v\n", r)
			fmt.Printf("%s (%s): '%s' on commit %s\n", review.User.GetLogin(), review.GetAuthorAssociation(), review.GetState(), review.GetCommitID())
		}
	}
}

