package cmd

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/cobra"
)

func newPingCommand(c *memcache.Client) *cobra.Command {
	pingCmd := &cobra.Command{
		Use:   "ping",
		Short: "Pings all server instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := c.Ping(); err != nil {
				return err
			}
			cmd.OutOrStdout().Write([]byte("OK"))
			return nil
		},
	}
	return pingCmd
}
