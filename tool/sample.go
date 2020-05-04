package main

import (
	"fmt"
	"os"

	"github.com/short-d/app/fw/cli"
	"github.com/short-d/app/fw/cli/ui"
	"github.com/short-d/app/fw/terminal"

	"github.com/short-d/eventbus"
	"github.com/spf13/cobra"
)

type SampleTool struct {
	exitChannel     eventbus.DataChannel
	keyUpChannel    eventbus.DataChannel
	keyDownChannel  eventbus.DataChannel
	keyEnterChannel eventbus.DataChannel
	gui             cli.GUI
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
	s.gui.OnKeyPress(terminal.CtrlEName, s.exitChannel)
	s.gui.OnKeyPress(terminal.CursorUpName, s.keyUpChannel)
	s.gui.OnKeyPress(terminal.CursorDownName, s.keyDownChannel)
	s.gui.OnKeyPress(terminal.EnterName, s.keyEnterChannel)
	fmt.Println("To exit, press Ctrl + E")
	fmt.Println("To select an item, press Enter")
}

func (s SampleTool) handleEvents() {
	s.gui.EnterMainLoop(func() {
		select {
		case <-s.exitChannel:
			s.radio.Remove()
			fmt.Println("Terminating process...")
			s.gui.Exit()
		case <-s.keyUpChannel:
			s.radio.Prev()
		case <-s.keyDownChannel:
			s.radio.Next()
		case <-s.keyEnterChannel:
			s.radio.Remove()
			selectedItem := s.languages[s.radio.SelectedIdx()]
			fmt.Printf("Selected %s\n", selectedItem)
			s.gui.Exit()
		}
	})
}

func NewSampleTool() SampleTool {
	term := terminal.NewTerminal()
	languages := []string{
		"Go",
		"Rust",
		"C",
		"C++",
		"Java",
		"Python",
		"C#",
		"JavaScript",
		"TypeScript",
		"Swift",
		"Kotlin",
	}

	sampleTool := SampleTool{
		gui:             cli.NewGUI(term),
		exitChannel:     make(eventbus.DataChannel),
		keyUpChannel:    make(eventbus.DataChannel),
		keyDownChannel:  make(eventbus.DataChannel),
		keyEnterChannel: make(eventbus.DataChannel),
		radio:           ui.NewRadio(languages, 3, term),
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

func main() {
	sampleTool := NewSampleTool()
	sampleTool.Execute()
}
