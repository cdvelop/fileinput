package fileinput

func (f FileInput) SetObjectInDomAfterDelete(data ...map[string]string) (err string) {
	// f.Log("SET DOM DESPUÉS DE ELIMINAR OBJETO EN DB")
	for _, data := range data {
		// f.Log("delete object in dom:", data)
		_, err = f.dom.CallFunction("deleteFile", data["id_file"])
		if err != "" {
			f.Log("SetObjectInDomAfterDelete error:", err)
		}
	}

	return ""
}
