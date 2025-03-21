package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/duanwei33/learn-golang/example/cobra-command/operator"
)

// NewOperatorCommand creates and returns the root command
func NewOperatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "my-operator",
		Short: "My Custom Operator",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	// Create and add the "start" subcommand
	ctrlCmd := &cobra.Command{
		Use:   "start",
		Short: "Start My Operator",
		Run: func(cmd *cobra.Command, args []string) {
			operator.RunOperator()
		},
	}

	// Add the "start" subcommand to the root command
	cmd.AddCommand(ctrlCmd)

	return cmd
}

func main() {
	if err := NewOperatorCommand().Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

