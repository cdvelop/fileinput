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
func New(m *model.Module, db model.DataBaseAdapter, c model.FileConfig, h *model.Handlers) (*File, error) {

	err := m.AddInputs([]*model.Input{unixid.InputPK(), input.TextNumCode(), input.TextNum(), input.FilePath(), input.Text()}, "file pkg")
	if err != nil {
		return nil, err
	}

	f := File{}

	input := f.Input()

	// agregar inputs al modulo
	m.AddInputs([]*model.Input{input}, "file input")

	// crear objeto file
	err = object.New(&f, m, h)
	if err != nil {
		return nil, err
	}

	f.FileConfig = model.FileConfig{
		MaximumFilesAllowed: 6,
		InputNameWithFiles:  f.Files,
		MaximumFileSize:     0,
		MaximumKbSize:       50,
		AllowedExtensions:   ",.jpg, .png, .jpeg",

		RootFolder: "app_files",
		FileType:   "imagen",

		IdFieldName: "",
		Name:        "",
		Legend:      "Imágenes",
	}

	f.db = db

	if c.Name == "" {
		return nil, model.Error(`fileinput error FileConfig.Name:"nombre_campo" no ingresado`)
	} else {
		f.FileConfig.Name = c.Name
	}

	f.DefaultEnableInput = c.DefaultEnableInput

	if c.RootFolder != "" {
		f.RootFolder = c.RootFolder
	}

	if c.FileType != "" {
		f.FileConfig.FileType = c.FileType

		switch c.FileType {
		case "video":
			f.Legend = "Videos"
			f.FileConfig.AllowedExtensions = ".avi, .mkv, .mp4"
		case "document":
			f.Legend = "Documentos"
			f.FileConfig.AllowedExtensions = ".doc, .xlsx, .txt"
		case "pdf":
			f.Legend = "Documentos PDF"
			f.FileConfig.AllowedExtensions = ".pdf"
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

	if !f.db.RunOnClientDB() { // verificamos la base de datos solo si estamos en el servidor
		err = db.CreateTablesInDB([]*model.Object{f.Object}, nil)
		// fmt.Println("ESTAMOS EN SERVIDOR CREAMOS TABLA ", f.Object.Table, " EN DB CON ERROR", err)
		if err != nil {
			return nil, err
		}
	}

	//nota: al no declarar punteros se pierden posteriormente

	f.Object.ViewHandler = f

	return &f, nil
}

func (f File) ConfigFile() *model.FileConfig {
	return &f.FileConfig
}
