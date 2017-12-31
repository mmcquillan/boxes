package boxes

import (
	"github.com/nsf/termbox-go"
)

type Box struct {
	BackgroundColor termbox.Attribute
	ForegroundColor termbox.Attribute
	BorderWidth     int
	BorderColor     termbox.Attribute
	PaddingTop      int
	PaddingBottom   int
	PaddingLeft     int
	PaddingRight    int
	Input           bool
	Anchor          string
	Size            string
	Z               int
	Align           string
	Wrap            bool
	Hidden          bool
}
