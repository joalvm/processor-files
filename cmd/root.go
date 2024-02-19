package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "processor-medias",
	Short: "Processor de medias para archivos de video e imagenes.",
	Long:  "\nEl proposito de este programa es procesar archivos de video e imagenes para la web creando diferentes versiones de los mismos.",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Do Stuff Here
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
