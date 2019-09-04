package terminal

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/byliuyang/eventbus"
)

const esc = 27

type ColorCode string
type vt100ColorCodes struct {
	Black   ColorCode
	Red     ColorCode
	Green   ColorCode
	Yellow  ColorCode
	Blue    ColorCode
	Magenta ColorCode
	Cyan    ColorCode
	White   ColorCode
}

var ForegroundColor = vt100ColorCodes{
	Black:   "30",
	Red:     "31",
	Green:   "32",
	Yellow:  "33",
	Blue:    "34",
	Magenta: "35",
	Cyan:    "36",
	White:   "37",
}

var BackgroundColor = vt100ColorCodes{
	Black:   "40",
	Red:     "41",
	Green:   "42",
	Yellow:  "43",
	Blue:    "44",
	Magenta: "45",
	Cyan:    "46",
	White:   "47",
}

type Key struct {
	escapeSequence string
	name           string
}

const (
	CursorUpName       = "cursorUp"
	CursorDownName     = "cursorDown"
	CursorForwardName  = "cursorForward"
	CursorBackwardName = "cursorDownward"
	CtrlEName          = "Ctrl+E"
)

var keyNames = map[string]string{
	"\033[A": CursorUpName,
	"\033[B": CursorDownName,
	"\033[C": CursorForwardName,
	"\033[D": CursorBackwardName,
	"\005":   CtrlEName,
}

// Terminal manipulates terminals that implements VT100 standard:
// http://ascii-table.com/ansi-escape-sequences.php
type Terminal struct {
	out             io.Writer
	in              io.Reader
	foregroundColor ColorCode
	backgroundColor ColorCode
	eventBus        eventbus.EventBus
}

func (t *Terminal) StopWaitForEnter() {
	t.SetMinCharsForRead(0)
}

// SetMinReadLen sets charCount characters minimum for a completed read
func (t *Terminal) SetMinCharsForRead(charCount int) {
	t.execute("stty", "-icanon", "min", fmt.Sprintf("%d", charCount))
}

func (t *Terminal) TurnOffEcho() {
	t.execute("stty", "-echo")
}

func (t *Terminal) TurnOnEcho() {
	t.execute("stty", "echo")
}

func (t Terminal) HideCursor() {
	t.execute("tput", "civis")
}

func (t Terminal) ShowCursor() {
	t.execute("tput", "cnorm")
}

func (t *Terminal) SetForegroundColor(color ColorCode) {
	t.foregroundColor = color
}

func (t *Terminal) SetBackgroundColor(color ColorCode) {
	t.backgroundColor = color
}

func (t Terminal) UpdateGraphicsMode() {
	t.turnOffTextAttributes()
	t.escape(fmt.Sprintf("%s%sm", t.foregroundColor, t.backgroundColor))
}

func (t Terminal) Print(text string) {
	_, err := fmt.Fprint(t.out, text)
	if err != nil {
		panic(err)
	}
}

func (t Terminal) ClearLine() {
	t.escape("2K\r")
}

func (t Terminal) NewLine() {
	t.Print("\n")
}

func (t Terminal) MoveCursorUp(numLines int) {
	t.escape(fmt.Sprintf("%dA", numLines))
}

func (t Terminal) MoveCursorDown(numLines int) {
	t.escape(fmt.Sprintf("%dB", numLines))
}

func (t Terminal) MoveCursorForward(numLines int) {
	t.escape(fmt.Sprintf("%dC", numLines))
}

func (t Terminal) MoveCursorBackward(numLines int) {
	t.escape(fmt.Sprintf("%dD", numLines))
}

func (t Terminal) Read() []byte {
	buf := make([]byte, 1)
	n, err := t.in.Read(buf)
	if err != nil {
		return []byte(nil)
	}
	return buf[:n]
}

func (t Terminal) StartEventLoop() {
	slowDownPeriod := 60 * time.Millisecond
	resetPeriod := 60 * time.Millisecond
	reset := make(chan time.Time)

	go func() {
		var buf []byte

		// Pressing the arrow keys produces escape sequence in the standard input.
		// For example, pressing up arrow key produces \033[A, which contains three bytes: \033, [, and A.
		// Each byte will be read sequentially ( going through the for loop 3 times for up arrow key).
		for {
			select {
			case <-reset:
				buf = nil
			default:
				char := t.Read()
				if len(char) < 1 {
					time.Sleep(slowDownPeriod)
					continue
				}

				buf = append(buf, char...)
				sequence := string(buf)
				keyName, ok := keyNames[sequence]
				if !ok && buf[0] == esc {
					// The chars read before and after resetPeriod do not belong to a single key press.
					// Clear the bytes of the previous keypress after resetPeriod.
					go func() {
						reset <- <-time.After(resetPeriod)
					}()
					continue
				}

				if !ok {
					buf = nil
					time.Sleep(slowDownPeriod)
					continue
				}

				t.eventBus.Publish(keyName, nil)
				buf = nil
			}
		}
	}()
}

func (t Terminal) SaveCursorPosition() {
	t.escape("s")
}

func (t Terminal) RestoreCursorPosition() {
	t.escape("u")
}

func (t Terminal) OnKeyPress(keyName string, ch eventbus.DataChannel) {
	t.eventBus.Subscribe(keyName, ch)
}

func (t Terminal) turnOffTextAttributes() {
	t.escape("0m")
}

func (t Terminal) escape(sequence string) {
	_, err := fmt.Fprintf(t.out, "\033[%s", sequence)
	if err != nil {
		panic(err)
	}
}

func (t Terminal) execute(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = t.in
	cmd.Stdout = t.out

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func NewTerminal() Terminal {
	return Terminal{
		in:       os.Stdin,
		out:      os.Stdout,
		eventBus: eventbus.NewEventBus(),
	}
}
