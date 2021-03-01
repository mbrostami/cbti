package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read bigtable.",
	Long: `For development purposes only.
	Usage: cbti list tables 
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hii there %+v", args)
	},
}
