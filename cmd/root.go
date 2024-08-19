package cmd

import (
	"os"
	"time"

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

	//start commands

	//ping command
	pingCmd := newPingCommand(&client)
	pingCmd.PersistentFlags().StringSliceVarP(&hosts, "hosts", "h", []string{}, "Comma seprated list of hosts")
	pingCmd.MarkPersistentFlagRequired("hosts")

	//set command
	setCmd := newSetCommand(&client)
	setCmd.PersistentFlags().StringSliceVarP(&hosts, "hosts", "h", []string{}, "Comma seprated list of hosts")
	setCmd.PersistentFlags().StringP("key", "k", "", "Key of the cached entity")
	setCmd.PersistentFlags().StringP("val", "v", "", "String value of the cached entity")
	setCmd.PersistentFlags().StringP("file", "f", "", "File path value for the cached entity")
	setCmd.PersistentFlags().DurationP("expiration", "e", time.Duration(0), `The cache expiration time, sample values are 300ms, 2m, 1h, 2h45m etc. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h", "d", "w"`)
	setCmd.MarkPersistentFlagRequired("hosts")
	setCmd.MarkPersistentFlagRequired("key")
	setCmd.MarkFlagsOneRequired("val", "file")
	setCmd.MarkFlagsMutuallyExclusive("val", "file")

	//get command
	getCmd := newGetCommand(&client)
	getCmd.PersistentFlags().StringSliceVarP(&hosts, "hosts", "h", []string{}, "Comma seprated list of hosts")
	getCmd.PersistentFlags().StringP("key", "k", "", "Key of the cached entity")
	getCmd.MarkPersistentFlagRequired("key")

	//delete command
	deleteCmd := newDeleteCommand(&client)
	deleteCmd.PersistentFlags().StringSliceVarP(&hosts, "hosts", "h", []string{}, "Comma seprated list of hosts")
	deleteCmd.PersistentFlags().StringP("key", "k", "", "Key of the cached entity")
	deleteCmd.MarkPersistentFlagRequired("key")

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(deleteCmd)
	//end commands

	return rootCmd
}
