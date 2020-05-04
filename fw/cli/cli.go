package cli

type CommandConfig struct {
	Usage           string
	ShortHelpMsg    string
	DetailedHelpMsg string
	OnExecute       func(cmd *Command, args []string)
}

type Command interface {
	Execute() error
	AddSubCommand(subCommand Command) error
	AddStringFlag(valueHolder *string, name string, defaultValue string, shortDescription string)
}

type CommandFactory interface {
	NewCommand(config CommandConfig) Command
}
