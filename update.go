package fileinput

import (
	"github.com/cdvelop/model"
)

func (f File) Update(u *model.User, data ...map[string]string) ([]map[string]string, error) {

	// fmt.Println("UPDATE DATA:", data)

	return f.db.UpdateObjectsInDB(f.Object.Table, data...)
}
