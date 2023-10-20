package fileinput

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (f File) GetFilePathByID(params map[string]string) (file_path, file_area string, err error) {
	fmt.Println("parámetros FilePath recibidos: ", params)

	if len(params) != 1 {
		return "", "", model.Error("solo se puede recibir un parámetro para leer un archivo")
	}

	id, ok := params["id"]
	if !ok {
		return "", "", model.Error("parámetro id no enviado para leer archivo")
	}

	err = f.ValidateField(id, false)
	if err != nil {
		return "", "", err
	}

	data, err := f.db.ReadObjectsInDB(f.Object.Table, map[string]string{
		f.Id_file: id,
	})
	if err != nil {
		return "", "", err
	}

	if len(data) != 1 {
		return "", "", model.Error("parámetros incorrectos al recuperar archivo")
	}

	file_path = data[0][f.File_path]
	file_area = data[0][f.File_area]

	fmt.Println("AREA ARCHIVO = A, DB=", file_area)
	file_area = "A"

	return file_path, file_area, nil
}

func (f File) Read(u *model.User, params ...map[string]string) ([]map[string]string, error) {

	// fmt.Println("parámetros leer recibidos:", params)

	for _, v := range params {
		v["SELECT"] = f.Module_name + "," + f.Field_name + "," + f.Folder_id + "," + f.Description
	}

	data, err := f.db.ReadObjectsInDB(f.Object.Table, params...)
	if err != nil {
		return nil, err
	}

	return data, nil
}
