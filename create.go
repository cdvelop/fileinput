package fileinput

import "github.com/cdvelop/model"

func (f File) RegisterNewFile(new *model.FileNewToStore, form_data map[string]string) (map[string]string, error) {

	form_data[f.Id_file] = new.FileNameOnDisk
	form_data[f.File_area] = new.FileArea
	form_data[f.Extension] = new.Extension

	// cortar el nombre del archivo para eliminar la extensión antes de almacenarlo
	if len(new.DescriptionInputName) > 5 {
		form_data[f.Description] = new.DescriptionInputName[:len(new.DescriptionInputName)-len(new.Extension)]
	}

	// borramos el campo files
	delete(form_data, f.Files)

	err := f.App.CreateObjectsInDB(f.Object.Table, false, form_data)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		f.Id_file:     new.FileNameOnDisk,
		f.Description: form_data[f.Description],
	}, nil
}

func (f File) GenerateFileNameOnDisk() string {
	return f.App.GetNewID()
}
