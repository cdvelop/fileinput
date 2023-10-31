package fileinput

import (
	"github.com/cdvelop/model"
)

type File struct {
	//table
	Id_file     string `Legend:"File Id" Input:"InputPK"`
	Module_name string `Legend:"Modulo" Input:"TextNumCode"`
	Field_name  string `Legend:"Carpeta Campo" Input:"TextNum" PrincipalField:"ok"`
	Folder_id   string `Legend:"Carpeta Registro" Input:"InputPK"`
	File_area   string `Legend:"Area archivo" Input:"Text" PrincipalField:"ok" SkipCompletionAllowed:"true"`
	Extension   string `Legend:"Tipo Archivo" Input:"Text" SkipCompletionAllowed:"true"`
	Description string `Legend:"Descripción" Input:"Text" PrincipalField:"ok" SkipCompletionAllowed:"true"`
	File_data   string `Legend:"Ubicación" Input:"Text" SkipCompletionAllowed:"true"`

	Files string `Input:"Text" NotRequiredInDB:"true"` //nombre por defecto input html = files

	Object *model.Object

	db model.DataBaseAdapter

	input_id *model.Input

	App *model.Handlers

	model.FileConfig

	source *model.Object // objeto origen

}
