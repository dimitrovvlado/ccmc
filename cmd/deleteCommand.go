package cmd

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/cobra"
)

func newDeleteCommand(c *memcache.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Deletes a cache entry for the provided key",
		RunE: func(cmd *cobra.Command, args []string) error {
			var key string
			var err error
			if key, err = cmd.Flags().GetString("key"); err != nil {
				return err
			}
			if err = c.Delete(key); err != nil {
				return err
			}
			cmd.OutOrStdout().Write([]byte("OK"))
			return nil
		},
	}
}
