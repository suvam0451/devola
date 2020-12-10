/**/
package cmd

import (
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `Tools to work with remote servers.`,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
