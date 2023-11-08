package fileinput

import (
	"github.com/cdvelop/model"
)

func (f File) AddBootFiles(u *model.User, out *[]model.Response, from_data ...map[string]string) error {

	field_ids := []map[string]string{}
	for _, data := range from_data {
		field_ids = append(field_ids, map[string]string{
			f.Object_id: data[f.FieldNameWithObjectID],
			"SELECT":    f.Id_file + "," + f.Object_id + "," + f.Description,
		})
	}

	our_files, err := f.Read(u, field_ids...)
	if err != nil {
		return err
	}

	if len(our_files) >= 1 {
		*out = append(*out, f.Object.Response(our_files))
	}
	// fmt.Println("ARCHIVOS LEÍDOS:", len(our_files))

	return nil
}
