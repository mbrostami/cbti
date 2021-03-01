package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mbrostami/cbti/internal/config"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read bigtable.",
	Long: `For development purposes only.
	Usage: cbti list tables 
`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
