package cmd

import (
	"github.com/joalvm/processor-medias/internal"
	"github.com/spf13/cobra"
)

var processorCmd = &cobra.Command{
	Use:   "run",
	Short: "Inicia el procesamiento de archivos de video e imagenes.",
	Long:  `El proposito de este comando es iniciar el procesamiento de archivos de video e imagenes para la web creando diferentes versiones de los mismos.`,
	Run:   internal.Run,
}

func init() {
	rootCmd.AddCommand(processorCmd)
}
