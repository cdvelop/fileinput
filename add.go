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
// MaximumFilesAllowed:1, 4, 6.. default 6
// max_kb_size:100, 400 default 50

func NewUploadFileApi(h *model.Handlers, o *model.Object, s filehandler.FileSetting) (*FileInput, error) {

	f := FileInput{}

	// agregar inputs usados en tabla file al modulo
	err := o.Module.AddInputs([]*model.Input{unixid.InputPK(), f.Input(), FileType(), SavedMode(), input.Text(), input.TextNumCode(), input.TextNum(), input.FilePath(), input.TextOnly()}, "file pkg")
	if err != nil {
		return nil, err
	}

	// crear objeto file
	err = object.AddToHandlerFromStructs(&f, o.Module, h)
	if err != nil {
		return nil, err
	}

	f.DomAdapter = h
	f.ThemeAdapter = h

	fileHandler, err := filehandler.Add(h, o.Module)
	if err != nil {
		return nil, err
	}
	f.FileHandler = fileHandler

	// configuración por defecto file
	f.FileSetting = &filehandler.FileSetting{
		MaximumFilesAllowed:   6,
		MaximumKbSize:         50,
		ImagenWidth:           "800",
		ImagenHeight:          "600",
		FileType:              "i",
		FieldNameWithObjectID: o.PrimaryKeyName(),
		DescriptiveName:       s.DescriptiveName,
		Legend:                "Imágenes",
		Source:                o, // asignamos el objeto origen
	}

	if f.FileSetting.DescriptiveName == "" {
		return nil, model.Error(`fileinput error FileSetting.DescriptiveName no ingresado`)
	}

	f.FileSetting.DefaultEnableInput = s.DefaultEnableInput

	if s.ImagenWidth != "" {
		f.FileSetting.ImagenWidth = s.ImagenWidth
	}

	if s.ImagenHeight != "" {
		f.FileSetting.ImagenHeight = s.ImagenHeight
	}

	if s.MaximumFilesAllowed != f.FileSetting.MaximumFilesAllowed {
		f.FileSetting.MaximumFilesAllowed = s.MaximumFilesAllowed
	}

	if s.MaximumKbSize != f.FileSetting.MaximumKbSize {
		f.FileSetting.MaximumKbSize = s.MaximumKbSize
	}

	f.Object.Name = f.FileSetting.DescriptiveName

	f.FileSetting.SetMaximumFileSize()

	f.FileSetting.SetAllowedExtensions()

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
