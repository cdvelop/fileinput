package fileinput

func (f File) RegisterNewFile(header_name, upload_folder, file_name, extension string, form_data map[string]string) (map[string]string, error) {

	form_data[f.Id_file] = file_name
	form_data[f.File_path] = upload_folder + "/" + file_name

	// cortar el nombre del archivo para eliminar la extensión antes de almacenarlo
	if len(header_name) > 5 {
		form_data[f.Description] = header_name[:len(header_name)-len(extension)]
	}

	// borramos el campo files
	delete(form_data, f.Files)

	err := f.db.CreateObjectsInDB(f.Object.Table, form_data)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		f.Id_file:     file_name,
		f.Description: form_data[f.Description],
	}, nil
}

func (f File) FileName(extension string) string {
	file_name := f.idh.GetNewID()
	return file_name + extension
}

func (f File) UploadFolderPath(form_data map[string]string) string {
	return f.root_folder + "/" + form_data[f.Module_name] + "/" + form_data[f.Field_name] + "/" + form_data[f.Folder_id]
}
