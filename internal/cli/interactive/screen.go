package interactive

import "github.com/wavesoftware/serverless-installer/internal/domain/model"

type screenData struct {
	next    Screen
	answers *model.Answers
}

// Screen - represents a CLI screen that user sees.
type Screen interface {
	Display()
	Next() Screen
}

func (d screenData) Next() Screen {
	return d.next
}
