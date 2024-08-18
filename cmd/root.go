package cmd

import (
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/cobra"
)

var (
	client memcache.Client
	hosts  []string
)

func NewRootCmd() *cobra.Command {

	rootCmd := &cobra.Command{
		Use:           "ccmc",
		Short:         "ccmc is a CLI tool for interacting with memcached",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			client = *memcache.New(hosts...)
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	//overwrite the help command so that the hosts command can use the -h flag
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for this command")

	//commands
	pingCmd := newPingCommand(&client)
	pingCmd.PersistentFlags().StringSliceVarP(&hosts, "hosts", "h", []string{}, "Comma seprated list of hosts")
	pingCmd.MarkPersistentFlagRequired("hosts")

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(pingCmd)
	return rootCmd
}
