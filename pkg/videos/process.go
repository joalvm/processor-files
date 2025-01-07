package videos

import (
	"os"

	"github.com/joalvm/processor-medias/pkg/ffmpeg"
	"github.com/joalvm/processor-medias/pkg/utils"
)

func (v *Video) Process() error {
	ffmpeg := ffmpeg.NewFfmpeg().Input(v.file.Name()).Output(v.getFilename())

	err := v.makeFolder()
	if err != nil {
		return err
	}

	width, height := ffmpeg.Probe.Dimesions()
	v.model.Width = width
	v.model.Height = height
	v.model.Fps = ffmpeg.Probe.Fps()
	v.model.Duration = ffmpeg.Probe.Duration()

	// v.makePreview(ffmpeg)
	// // panic("implement me")
	// v.makeScreenshot(ffmpeg)

	return nil
}

func (v *Video) makePreview(ffmpeg *ffmpeg.Ffmpeg) error {
	_, err := ffmpeg.Preview()
	if err != nil {
		return err
	}

	return nil
}

func (v *Video) makeScreenshot(ffmpeg *ffmpeg.Ffmpeg) error {
	_, err := ffmpeg.Screenshot()
	if err != nil {
		return err
	}

	return nil
}

func (v *Video) makeFolder() error {
	err := os.MkdirAll(v.getFolder(), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (v *Video) getFolder() string {
	return utils.Resolve(v.destinationDir, v.directory, v.model.Code)
}

func (v *Video) getFilename() string {
	return utils.Resolve(v.getFolder(), v.model.Code+".mp4")
}
