package cmd

import (

	"github.com/spf13/cobra"
)


var (

	rootCmd = &cobra.Command{
		Use: "static",
		Short: "pmt hub cli",
		Long: "prompt hub cli tool",
	}

)


func Execute() error {

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(versionCmd)

	return rootCmd.Execute()

} // Execute


func init() {

	cobra.OnInitialize()

} // init
