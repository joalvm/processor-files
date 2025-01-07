package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joalvm/processor-medias/pkg/database"
	"github.com/joalvm/processor-medias/pkg/utils"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var orderScripts = []string{
	"scheme.sql",
	"enums/",
	"types/",
	"tables.sql",
	"indices/",
	"functions/",
	"views/",
	"materialized_views/",
	"procedures/",
	"triggers/",
}

// G migrates the database.
func GenerateDb(cmd *cobra.Command, args []string) error {
	dbName := dbName(cmd)

	err := createDatabase(dbName)
	if err != nil {
		return err
	}

	db, err := database.NewWithDbName(dbName)
	if err != nil {
		return err
	}
	defer database.CloseConnection(db)

	err = handleScripts(db)
	if err != nil {
		return err
	}

	// Si es una carpeta
	return nil
}

func createDatabase(dbName string) error {
	db, err := database.NewWithPostgresDb()
	if err != nil {
		return err
	}
	defer database.CloseConnection(db)

	err = handleCurrentDatabase(db, dbName)
	if err != nil {
		return err
	}

	tx := db.Exec(
		fmt.Sprintf(`
			CREATE DATABASE %s
			WITH OWNER=postgres
			ENCODING='UTF8'
			TABLESPACE=default
			CONNECTION LIMIT=-1`,
			dbName,
		),
	)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Si existe una base de datos con el nombre dbName renombrarla a dbName + "_20240423_050034"
func handleCurrentDatabase(db *gorm.DB, dbName string) error {
	// Verificar si la base de datos existe:
	tx := db.Exec(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbName))
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil
	}

	// Cerrar todas las conexiones a la base de datos
	tx = db.Exec(fmt.Sprintf("SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '%s'", dbName))
	if tx.Error != nil {
		return tx.Error
	}

	timestamp := time.Now().Format("20060102_150405")
	tx = db.Exec(fmt.Sprintf("ALTER DATABASE %s RENAME TO %s", dbName, dbName+"_"+timestamp))
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func dbName(cmd *cobra.Command) string {
	dbName := cmd.Flag("database").Value.String()
	if dbName == "" {
		dbName = utils.Env("DB_NAME")
	}

	return dbName
}

func handleScripts(db *gorm.DB) error {
	scripts := []string{}

	for _, folder := range orderScripts {
		if strings.Contains(folder, ".sql") {
			script := utils.Resolve("database/scripts", folder)

			if _, err := os.Stat(script); os.IsNotExist(err) {
				continue
			}

			scripts = append(scripts, script)

			continue
		}

		scripts = append(scripts, getScripts(folder)...)
	}

	return execScripts(db, scripts)
}

func execScripts(db *gorm.DB, scripts []string) error {
	for _, script := range scripts {
		// Leer el archivo
		file, err := os.ReadFile(script)
		if err != nil {
			return err
		}

		if len(file) == 0 {
			continue
		}

		tx := db.Exec(string(file))
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func getScripts(folder string) []string {
	dir := utils.Resolve("database/scripts", folder)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return []string{}
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	scripts := []string{}

	for _, file := range files {
		if file.IsDir() {
			scripts = append(scripts, getScripts(filepath.Join(folder, file.Name()))...)
			continue
		}

		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}

		scripts = append(scripts, utils.Resolve(dir, file.Name()))
	}

	return scripts
}
