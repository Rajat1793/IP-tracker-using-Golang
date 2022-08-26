package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:   "iptracker",
		Short: "IP Tracker",
		Long: `		IP tracking is the technology which allows you to easily search, 
		find, track and trace the location not only of your public IP, but also the 
		location of any other publicly accessible IP or domain in the world.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
