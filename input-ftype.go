package fileinput

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

type fileType struct{}

func (fileType) SourceData() map[string]string {
	return map[string]string{"i": "Imagen", "v": "Video", "p": "pdf", "w": "Word", "e": "Excel"}
}

func FileType() *model.Input {
	return input.Radio("FileType", fileType{})
}
