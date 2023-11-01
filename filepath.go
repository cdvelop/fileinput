package fileinput

func (f File) BuildFilePath(data map[string]string) (file_path string) {

	return f.UploadFolderPath(data) + "/" + data[f.Id_file] + data[f.Extension]
}

func (f File) UploadFolderPath(data map[string]string) string {
	return f.RootFolder + "/" + data[f.Module_name] + "/" + data[f.Field_name] + "/" + data[f.Object_id]
}
