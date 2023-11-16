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

func NewUploadFileApi(h *model.Handlers, o *model.Object, s filehandler.FileSetting) (*FileInput, error) {

	f := FileInput{}

	// agregar inputs usados en tabla file al modulo
	err := o.Module.AddInputs([]*model.Input{f.Input(), FileType(), SavedMode(), input.FilePath(), input.TextOnly(), input.Text()}, "fileinput pkg")
	if err != nil {
		return nil, err
	}

	// crear objeto file upload
	err = object.AddToHandlerFromStructs(&f, o.Module, h)
	if err != nil {
		return nil, err
	}
	f.Logger = h.Logger

	f.dom = h
	f.theme = h

	handler, err := filehandler.Add(h, h, h)
	if err != nil {
		return nil, err
	}

	// configuración por defecto file
	f.conf = &filehandler.FileSetting{
		MaximumFilesAllowed:   6,
		MaximumKbSize:         50,
		ImagenWidth:           "800",
		ImagenHeight:          "600",
		FileType:              "imagen",
		FieldNameWithObjectID: o.PrimaryKeyName(),
		DescriptiveName:       s.DescriptiveName,
		Legend:                "Imágenes",
		Source:                o, // asignamos el objeto origen
	}

	if f.conf.DescriptiveName == "" {
		return nil, model.Error(`fileinput error FileSetting DescriptiveName no ingresado`)
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

	f.Object.Name = f.conf.DescriptiveName

	f.conf.SetMaximumFileSize()

	f.conf.SetAllowedExtensions()

	//nota: al no declarar punteros se pierden posteriormente

	f.Object.ViewHandler = f

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

	return &f, nil

}

func (f FileInput) ViewHandlerName() string {
	return "FileInput"
}
