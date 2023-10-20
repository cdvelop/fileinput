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
	Description string `Legend:"Descripción" Input:"Text" PrincipalField:"ok" SkipCompletionAllowed:"true"`
	File_path   string `Legend:"Ubicación" Input:"FilePath" SkipCompletionAllowed:"true"`

	//config
	Files string `Input:"Text" NotRequiredInDB:"true"` //nombre por defecto input html = files

	Object *model.Object
	db     model.DataBaseAdapter

	filetype    string //imagen, video, document
	root_folder string //ej: "./app_files"

	model.FileConfig
}
