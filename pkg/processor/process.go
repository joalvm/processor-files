package processor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
	"github.com/joalvm/processor-medias/pkg/enums"
	"github.com/joalvm/processor-medias/pkg/images"
	"github.com/joalvm/processor-medias/pkg/models"
	"github.com/joalvm/processor-medias/pkg/utils"
	"github.com/joalvm/processor-medias/pkg/videos"
)

func (p *Processor) Process() error {
	err := p.HandleThirdPartyLibraries()

	if err != nil {
		return err
	}

	files, err := SearchFiles(p.sourceDir)
	if err != nil {
		return err
	}

	bar := utils.ProgressBar(len(files))
	bar.Start()

	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			return err
		}

		if info.IsDir() {
			p.makeDirectory(filePath)
			bar.Increment()

			p.directoryService.Create(models.Directory{
				Code:     utils.StrRandom(10),
				RealName: filepath.Base(filePath),
				Index:    p.GlobalIndex,
				DirIndex: p.FilesIndexer[filePath[len(p.sourceDir):]],
			})

			continue
		}

		p.GlobalIndex++
		bar.Increment()

		continue

		bar.Set("filename", formatFilePath(filePath[len(p.sourceDir):]))

		directory := filepath.Dir(filePath)[len(p.sourceDir):]

		model, err := p.handleFile(file, p.FilesIndexer[directory], p.destinationDir, directory)
		if err != nil {
			return err
		}

		p.DirectoryMap[directory] = append(p.DirectoryMap[directory], &model)

		p.FilesIndexer[directory]++
		p.GlobalIndex++

		bar.Increment()
	}

	bar.Set("filename", "")
	bar.Finish()

	return nil
}

func (p *Processor) handleFile(file *os.File, dirIndex int, destinationDir string, directory string) (models.Media, error) {
	info, err := file.Stat()
	if err != nil {
		return models.Media{}, err
	}

	model := models.Media{
		Code:     utils.StrRandom(10),
		RealName: filepath.Base(info.Name()),
		Size:     info.Size(),
		DirIndex: dirIndex,
		Index:    p.GlobalIndex,
	}

	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)

	file.Read(head)

	kind, err := filetype.Match(head)
	if err != nil {
		return model, err
	}

	if kind == filetype.Unknown {
		return model, err
	}

	model.MimeType = kind.MIME.Value

	if kind.MIME.Type == "image" {
		err = p.handleImageFile(file, &model, destinationDir, directory)
	} else if kind.MIME.Type == "video" {
		err = p.handleVideoFile(file, &model, destinationDir, directory)
	}
	if err != nil {
		return model, err
	}

	return model, nil
}

func (p *Processor) handleImageFile(file *os.File, model *models.Media, destinationDir string, directory string) error {
	thumbSizes := struct{ Md, Sm, Xs int }{Xs: 100, Sm: 300, Md: 500}

	img := images.New(
		images.WithFile(file),
		images.WithModel(model),
		images.WithDestinationDir(destinationDir),
		images.WithThumbSizes(thumbSizes),
		images.WithDirectory(directory),
		images.WithDb(p.db),
	)

	err := img.Process()
	if err != nil {
		return err
	}

	return nil
}

func (p *Processor) handleVideoFile(file *os.File, model *models.Media, destinationDir string, directory string) error {
	formats := []enums.FormatExt{enums.MP4, enums.WEBM}
	thumbSizes := struct{ Md, Sm, Xs int }{Xs: 100, Sm: 300, Md: 500}

	video := videos.New(
		videos.WithFile(file),
		videos.WithModel(model),
		videos.WithDestinationDir(destinationDir),
		videos.WithFormats(formats),
		videos.WithThumbSizes(thumbSizes),
		videos.WithDirectory(directory),
		videos.WithDb(p.db),
	)

	err := video.Process()
	if err != nil {
		return err
	}

	return nil
}

func formatFilePath(filePath string) string {
	if len(filePath) > 50 {
		return "..." + filePath[len(filePath)-47:]
	}

	// Si tiene menos de 50 caracteres, rellenamos con espacios
	spaces := 50 - len(filePath)
	return filePath + fmt.Sprintf("%*s", spaces, " ")
}

func (p *Processor) makeDirectory(path string) {
	directory := path[len(p.sourceDir):]

	if _, exists := p.DirectoryMap[directory]; !exists {
		p.DirectoryMap[directory] = []*models.Media{}
	}
}
