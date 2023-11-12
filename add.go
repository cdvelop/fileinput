package fileinput

import (
	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
	"github.com/cdvelop/unixid"
)

// conf:
// field_name:voucher,user_photo,boleta... default file
// filetype:video, pdf, document. default imagen
// root_folder:static_files default "app_files"
// max_files:1, 4, 6.. default 6
// max_kb_size:100, 400 default 50

func NewUploadFileApi(h *model.Handlers, o *model.Object, s filehandler.FileSetting) (*FileInput, error) {

	f := FileInput{}

	input_id := unixid.InputPK()

	// agregar inputs usados en tabla file al modulo
	err := o.Module.AddInputs([]*model.Input{input_id, f.Input(), FileType(), SavedMode(), input.Text(), input.TextNumCode(), input.TextNum(), input.FilePath(), input.TextOnly()}, "file pkg")
	if err != nil {
		return nil, err
	}

	// crear objeto file
	err = object.New(&f, o.Module, h)
	if err != nil {
		return nil, err
	}

	fileHandler, err := filehandler.Add(h.FileRootFolder, h, h)
	if err != nil {
		return nil, err
	}
	f.FileHandler = fileHandler

	// configuración por defecto file
	new_file := filehandler.FileSetting{
		MaximumFilesAllowed: 6,
		InputNameWithFiles:  f.Files,
		MaximumFileSize:     0,
		MaximumKbSize:       50,
		AllowedExtensions:   ",.jpg, .png, .jpeg",
		ImagenWidth:         "800",
		ImagenHeight:        "600",

		FileType: "i",

		FieldNameWithObjectID: o.PrimaryKeyName(),
		Name:                  s.Name,
		Legend:                "Imágenes",

		Source: o, // asignamos el objeto origen
	}

	if new_file.Name == "" {
		return nil, model.Error(`fileinput error FileSetting.Name:"nombre_campo" no ingresado`)
	}

	new_file.DefaultEnableInput = s.DefaultEnableInput

	if s.ImagenWidth != "" {
		new_file.ImagenWidth = s.ImagenWidth
	}

	if s.ImagenHeight != "" {
		new_file.ImagenHeight = s.ImagenHeight
	}

	if s.FileType != "" {
		new_file.FileType = s.FileType

		switch s.FileType {
		case "video":
			new_file.Legend = "Videos"
			new_file.AllowedExtensions = ".avi, .mkv, .mp4"
			new_file.FileType = "v"
		case "document":
			new_file.Legend = "Documentos"
			new_file.AllowedExtensions = ".doc, .xlsx, .txt"
			new_file.FileType = "d"
		case "pdf":
			new_file.Legend = "Documentos PDF"
			new_file.AllowedExtensions = ".pdf"
			new_file.FileType = "p"
		}
	}

	if s.MaximumFilesAllowed != new_file.MaximumFilesAllowed {
		new_file.MaximumFilesAllowed = s.MaximumFilesAllowed
	}

	if s.MaximumKbSize != new_file.MaximumKbSize {
		new_file.MaximumKbSize = s.MaximumKbSize
	}

	f.Object.Name += "." + new_file.Name

	new_file.MaximumFileSize = int64(float64(new_file.MaximumFilesAllowed*new_file.MaximumKbSize*1024) * 1.05)

	//nota: al no declarar punteros se pierden posteriormente

	f.Object.ViewHandler = f

	// agrego el campo input file al objeto para mostrarlo en la vista
	o.Fields = append(o.Fields, model.Field{
		Name:                  f.Files,
		Legend:                f.Legend,
		SourceTable:           f.Object.Table,
		Input:                 f.Input(),
		SkipCompletionAllowed: true,
		NotRequiredInDB:       true,
	})

	f.FileHandler.AddNewFileSetting(f.Object, f.FileSetting)

	return &f, nil

}

func (f FileInput) ViewHandlerName() string {
	return "FileInput"
}
