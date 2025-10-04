package cmd

import (

	"github.com/spf13/cobra"
)


var pullCmd = &cobra.Command{
	Use: "pull",
	Short: "get prompt",
	Long: "get prompt from prompt hub",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
