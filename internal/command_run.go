package internal

import (
	"github.com/joalvm/processor-medias/pkg/ffmpeg"
	"github.com/joalvm/processor-medias/pkg/imagemagick"
	"github.com/spf13/cobra"
)

// Run is the command to run the processor-medias
func Run(cmd *cobra.Command, args []string) {
	ffmpeg.Install()
	imagemagick.Install()
}
