package cli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/byliuyang/app/tool/terminal"
)

type CommandLineTool struct {
	term terminal.Terminal
}

func (c *CommandLineTool) EnterMainLoop(handleEvents func()) {
	// Prevent process from being killed by pressing Ctrl + C
	signal.Ignore(syscall.SIGINT)
	c.term.StopWaitForEnter()
	c.term.StartEventLoop()
	c.term.TurnOffEcho()

	for {
		handleEvents()
	}
}

func (c *CommandLineTool) Exit() {
	c.term.TurnOnEcho()
	os.Exit(0)
}

func NewCommandLineTool(term terminal.Terminal) CommandLineTool {
	return CommandLineTool{
		term: term,
	}
}
