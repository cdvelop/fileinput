package fileinput

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (f FileInput) Delete(u *model.User, params ...map[string]string) ([]map[string]string, error) {

	// fmt.Println("parámetros Delete recibidos:", params)
	recover_data, err := f.App.ReadObjectsInDB(f.Object.Table, params...)
	if err != nil {
		return nil, err
	}

	err = f.App.DeleteObjectsInDB(f.Object.Table, params...)
	if err != nil {
		return nil, err
	}

	fmt.Println("DATA RECOBRADA DESPUÉS DE BORRAR: ", recover_data)
	if len(recover_data) == 0 {
		return nil, model.Error("el archivo ya no existe en la base de datos")
	}

	for _, data := range recover_data {

		file_path := f.BuildFilePath(data)

		// Borrar archivos desde hdd
		err := f.App.FileDelete(file_path)
		if err != nil {
			return nil, model.Error("archivo", data[f.Field_name], "fue eliminado de la db pero no del hdd", err)
		}
	}

	fmt.Println("DATA RECOBRADA DESPUÉS DE BORRAR: ", recover_data)

	return recover_data, nil
}

func (f FileInput) SetObjectInDomAfterDelete(data ...map[string]string) error {

	f.App.Log("SET DOM DESPUÉS DE ELIMINAR OBJETO EN DB")

	for _, data := range data {
		f.App.Log("data:", data)
	}

	return nil
}
