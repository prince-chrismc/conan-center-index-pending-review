package pending_review

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	defer gock.Off()

	gock.New("http://foo.com").
		Get("/bar").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	res, err := http.Get("http://foo.com/bar")

	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, 200)

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, string(body)[:13], `{"foo":"bar"}`)

	// Verify that we don't have pending mocks
	assert.Equal(t, gock.IsDone(), true)
}
