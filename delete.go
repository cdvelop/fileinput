package fileinput

func (f FileInput) SetObjectInDomAfterDelete(data ...map[string]string) error {

	f.Log("SET DOM DESPUÉS DE ELIMINAR OBJETO EN DB")

	for _, data := range data {
		f.Log("data:", data)
	}

	return nil
}
