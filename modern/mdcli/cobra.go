package mdcli

import (
	"errors"

	"github.com/short-d/app/fw"
	"github.com/spf13/cobra"
)

var _ fw.CommandFactory = (*CobraFactory)(nil)
var _ fw.Command = (*CobraCommand)(nil)

type CobraFactory struct{}

func (c CobraFactory) NewCommand(config fw.CommandConfig) fw.Command {
	return CobraCommand{
		cmd: &cobra.Command{
			Use:   config.Usage,
			Short: config.ShortHelpMsg,
			Long:  config.DetailedHelpMsg,
			Run: func(cmd *cobra.Command, args []string) {
				var cmdWrapper fw.Command = CobraCommand{cmd: cmd}
				config.OnExecute(&cmdWrapper, args)
			},
		},
	}
}

func NewCobraFactory() CobraFactory {
	return CobraFactory{}
}

type CobraCommand struct {
	cmd *cobra.Command
}

func (c CobraCommand) Execute() error {
	return c.cmd.Execute()
}

func (c CobraCommand) AddSubCommand(subCommand fw.Command) error {
	subCobraCmd, ok := subCommand.(CobraCommand)
	if !ok {
		return errors.New("fail casting fw.Command to CobraCommand")
	}
	c.cmd.AddCommand(subCobraCmd.cmd)
	return nil
}

func (c CobraCommand) AddStringFlag(valueHolder *string, name string, defaultValue string, shortDescription string) {
	c.cmd.Flags().StringVar(valueHolder, name, defaultValue, shortDescription)
}
