package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newTaskCmd() *cobra.Command {
	commitCmd := &cobra.Command{
		Use:   "task",
		Short: "used to commit with ticket ID and #comment flag",
		Long:  `used to commit with ticket ID and #comment flag`,
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Her in prerun")
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Here inside run")
		},
	}

	//commitCmd.Flags().BoolVar(&git.Options.Header, "header", false, "Remove header")
	//commitCmd.Flags().BoolVarP(&git.Options.Comment, "comment", "c", false, "Remove comment")
	//commitCmd.Flags().BoolVarP(&git.Options.Sign, "sign", "s", false, "Sign")

	return commitCmd
}
