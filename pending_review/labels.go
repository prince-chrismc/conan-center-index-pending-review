package pending_review

import "fmt"

func processLabels(labels []*Label) error {
	for _, label := range labels {
		switch label.GetName() {
		case "bug", "stale", "Failed", "infrastructure", "blocked", "duplicate", "invalid":
			return fmt.Errorf("%w", ErrStoppedLabel)
		case "Bump version", "Bump dependencies":
			return fmt.Errorf("%w", ErrBumpLabel)
		default:
			continue
		}
	}

	return nil
}
