package cmd

import "github.com/spf13/cobra"

type Command interface {
	getCLI() *cobra.Command
}
