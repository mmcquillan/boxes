package boxes

import (
	"os"
	"sort"
	"strings"

	"github.com/nsf/termbox-go"
)

// TODO:
// - bind keys

type Window struct {
	Height          int
	Width           int
	OutputMode      OutputMode
	BackgroundColor termbox.Attribute
	BorderWidth     int
	BorderColor     termbox.Attribute
	Boxes           map[string]Box
	BoxContent      map[string]string
	BoxCursor       map[string]int
}

type (
	OutputMode int
)

const (
	OutputCurrent OutputMode = iota
	OutputNormal
	Output256
	Output216
	OutputGrayscale
)

func NewWindow() (window Window) {

	// initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	// set defaults
	window.OutputMode = Output256
	window.BackgroundColor = termbox.ColorDefault
	window.BorderWidth = 0
	window.BorderColor = 242
	window.Boxes = make(map[string]Box)
	window.BoxContent = make(map[string]string)

	// bind input

	return window
}

func (window Window) Draw() {

	// interpret window type
	var outputMode termbox.OutputMode
	switch window.OutputMode {
	case OutputCurrent:
		outputMode = termbox.OutputCurrent
	case OutputNormal:
		outputMode = termbox.OutputNormal
	case Output256:
		outputMode = termbox.Output256
	case Output216:
		outputMode = termbox.Output216
	case OutputGrayscale:
		outputMode = termbox.OutputGrayscale
	}

	// set window values
	termbox.SetOutputMode(outputMode)
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, window.BackgroundColor)

	// reset window
	w, h := termbox.Size()
	window.Width = w
	window.Height = h

	// window border
	if window.BorderWidth > 0 {
		for y := 0; y < window.Height; y++ {
			for x := 0; x < window.Width; x++ {
				if y < window.BorderWidth || y >= window.Height-window.BorderWidth {
					termbox.SetCell(x, y, ' ', coldef, window.BorderColor)
				} else if x < window.BorderWidth || x >= window.Width-window.BorderWidth {
					termbox.SetCell(x, y, ' ', coldef, window.BorderColor)
				}
			}
		}
	}

	// organize boxes
	boxMap := make(map[int]string)
	var boxSort []int
	for key, box := range window.Boxes {
		if keys, chk := boxMap[box.Z]; chk {
			boxMap[box.Z] = keys + "," + key
		} else {
			boxMap[box.Z] = key
			boxSort = append(boxSort, box.Z)
		}
	}
	sort.Ints(boxSort)

	// draw boxes in z order
	for _, s := range boxSort {
		for _, key := range strings.Split(boxMap[s], ",") {
			box := window.Boxes[key]
			plot := Plotter(window, box)
			runes := []rune(window.BoxContent[key])
			runeMarker := 0
			runeLine := 0
			for y := plot.Y; y < (plot.Y + plot.H); y++ {
				for x := plot.X; x < (plot.X + plot.W); x++ {

					// box border
					if y < plot.Y+box.BorderWidth || y >= plot.Y+plot.H-box.BorderWidth {
						termbox.SetCell(x, y, ' ', coldef, box.BorderColor)
					} else if x < plot.X+box.BorderWidth || x >= plot.X+plot.W-box.BorderWidth {
						termbox.SetCell(x, y, ' ', coldef, box.BorderColor)
					} else {

						// box padding
						if y < plot.Y+box.BorderWidth+box.PaddingTop || y >= plot.Y+plot.H-box.BorderWidth-box.PaddingBottom {
							termbox.SetCell(x, y, ' ', box.ForegroundColor, box.BackgroundColor)
						} else if x < plot.X+box.BorderWidth+box.PaddingLeft || x >= plot.X+plot.W-box.BorderWidth-box.PaddingRight {
							termbox.SetCell(x, y, ' ', box.ForegroundColor, box.BackgroundColor)
						} else {

							// write content
							if len(runes) > runeMarker {
								if runes[runeMarker] == '\n' {
									runeMarker++
									runeLine++
								}
								if runeLine <= y {
									termbox.SetCell(x, y, runes[runeMarker], box.ForegroundColor, box.BackgroundColor)
									runeMarker++
									runeLine = y
								} else {
									termbox.SetCell(x, y, ' ', box.ForegroundColor, box.BackgroundColor)
								}
							} else {
								termbox.SetCell(x, y, ' ', box.ForegroundColor, box.BackgroundColor)
							}
						}

					}

					// input
					if box.Input {
						termbox.SetCursor(plot.X+box.BorderWidth+box.PaddingLeft, plot.Y+box.BorderWidth+box.PaddingTop)
					}

				}
			}
		}
	}

	// flush the screen to draw
	termbox.Flush()
}

func (window Window) BindKeys() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				window.Close()
				os.Exit(0)
			case termbox.KeyArrowLeft, termbox.KeyCtrlB:
				//edit_box.MoveCursorOneRuneBackward()
			case termbox.KeyArrowRight, termbox.KeyCtrlF:
				//edit_box.MoveCursorOneRuneForward()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				window.deleteContent()
				//edit_box.DeleteRuneBackward()
			case termbox.KeyDelete, termbox.KeyCtrlD:
				//edit_box.DeleteRuneForward()
			case termbox.KeyEnter:
				window.appendContent('\n')
			case termbox.KeyTab:
				window.appendContent('\t')
			case termbox.KeySpace:
				window.appendContent(' ')
			case termbox.KeyHome, termbox.KeyCtrlA:
				//edit_box.MoveCursorToBeginningOfTheLine()
			case termbox.KeyEnd, termbox.KeyCtrlE:
				//edit_box.MoveCursorToEndOfTheLine()
			default:
				if ev.Ch != 0 {
					window.appendContent(ev.Ch)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func (window Window) Close() {
	termbox.Close()
}

func (window Window) Update(name string, box Box) {
	name = strings.Replace(name, ",", "_", -1)
	window.Boxes[name] = box
	window.BoxContent[name] = ""
}

func (window Window) Content(name string, content string) {
	name = strings.Replace(name, ",", "_", -1)
	window.BoxContent[name] = content
}

func (window Window) appendContent(r rune) {
	for key, _ := range window.Boxes {
		if window.Boxes[key].Input {
			window.BoxContent[key] = window.BoxContent[key] + string(r)
			window.Draw()
		}
	}
}

func (window Window) deleteContent() {
	for key, _ := range window.Boxes {
		if window.Boxes[key].Input {
			window.BoxContent[key] = window.BoxContent[key][0 : len(window.BoxContent[key])-1]
			window.Draw()
		}
	}
}
