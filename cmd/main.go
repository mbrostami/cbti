package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cbti",
	Short: "Bigtable reader",
	Long:  `cbti executable`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Run() {
	rootCmd.AddCommand(readCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}