package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "migrator",
		Short: "Cassandra migrator cli",
	}

	rootCmd.AddCommand(NewMigrationCmd())

	return rootCmd
}
