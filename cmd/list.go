package cmd

import (

	"github.com/spf13/cobra"
)


var listCmd = &cobra.Command{
	Use: "list",
	Short: "show all prompts",
	Long: "show all prompts",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
