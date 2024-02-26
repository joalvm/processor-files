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
			makeDirectory(filePath, p.sourceDir)
			bar.Increment()

			continue
		}

		bar.Set("filename", formatFilePath(filePath[len(p.sourceDir):]))

		directory := filepath.Dir(filePath)[len(p.sourceDir):]

		model, err := handleFile(file, DirectoryIndexes[directory], p.destinationDir, directory)
		if err != nil {
			return err
		}

		DirectoryMap[directory] = append(DirectoryMap[directory], &model)

		DirectoryIndexes[directory]++
		GlobalIndex++

		bar.Increment()
	}

	bar.Set("filename", "")
	bar.Finish()

	return nil
}

func handleFile(file *os.File, dirIndex int, destinationDir string, directory string) (models.File, error) {
	info, err := file.Stat()
	if err != nil {
		return models.File{}, err
	}

	model := models.File{
		Code:     utils.StrRandom(10),
		Name:     utils.StrRandom(10),
		RealName: filepath.Base(info.Name()),
		Size:     info.Size(),
		DirIndex: dirIndex,
		Index:    GlobalIndex,
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
		err := handleImageFile(file, &model, destinationDir, directory)
		if err != nil {
			return model, err
		}

		// panic("stop")
	}

	return model, nil
}

func handleImageFile(file *os.File, model *models.File, destinationDir string, directory string) error {
	formats := []enums.FormatExt{enums.JPEG, enums.WEBP}
	thumbSizes := struct{ Md, Sm, Xs int }{Xs: 100, Sm: 300, Md: 500}

	img := images.New(
		images.WithFile(file),
		images.WithModel(model),
		images.WithDestinationDir(destinationDir),
		images.WithFormats(formats),
		images.WithThumbSizes(thumbSizes),
		images.WithDirectory(directory),
	)

	err := img.Process()
	if err != nil {
		return err
	}

	// jsonData, err := json.MarshalIndent(model, "", "  ")
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(string(jsonData))

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

func makeDirectory(path string, sourceDir string) {
	directory := path[len(sourceDir):]

	if _, exists := DirectoryMap[directory]; !exists {
		DirectoryMap[directory] = []*models.File{}
	}
}
