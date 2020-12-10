/**/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/suvam0451/devola/collection"
)

// serverListCmd represents the serverList command
var serverListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `List all server entries`,
	Run: func(cmd *cobra.Command, args []string) {
		collection.ListServers()
	},
}

func init() {
	serverCmd.AddCommand(serverListCmd)
}
