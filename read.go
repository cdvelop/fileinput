package fileinput

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (f File) GetFilePathByID(params map[string]string) (file_path, file_area string, err error) {
	// fmt.Println("parámetros FilePath recibidos: ", params)

	if len(params) != 1 {
		return "", "", model.Error("solo se puede recibir un parámetro para leer un archivo")
	}

	id, ok := params["id"]
	if !ok {
		return "", "", model.Error("parámetro id no enviado para leer archivo")
	}

	err = f.input_id.ValidateField(id, false)
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

	file_path = f.BuildFilePath(data[0])
	file_area = data[0][f.File_area]

	// fmt.Println("AREA ARCHIVO: s, DB:" + file_area)
	// fmt.Println("PATH:" + file_path)

	return file_path, file_area, nil
}

func (f File) Read(u *model.User, params ...map[string]string) ([]map[string]string, error) {

	// var file_data string
	// if f.AddBinaryInReadResponse {
	// 	file_data = "," + f.File_data
	// }
	fmt.Println("parámetros leer recibidos:", params)

	// for _, v := range params {
	// 	v["SELECT"] = f.Id_file + ","+ f.Module_name + "," + f.Folder_id + "," + f.Description + file_data
	// }

	data, err := f.db.ReadObjectsInDB(f.Object.Table, params...)
	if err != nil {
		return nil, err
	}
	fmt.Println("data devuelta:", data)

	return data, nil
}
