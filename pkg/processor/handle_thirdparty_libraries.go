package processor

import (
	"github.com/joalvm/processor-medias/pkg/ffmpeg"
	"github.com/joalvm/processor-medias/pkg/imagemagick"
)

func (p *Processor) HandleThirdPartyLibraries() error {
	err := ffmpeg.Install(p.ffmpegUrl)
	if err != nil {
		return err
	}

	err = imagemagick.Install(p.imagemagickUrl)
	if err != nil {
		return err
	}

	return nil
}
