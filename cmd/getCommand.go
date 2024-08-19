package cmd

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/cobra"
)

func newGetCommand(c *memcache.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Gets a value to a new or existing key",
		RunE: func(cmd *cobra.Command, args []string) error {
			var key string
			var err error
			var item *memcache.Item
			if key, err = cmd.Flags().GetString("key"); err != nil {
				return err
			}
			if item, err = c.Get(key); err != nil {
				return nil
			}
			cmd.OutOrStdout().Write(item.Value)
			return nil
		},
	}
}
