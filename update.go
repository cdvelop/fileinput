package fileinput

import (
	"github.com/cdvelop/model"
)

func (f File) Update(u *model.User, data ...map[string]string) error {

	// fmt.Println("UPDATE DATA:", data)

	return f.App.UpdateObjectsInDB(f.Object.Table, data...)
}
