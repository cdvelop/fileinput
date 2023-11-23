package fileinput

import (
	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

// conf:
// field_name:voucher,user_photo,boleta... default file
// filetype:video, pdf, document. default imagen
// root_folder:static_files default "app_files"
// MaximumFilesAllowed:1, 4, 6.. default 6
// max_kb_size:100, 400 default 50

func NewUploadFileApi(h *model.Handlers, o *model.Object, s filehandler.FileSetting) (F *FileInput, err string) {

	f := FileInput{}

	// agregar inputs usados en tabla file al modulo
	err = o.Module.AddInputs([]*model.Input{f.Input(), FileType(), SavedMode(), input.FilePath(), input.TextOnly(), input.Text()}, "fileinput pkg")
	if err != "" {
		return nil, err
	}

	// crear objeto file upload
	err = object.AddToHandlerFromStructs(&f, o.Module, h)
	if err != "" {
		return nil, err
	}
	f.Logger = h.Logger
	f.Object.NoAddObjectInDB = true

	f.dom = h
	f.theme = h

	handler, err := filehandler.Add(h)
	if err != "" {
		return nil, err
	}

	// configuración por defecto file
	f.conf = &filehandler.FileSetting{
		AllowedExtensions:     s.AllowedExtensions,
		MaximumFilesAllowed:   6,
		MaximumKbSize:         50,
		ImagenWidth:           "800",
		ImagenHeight:          "600",
		FieldNameWithObjectID: o.PrimaryKeyName(),
		ModuleName:            o.ModuleName,
		DescriptiveName:       s.DescriptiveName,
		Legend:                "Imágenes",
		// source:                o, // asignamos el objeto origen
	}

	if f.conf.DescriptiveName == "" {
		return nil, `fileinput error FileSetting DescriptiveName no ingresado`
	}

	f.conf.DefaultEnableInput = s.DefaultEnableInput

	if s.ImagenWidth != "" {
		f.conf.ImagenWidth = s.ImagenWidth
	}

	if s.ImagenHeight != "" {
		f.conf.ImagenHeight = s.ImagenHeight
	}

	if s.MaximumFilesAllowed != f.conf.MaximumFilesAllowed {
		f.conf.MaximumFilesAllowed = s.MaximumFilesAllowed
	}

	if s.MaximumKbSize != f.conf.MaximumKbSize {
		f.conf.MaximumKbSize = s.MaximumKbSize
	}

	f.Object.ObjectName = o.ModuleName + ".file." + f.conf.DescriptiveName
	// f.Object.ObjectName = o.ObjectName+"files" + f.conf.DescriptiveName

	f.conf.SetMaximumFileSize()

	if len(f.conf.AllowedExtensions) == 0 {
		return nil, "AllowedExtensions error debes ingresar al menos una extension ej: .jpg, .txt"
	}

	//nota: al no declarar punteros se pierden posteriormente

	f.Object.FrontHandler.ViewAdapter = f

	// agrego el campo input file al objeto para mostrarlo en la vista
	o.Fields = append(o.Fields, model.Field{
		Name:                  "files",
		Legend:                f.conf.Legend,
		SourceTable:           o.Table,
		Input:                 f.Input(),
		SkipCompletionAllowed: true,
		NotRequiredInDB:       true,
	})

	handler.AddNewFileSetting(f.Object, f.conf)

	// ASIGNAMOS LOS MANEJADORES API CRUD CORRESPONDIENTES AL OBJETO
	f.Object.Table = handler.Object.Table
	f.Object.BackHandler.ReadApi = handler
	f.Object.BackHandler.DeleteApi = handler
	f.Object.BackHandler.UpdateApi = handler
	f.Object.AlternativeValidateAdapter = handler
	// f.Object.

	return &f, ""

}

func (f FileInput) NameViewAdapter() string {
	return "FileInput"
}
