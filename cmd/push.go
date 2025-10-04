package cmd

import (

	"github.com/spf13/cobra"
)


var pushCmd = &cobra.Command{
	Use: "push",
	Short: "push prompt",
	Long: "push prompt to prompt hub",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
