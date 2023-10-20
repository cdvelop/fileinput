package fileinput

import (
	"strconv"
	"strings"

	"github.com/cdvelop/gotools"
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
func New(m *model.Module, db model.DataBaseAdapter, conf ...string) (*File, error) {

	err := m.AddInputs([]*model.Input{unixid.InputPK(), input.TextNumCode(), input.TextNum(), input.FilePath(), input.Text()}, "file pkg")
	if err != nil {
		return nil, err
	}

	f := File{}

	input := f.Input()

	m.AddInputs([]*model.Input{input}, "file input")

	// crear objeto
	err = object.New(&f, m)
	if err != nil {
		return nil, err
	}

	f.FileConfig = model.FileConfig{
		MaximumFilesAllowed: 6,
		InputNameWithFiles:  f.Files,
		MaximumFileSize:     0,
		MaximumKbSize:       50,
		AllowedExtensions:   ".jpg, .png, .jpeg",
	}

	f.db = db
	f.filetype = "imagen"
	f.root_folder = "app_files"

	var field_name string

	for _, option := range conf {

		switch {

		case strings.Contains(option, "field_name:"):
			gotools.ExtractTwoPointArgument(option, &field_name)

		case strings.Contains(option, "root_folder:"):
			gotools.ExtractTwoPointArgument(option, &f.root_folder)

		case strings.Contains(option, "filetype:"):
			gotools.ExtractTwoPointArgument(option, &f.filetype)

			switch f.filetype {
			case "video":
				f.FileConfig.AllowedExtensions = ".avi, .mkv, .mp4"
			case "document":
				f.FileConfig.AllowedExtensions = ".doc, .xlsx, .txt"
			case "pdf":
				f.FileConfig.AllowedExtensions = ".pdf"
			}

		case strings.Contains(option, "max_files:"):
			var max_files string
			gotools.ExtractTwoPointArgument(option, &max_files)

			num, err := strconv.ParseInt(max_files, 10, 64)
			if err != nil {
				return nil, model.Error("Error al convertir max_files la cadena: " + max_files + " " + err.Error())
			}
			f.FileConfig.MaximumFilesAllowed = num

		case strings.Contains(option, "max_kb_size:"):
			var max_kb_size string
			gotools.ExtractTwoPointArgument(option, &max_kb_size)

			num, err := strconv.ParseInt(max_kb_size, 10, 64)
			if err != nil {
				return nil, model.Error("Error al convertir max_kb_size la cadena: " + max_kb_size + " " + err.Error())
			}
			f.FileConfig.MaximumKbSize = num

		}
	}

	if field_name == "" {
		return nil, model.Error(`fileinput error field_name:"nombre_campo" no ingresado`)
	}

	f.Object.Name += "." + field_name

	f.FileConfig.MaximumFileSize = int64(float64(f.FileConfig.MaximumFilesAllowed*f.FileConfig.MaximumKbSize*1024) * 1.05)

	if !f.db.RunOnClientDB() { // verificamos la base de datos solo si estamos en el servidor
		err = db.CreateTablesInDB([]*model.Object{f.Object}, nil)
		// fmt.Println("ESTAMOS EN SERVIDOR CREAMOS TABLA ", f.Object.Table, " EN DB CON ERROR", err)
		if err != nil {
			return nil, err
		}
	}

	//nota: al no declarar punteros se pierden posteriormente

	return &f, nil
}

func (f File) ConfigFile() *model.FileConfig {
	return &f.FileConfig
}
