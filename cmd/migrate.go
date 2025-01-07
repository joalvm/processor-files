package cmd

import (
	"github.com/joalvm/processor-medias/internal"
	"github.com/joalvm/processor-medias/pkg/utils"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "generate-db",
	Short: "Migrar la base de datos.",
	Long:  `El proposito de este comando es migrar la base de datos.`,
	RunE:  internal.GenerateDb,
}

var defaultDbName string = utils.Env("DB_NAME")

func init() {
	migrateCmd.Flags().StringP("database", "d", defaultDbName, "Nombre de la base de datos.")
	rootCmd.AddCommand(migrateCmd)
}
