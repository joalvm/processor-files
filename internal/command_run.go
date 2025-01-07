package internal

import (
	"fmt"

	"github.com/joalvm/processor-medias/pkg/database"
	"github.com/joalvm/processor-medias/pkg/processor"
	"github.com/joalvm/processor-medias/pkg/utils"
	"github.com/spf13/cobra"
)

// Run is the command to run the processor-medias
func Run(cmd *cobra.Command, args []string) {
	db, err := database.New()
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	processor := processor.New(
		processor.WithSourceDir(utils.Env("SOURCE_DIR")),
		processor.WithDestinationDir(utils.Env("DESTINATION_DIR")),
		processor.WithFfmpegUrl("https://www.gyan.dev/ffmpeg/builds/ffmpeg-release-full.7z"),
		processor.WithImagemagickUrl("https://imagemagick.org/archive/binaries/ImageMagick-7.1.1-31-portable-Q16-HDRI-x64.zip"),
		processor.WithDb(db),
	)

	processor.Process()
}
