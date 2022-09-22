package cmd

import (
	"log"

	"github.com/Kintuda/cassandra-schema-migration/pkg/migrator"
	"github.com/spf13/cobra"
)

var contactPoints string
var keySpace string
var table string
var bundlePath string

func NewMigrationCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migration CLI",
	}

	checkSchema := &cobra.Command{
		Use:  "check",
		RunE: CheckSchema,
	}

	rootCmd.PersistentFlags().StringVarP(&contactPoints, "contact-points", "c", "", "contact points")
	rootCmd.PersistentFlags().StringVarP(&keySpace, "keyspace", "k", "", "keyspace")
	rootCmd.PersistentFlags().StringVarP(&table, "table", "t", "", "table")
	rootCmd.PersistentFlags().StringVarP(&bundlePath, "bundle-path", "b", "", "bundle path")

	rootCmd.AddCommand(checkSchema)

	return rootCmd
}

func CheckSchema(cmd *cobra.Command, args []string) error {
	c, err := migrator.CreateConnection(contactPoints, &bundlePath)

	if err != nil {
		return err
	}

	m := migrator.NewMigrator(c)

	defer c.Close()

	r := m.CheckSchema(keySpace, table)

	log.Print(r)

	return nil
}
