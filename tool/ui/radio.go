package ui

import (
	"fmt"

	"github.com/byliuyang/app/tool/terminal"
)

type Radio struct {
	term        terminal.Terminal
	items       []string
	itemCount   int
	selectedIdx int
	isRendered  bool
}

func (r *Radio) Render() {
	for idx, item := range r.items {
		if idx == r.selectedIdx {
			r.term.Print(fmt.Sprintf(" %s %s", FilledSquare, item))
		} else {
			r.term.Print(fmt.Sprintf("   %s", item))
		}
		r.term.NewLine()
	}
	r.term.MoveCursorUp(1)
	r.isRendered = true
}

func (r Radio) Clear() {
	if !r.isRendered {
		return
	}

	for idx := 0; idx < r.itemCount-1; idx++ {
		r.term.ClearLine()
		r.term.MoveCursorUp(1)
	}
	r.term.ClearLine()
}

func maxInt(num1 int, num2 int) int {
	if num1 < num2 {
		return num2
	}
	return num1
}

func minInt(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func (r *Radio) Prev() {
	r.selectedIdx = maxInt(r.selectedIdx-1, 0)
	r.Clear()
	r.Render()
}

func (r *Radio) Next() {
	r.selectedIdx = minInt(r.selectedIdx+1, r.itemCount-1)
	r.Clear()
	r.Render()
}

func (r Radio) Remove() {
	r.Clear()
	r.term.ShowCursor()
}

func NewRadio(items []string, term terminal.Terminal) Radio {
	radio := Radio{
		term:        term,
		items:       items,
		itemCount:   len(items),
		selectedIdx: 0,
		isRendered:  false,
	}
	radio.term.HideCursor()
	return radio
}
