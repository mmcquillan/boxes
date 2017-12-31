package boxes

import (
	"testing"
)

func TestPlots(t *testing.T) {

	window := Window{
		Height:      100,
		Width:       100,
		BorderWidth: 0,
	}

	box := Box{
		BorderWidth: 0,
		Anchor:      "",
		Size:        "",
	}

	box.Anchor = "0,0"
	box.Size = "1,1"
	testPlot(window, box, Plot{X: 0, Y: 0, W: 1, H: 1}, t)

	box.Anchor = "2,4"
	box.Size = "5,3"
	testPlot(window, box, Plot{X: 2, Y: 4, W: 5, H: 3}, t)

	box.Anchor = "middle"
	box.Size = "10,10"
	testPlot(window, box, Plot{X: 45, Y: 45, W: 10, H: 10}, t)

	box.Anchor = "left"
	box.Size = "10,10"
	testPlot(window, box, Plot{X: 0, Y: 45, W: 10, H: 10}, t)

	box.Anchor = "right"
	box.Size = "10,10"
	testPlot(window, box, Plot{X: 90, Y: 45, W: 10, H: 10}, t)

	box.Anchor = "top"
	box.Size = "10,10"
	testPlot(window, box, Plot{X: 45, Y: 0, W: 10, H: 10}, t)

	box.Anchor = "bottom"
	box.Size = "10,10"
	testPlot(window, box, Plot{X: 45, Y: 90, W: 10, H: 10}, t)

	box.Anchor = "middle"
	box.Size = "9,9"
	testPlot(window, box, Plot{X: 46, Y: 46, W: 9, H: 9}, t)

	box.Anchor = "middle"
	box.Size = "*,10"
	testPlot(window, box, Plot{X: 0, Y: 45, W: 100, H: 10}, t)

	box.Anchor = "middle"
	box.Size = "10,*"
	testPlot(window, box, Plot{X: 45, Y: 0, W: 10, H: 100}, t)

	box.Anchor = "10,10"
	box.Size = "*,10"
	testPlot(window, box, Plot{X: 10, Y: 10, W: 90, H: 10}, t)

	box.Anchor = "left"
	box.Size = "20%,100%"
	testPlot(window, box, Plot{X: 0, Y: 0, W: 20, H: 100}, t)

	window.BorderWidth = 5
	box.Anchor = "left"
	box.Size = "20%,100%"
	testPlot(window, box, Plot{X: 5, Y: 5, W: 18, H: 90}, t)

}

func testPlot(window Window, box Box, plot Plot, t *testing.T) {
	p := Plotter(window, box)
	if p.X != plot.X {
		t.Fatalf("Plot Test Failed - Want: %+v , Have: %+v for Window: %+v , Box: %+v", plot, p, window, box)
	}
	if p.Y != plot.Y {
		t.Fatalf("Plot Test Failed - Want: %+v , Have: %+v for Window: %+v , Box: %+v", plot, p, window, box)
	}
	if p.H != plot.H {
		t.Fatalf("Plot Test Failed - Want: %+v , Have: %+v for Window: %+v , Box: %+v", plot, p, window, box)
	}
	if p.W != plot.W {
		t.Fatalf("Plot Test Failed - Want: %+v , Have: %+v for Window: %+v , Box: %+v", plot, p, window, box)
	}
}
