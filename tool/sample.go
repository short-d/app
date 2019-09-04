package tool

import (
	"fmt"
	"os"

	"github.com/byliuyang/app/tool/cli"
	"github.com/byliuyang/app/tool/terminal"
	"github.com/byliuyang/app/tool/ui"
	"github.com/byliuyang/eventbus"
	"github.com/spf13/cobra"
)

type SampleTool struct {
	term            terminal.Terminal
	exitChannel     eventbus.DataChannel
	keyUpChannel    eventbus.DataChannel
	keyDownChannel  eventbus.DataChannel
	keyEnterChannel eventbus.DataChannel
	cli             cli.CommandLineTool
	rootCmd         *cobra.Command
	radio           ui.Radio
	languages       []string
}

func (s SampleTool) Execute() {
	if err := s.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (s SampleTool) bindKeys() {
	s.term.OnKeyPress(terminal.CtrlEName, s.exitChannel)
	s.term.OnKeyPress(terminal.CursorUpName, s.keyUpChannel)
	s.term.OnKeyPress(terminal.CursorDownName, s.keyDownChannel)
	s.term.OnKeyPress(terminal.EnterName, s.keyEnterChannel)
	fmt.Println("To exit, press Ctrl + E")
	fmt.Println("To select an item, press Enter")
}

func (s SampleTool) handleEvents() {
	s.cli.EnterMainLoop(func() {
		select {
		case <-s.exitChannel:
			s.radio.Remove()
			fmt.Println("Terminating process...")
			s.cli.Exit()
		case <-s.keyUpChannel:
			s.radio.Prev()
		case <-s.keyDownChannel:
			s.radio.Next()
		case <-s.keyEnterChannel:
			s.radio.Remove()
			selectedItem := s.languages[s.radio.SelectedIdx()]
			fmt.Printf("Selected %s\n", selectedItem)
			s.cli.Exit()
		}
	})
}

func NewSampleTool() SampleTool {
	term := terminal.NewTerminal()
	languages := []string{
		"Go",
		"C++",
		"Java",
		"Python",
		"JavaScript",
		"Rust",
	}

	sampleTool := SampleTool{
		term:            term,
		cli:             cli.NewCommandLineTool(term),
		exitChannel:     make(eventbus.DataChannel),
		keyUpChannel:    make(eventbus.DataChannel),
		keyDownChannel:  make(eventbus.DataChannel),
		keyEnterChannel: make(eventbus.DataChannel),
		radio:           ui.NewRadio(languages, term),
		languages:       languages,
	}
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			sampleTool.bindKeys()
			sampleTool.radio.Render()
			sampleTool.handleEvents()
		},
	}
	sampleTool.rootCmd = rootCmd
	return sampleTool
}
