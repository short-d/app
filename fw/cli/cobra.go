package cli

import (
	"errors"

	"github.com/spf13/cobra"
)

var _ CommandFactory = (*CobraFactory)(nil)
var _ Command = (*CobraCommand)(nil)

type CobraFactory struct{}

func (c CobraFactory) NewCommand(config CommandConfig) Command {
	return CobraCommand{
		cmd: &cobra.Command{
			Use:   config.Usage,
			Short: config.ShortHelpMsg,
			Long:  config.DetailedHelpMsg,
			Run: func(cmd *cobra.Command, args []string) {
				var cmdWrapper Command = CobraCommand{cmd: cmd}
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

func (c CobraCommand) AddSubCommand(subCommand Command) error {
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
