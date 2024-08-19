package cmd

import (
	"fmt"

	"github.com/dimitrovvlado/ccmc/version"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ccmc",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(fmt.Sprintf("ccmc version %s\n", version.VERSION)))
		},
	}
}
