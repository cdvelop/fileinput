package fileinput

import (
	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/model"
)

type FileInput struct {
	// Files string `Input:"Text" NotRequiredInDB:"true"` //nombre por defecto input html = files

	Object *model.Object

	notifyDelete model.NotifyFileDelete

	model.Logger

	// *filehandler.FileHandler

	conf *filehandler.FileSetting

	dom   model.DomAdapter
	theme model.ThemeAdapter

	input_enable model.CallJsOptions

	input_reset model.CallJsOptions

	gamepad model.CallJsOptions
}
