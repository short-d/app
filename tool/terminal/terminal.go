package terminal

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/byliuyang/eventbus"
)

const (
	blackForeground = 0x3330 + iota
	redForeground
	greenForeground
	yellowForeground
	blueForeground
	magentaForeground
	cyanForeground
	whiteForeground
)

const (
	blackBackground = 0x3430 + iota
	redBackground
	greenBackground
	yellowBackground
	blueBackground
	magentaBackground
	cyanBackground
	whiteBackground
)

type Key struct {
	escapeSequence string
	name           string
}

const esc = 0x1b
const carriageReturn = 0x0d

var (
	cursorUpKey       = string([]byte{esc, 0x5b, 0x41})
	cursorDownKey     = string([]byte{esc, 0x5b, 0x42})
	cursorForwardKey  = string([]byte{esc, 0x5b, 0x43})
	cursorBackwardKey = string([]byte{esc, 0x5b, 0x44})
	CtrlEKey          = string([]byte{0x05})
	enterKey          = string([]byte{0x0a})
)

const (
	CursorUpName       = "cursorUp"
	CursorDownName     = "cursorDown"
	CursorForwardName  = "cursorForward"
	CursorBackwardName = "cursorDownward"
	CtrlEName          = "Ctrl+E"
	EnterName          = "enter"
)

var keyNames = map[string]string{
	cursorUpKey:       CursorUpName,
	cursorDownKey:     CursorDownName,
	cursorForwardKey:  CursorForwardName,
	cursorBackwardKey: CursorBackwardName,
	CtrlEKey:          CtrlEName,
	enterKey:          EnterName,
}

// Terminal manipulates terminals that implements VT100 standard:
// http://ascii-table.com/ansi-escape-sequences.php
type Terminal struct {
	out             io.Writer
	in              io.Reader
	foregroundColor int
	backgroundColor int
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

func (t *Terminal) SetForegroundColor(color int) {
	t.foregroundColor = color
}

func (t *Terminal) SetBackgroundColor(color int) {
	t.backgroundColor = color
}

func (t Terminal) Print(text string) {
	_, err := fmt.Fprint(t.out, text)
	if err != nil {
		panic(err)
	}
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

func (t Terminal) UpdateGraphicsMode() {
	t.turnOffTextAttributes()
	t.escape(append(
		[]byte{},
		byte(t.foregroundColor),
		byte(t.backgroundColor),
		0x6d,
	))
}

func (t Terminal) ClearLine() {
	t.escape([]byte{0x32, 0x4b, carriageReturn})
}

func (t Terminal) NewLine() {
	t.Print("\n")
}

func (t Terminal) MoveCursorUp(numLines int) {
	t.escape([]byte{byte(numLines), 0x41})
}

func (t Terminal) MoveCursorDown(numLines int) {
	t.escape([]byte{byte(numLines), 0x42})
}

func (t Terminal) MoveCursorForward(numLines int) {
	t.escape([]byte{byte(numLines), 0x43})
}

func (t Terminal) MoveCursorBackward(numLines int) {
	t.escape([]byte{byte(numLines), 0x44})
}

func (t Terminal) SaveCursorPosition() {
	t.escape([]byte{0x73})
}

func (t Terminal) RestoreCursorPosition() {
	t.escape([]byte{0x75})
}

func (t Terminal) OnKeyPress(keyName string, ch eventbus.DataChannel) {
	t.eventBus.Subscribe(keyName, ch)
}

func (t Terminal) turnOffTextAttributes() {
	t.escape([]byte{0x30, 0x6d})
}

func (t Terminal) escape(sequence []byte) {
	_, err := t.out.Write(append([]byte{esc, 0x5b}, sequence...))
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
