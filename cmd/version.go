package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	APP_VERSION				= "0.1"
)


func init() {

} // init


var versionCmd = &cobra.Command{
	Use: "version",
	Short: "version",
	Long: "cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("prompt hub cli %s", APP_VERSION)
	},
}
