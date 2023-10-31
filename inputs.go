package fileinput

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

type fileType struct{}

func FileType() *model.Input {
	return input.Radio("FileType", fileType{})
}

func (fileType) SourceData() map[string]string {
	return map[string]string{"i": "Imagen", "v": "Video", "p": "pdf", "w": "Word", "e": "Excel"}
}

type savedMode struct{}

func SavedMode() *model.Input {
	return input.Radio("SavedMode", savedMode{})
}

func (savedMode) SourceData() map[string]string {
	return map[string]string{"p": "en disco hdd (Path)", "b": "en db (Blob)"}
}
