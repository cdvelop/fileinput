package fileinput

import (
	"github.com/cdvelop/model"
)

func (f File) Delete(u *model.User, params ...map[string]string) ([]map[string]string, error) {

	// fmt.Println("parámetros Delete recibidos:", params)

	recover_data, err := f.App.DeleteObjectsInDB(f.Object.Table, params...)
	if err != nil {
		return nil, err
	}

	// fmt.Println("DATA RECOBRADA DESPUÉS DE BORRAR: ", recover_data)

	for _, data := range recover_data {

		file_path := f.BuildFilePath(data)

		// Borrar archivos desde hdd
		err := f.App.DeleteFile(file_path)
		if err != nil {
			return nil, model.Error("archivo", data[f.Field_name], "fue eliminado de la db pero no del hdd", err)
		}
	}

	return recover_data, nil
}
