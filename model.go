package fileinput

import (
	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/model"
)

type FileInput struct {
	Files string `Input:"Text" NotRequiredInDB:"true"` //nombre por defecto input html = files

	Object *model.Object

	App *model.Handlers

	*filehandler.FileHandler

	*filehandler.FileSetting
}
