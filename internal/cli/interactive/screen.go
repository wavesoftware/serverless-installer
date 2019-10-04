package interactive

import "github.com/wavesoftware/serverless-installer/internal/domain/contract"

type screen struct {
	next    Screen
	answers *contract.LocationAwareAnswers
}

// Screen - represents a CLI screen that user sees.
type Screen interface {
	Display()
	Next() Screen
}

func (s screen) Next() Screen {
	return s.next
}
