package cmd

import (
	"gencode/cmd/install"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install gf binary to system (might need root/admin permission)",
	Long:  `install gf binary to system (might need root/admin permission)`,
	Run: func(cmd *cobra.Command, args []string) {
		install.Install()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
