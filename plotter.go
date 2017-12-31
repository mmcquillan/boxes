package boxes

import (
	"strconv"
	"strings"
)

func Plotter(window Window, box Box) (plot Plot) {

	// decode size
	if strings.Contains(box.Size, ",") {
		s := strings.Split(box.Size, ",")
		if len(s) == 2 {
			if w, err := strconv.ParseInt(s[0], 10, 32); err == nil {
				plot.W = int(w)
			} else {
				if s[0] == "*" {
					plot.W = window.Width - (2 * window.BorderWidth)
				}
				if strings.Contains(s[0], "%") {
					sp := strings.Replace(s[0], "%", "", 1)
					if w, err := strconv.ParseInt(sp, 10, 32); err == nil {
						plot.W = (window.Width - (2 * window.BorderWidth)) * int(w) / 100
					}
				}
			}
			if h, err := strconv.ParseInt(s[1], 10, 32); err == nil {
				plot.H = int(h)
			} else {
				if s[1] == "*" {
					plot.H = window.Height - (2 * window.BorderWidth)
				}
				if strings.Contains(s[1], "%") {
					sp := strings.Replace(s[1], "%", "", 1)
					if h, err := strconv.ParseInt(sp, 10, 32); err == nil {
						plot.H = (window.Height - (2 * window.BorderWidth)) * int(h) / 100
					}
				}
			}
		}
	}

	// decode anchor
	if strings.Contains(box.Anchor, ",") {
		a := strings.Split(box.Anchor, ",")
		if len(a) == 2 {
			if x, err := strconv.ParseInt(a[0], 10, 32); err == nil {
				plot.X = int(x)
			}
			if y, err := strconv.ParseInt(a[1], 10, 32); err == nil {
				plot.Y = int(y)
			}
		}
	} else {
		if box.Anchor == "middle" {
			plot.X = (window.Width / 2) - (plot.W / 2)
			plot.Y = (window.Height / 2) - (plot.H / 2)
		}
		if box.Anchor == "left" {
			plot.X = 0 + window.BorderWidth
			plot.Y = (window.Height / 2) - (plot.H / 2)
		}
		if box.Anchor == "right" {
			plot.X = window.Width - plot.W - window.BorderWidth
			plot.Y = (window.Height / 2) - (plot.H / 2)
		}
		if box.Anchor == "top" {
			plot.X = (window.Width / 2) - (plot.W / 2)
			plot.Y = 0 + window.BorderWidth
		}
		if box.Anchor == "bottom" {
			plot.X = (window.Width / 2) - (plot.W / 2)
			plot.Y = window.Height - plot.H
		}
	}

	// double check
	if window.Width < (plot.X + plot.W) {
		plot.W = window.Width - plot.X
	}
	if window.Height < (plot.Y + plot.H) {
		plot.H = window.Height - plot.Y
	}

	return plot

}
