package cmd

import (
	"math"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/cobra"
)

func newSetCommand(c *memcache.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "Sets a new value to a new or existing key",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			var key string
			var val string
			var dur time.Duration
			if key, err = cmd.Flags().GetString("key"); err != nil {
				return err
			}
			if val, err = cmd.Flags().GetString("val"); err != nil {
				return err
			}
			if dur, err = cmd.Flags().GetDuration("expiration"); err != nil {
				return err
			}
			exp := int32(math.Floor(dur.Seconds()))
			if err := c.Set(&memcache.Item{Key: key, Value: []byte(val), Expiration: exp}); err != nil {
				return err
			}
			cmd.OutOrStdout().Write([]byte("OK"))
			return nil
		},
	}
}
