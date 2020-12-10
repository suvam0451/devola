/* */
package cmd

import (
	"github.com/suvam0451/devola/collection"
	"fmt"

	"github.com/spf13/cobra"
)

var ip string
var id string

// digitalOceanCmd represents the digitalOcean command
var digitalOceanCmd = &cobra.Command{
	Use:   "digitalOcean",
	Short: "Tools to manage/operate with DigitalOcean droplets",
	Long: `
keyring: Used to store SSH keys and IP sets
`,
	// Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("digitalOcean called")
		collection.AddServer(ip, "default")
	},
}

func init() {
	rootCmd.AddCommand(digitalOceanCmd)
	digitalOceanCmd.PersistentFlags().Int8("SSH Type", 0, "Key pair type being used (NOTE: Fedora 33 might require xyz for proper SSH handshake")
	digitalOceanCmd.Flags().StringVarP(&ip, "IP address", "a", "127.0.0.1", "IP address for the droplet entry")
	digitalOceanCmd.Flags().StringVarP(&id, "ID tag", "i", "default", "Name to use for key entry")
}
