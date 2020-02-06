package ui

import (
	"fmt"
	"os/signal"
	"syscall"

	"github.com/short-d/app/tool/terminal"
	"github.com/short-d/eventbus"
)

var _ Component = (*Radio)(nil)

type onDisappear func()
type onItemSelected func(selectedIdx int)

type Radio struct {
	terminal terminal.Terminal

	items       []string
	maxRows     int
	itemCount   int
	selectedIdx int
	currentTop  int
	isRendered  bool

	exitChannel     eventbus.DataChannel
	keyUpChannel    eventbus.DataChannel
	keyDownChannel  eventbus.DataChannel
	keyEnterChannel eventbus.DataChannel

	onDisappear    onDisappear
	onItemSelected onItemSelected
}

func (r *Radio) Show() {
	r.terminal.HideCursor()
	r.showInstructions()
	r.render()

	signal.Ignore(syscall.SIGINT)
	r.terminal.StopWaitForEnter()
	r.terminal.StartEventLoop()
	r.terminal.TurnOffEcho()

	for r.isRendered {
		r.processEvents()
	}
}

func (r *Radio) OnDisappear(onDisappear onDisappear) {
	r.onDisappear = onDisappear
}

func (r *Radio) OnItemSelected(onItemSelected onItemSelected) {
	r.onItemSelected = onItemSelected
}

func (r *Radio) Hide() {
	defer r.terminal.TurnOnEcho()
	defer r.terminal.ShowCursor()
	defer r.clearInstructions()

	r.onDisappear()
	r.clear()
	r.isRendered = false
}

func (r *Radio) SelectedIdx() int {
	return r.selectedIdx
}

func (r *Radio) showInstructions() {
	fmt.Println("To exit, press Ctrl + E")
	fmt.Println("To select an item, press Enter")
	fmt.Println("")
}

func (r *Radio) clearInstructions() {
	for idx := 0; idx < 3; idx++ {
		r.terminal.MoveCursorUp(1)
		r.terminal.ClearLine()
	}
}

func (r *Radio) processEvents() {
	select {
	case <-r.exitChannel:
		r.Hide()
	case <-r.keyUpChannel:
		r.prev()
	case <-r.keyDownChannel:
		r.next()
	case <-r.keyEnterChannel:
		r.Hide()
		r.onItemSelected(r.selectedIdx)
	}
}

func (r *Radio) bindEvents() {
	r.terminal.OnKeyPress(terminal.CtrlEName, r.exitChannel)
	r.terminal.OnKeyPress(terminal.CursorUpName, r.keyUpChannel)
	r.terminal.OnKeyPress(terminal.CursorDownName, r.keyDownChannel)
	r.terminal.OnKeyPress(terminal.EnterName, r.keyEnterChannel)
}

func (r *Radio) render() {
	endBefore := r.currentTop + r.maxRows
	for idx := r.currentTop; idx < endBefore; idx++ {
		item := r.items[idx]
		if idx == r.selectedIdx {
			r.terminal.Print(fmt.Sprintf(" %s %s", FilledSquare, item))
		} else {
			r.terminal.Print(fmt.Sprintf("   %s", item))
		}
		r.terminal.NewLine()
	}
	r.terminal.MoveCursorUp(1)
	r.isRendered = true
}

func (r Radio) clear() {
	if !r.isRendered {
		return
	}

	defer r.terminal.ClearLine()
	for idx := 0; idx < r.maxRows-1; idx++ {
		r.terminal.ClearLine()
		r.terminal.MoveCursorUp(1)
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

func (r *Radio) prev() {
	defer r.render()
	defer r.clear()

	r.selectedIdx = maxInt(r.selectedIdx-1, 0)
	if r.selectedIdx >= r.currentTop {
		return
	}
	r.currentTop--
}

func (r *Radio) next() {
	defer r.render()
	defer r.clear()

	r.selectedIdx = minInt(r.selectedIdx+1, r.itemCount-1)
	if r.selectedIdx < r.currentTop+r.maxRows {
		return
	}
	r.currentTop++
}

func NewRadio(items []string, maxRows int, terminal terminal.Terminal) Radio {
	if maxRows > len(items) {
		maxRows = len(items)
	}

	radio := Radio{
		terminal: terminal,

		items:       items,
		maxRows:     maxRows,
		itemCount:   len(items),
		selectedIdx: 0,
		currentTop:  0,
		isRendered:  false,

		exitChannel:     make(eventbus.DataChannel),
		keyUpChannel:    make(eventbus.DataChannel),
		keyDownChannel:  make(eventbus.DataChannel),
		keyEnterChannel: make(eventbus.DataChannel),

		onDisappear:    func() {},
		onItemSelected: func(selectedIdx int) {},
	}
	radio.bindEvents()
	return radio
}
