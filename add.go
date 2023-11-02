package fileinput

import (
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
func New(o *model.Object, db model.DataBaseAdapter, c model.FileConfig, h *model.Handlers) (*File, error) {

	f := File{}

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

	f.source = o          // asignamos el objeto origen
	f.input_id = input_id // asignamos input id para verificaciones en lectura

	// configuración por defecto file
	f.FileConfig = model.FileConfig{
		MaximumFilesAllowed: 6,
		InputNameWithFiles:  f.Files,
		MaximumFileSize:     0,
		MaximumKbSize:       50,
		AllowedExtensions:   ",.jpg, .png, .jpeg",
		ImagenWidth:         "800",
		ImagenHeight:        "600",

		RootFolder: "app_files",
		FileType:   "i",

		SavedAsBlobInDb: c.SavedAsBlobInDb,

		FieldNameWithObjectID: o.PrimaryKeyName(),
		Name:                  c.Name,
		Legend:                "Imágenes",
	}

	if f.Name == "" {
		return nil, model.Error(`fileinput error FileConfig.Name:"nombre_campo" no ingresado`)
	}

	f.DefaultEnableInput = c.DefaultEnableInput

	if c.RootFolder != "" {
		f.RootFolder = c.RootFolder
	}

	if c.ImagenWidth != "" {
		f.ImagenWidth = c.ImagenWidth
	}

	if c.ImagenHeight != "" {
		f.ImagenHeight = c.ImagenHeight
	}

	if c.FileType != "" {
		f.FileConfig.FileType = c.FileType

		switch c.FileType {
		case "video":
			f.Legend = "Videos"
			f.FileConfig.AllowedExtensions = ".avi, .mkv, .mp4"
			f.FileType = "v"
		case "document":
			f.Legend = "Documentos"
			f.FileConfig.AllowedExtensions = ".doc, .xlsx, .txt"
			f.FileType = "d"
		case "pdf":
			f.Legend = "Documentos PDF"
			f.FileConfig.AllowedExtensions = ".pdf"
			f.FileType = "p"
		}

	}

	if c.MaximumFilesAllowed != f.MaximumFilesAllowed {
		f.MaximumFilesAllowed = c.MaximumFilesAllowed
	}

	if c.MaximumKbSize != f.FileConfig.MaximumKbSize {
		f.FileConfig.MaximumKbSize = c.MaximumKbSize
	}

	f.Object.Name += "." + f.FileConfig.Name

	f.FileConfig.MaximumFileSize = int64(float64(f.FileConfig.MaximumFilesAllowed*f.FileConfig.MaximumKbSize*1024) * 1.05)

	if !f.App.RunOnClientDB() { // verificamos la base de datos solo si estamos en el servidor
		err = db.CreateTablesInDB([]*model.Object{f.Object}, nil)
		// fmt.Println("ESTAMOS EN SERVIDOR CREAMOS TABLA ", f.Object.Table, " EN DB CON ERROR", err)
		if err != nil {
			return nil, err
		}
	}

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

	return &f, nil
}

func (f File) ConfigFile() *model.FileConfig {
	return &f.FileConfig
}

func (f File) ViewHandlerName() string {
	return "FileInput"
}
