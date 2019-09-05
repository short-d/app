package ui

import (
	"fmt"

	"github.com/byliuyang/app/tool/terminal"
)

type Radio struct {
	term        terminal.Terminal
	items       []string
	maxRows     int
	itemCount   int
	selectedIdx int
	currentTop  int
	isRendered  bool
}

func (r *Radio) Render() {
	endBefore := r.currentTop + r.maxRows
	for idx := r.currentTop; idx < endBefore; idx++ {
		item := r.items[idx]
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

	defer r.term.ClearLine()
	for idx := 0; idx < r.maxRows-1; idx++ {
		r.term.ClearLine()
		r.term.MoveCursorUp(1)
	}
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
	defer r.Render()
	defer r.Clear()

	r.selectedIdx = maxInt(r.selectedIdx-1, 0)
	if r.selectedIdx >= r.currentTop {
		return
	}
	r.currentTop--
}

func (r *Radio) Next() {
	defer r.Render()
	defer r.Clear()

	r.selectedIdx = minInt(r.selectedIdx+1, r.itemCount-1)
	if r.selectedIdx < r.currentTop+r.maxRows {
		return
	}
	r.currentTop++
}

func (r Radio) Remove() {
	defer r.term.ShowCursor()
	r.Clear()
}

func (r *Radio) SelectedIdx() int {
	return r.selectedIdx
}

func NewRadio(items []string, maxRows int, term terminal.Terminal) Radio {
	radio := Radio{
		term:        term,
		items:       items,
		maxRows:     maxRows,
		itemCount:   len(items),
		selectedIdx: 0,
		currentTop:  0,
		isRendered:  false,
	}
	radio.term.HideCursor()
	return radio
}
