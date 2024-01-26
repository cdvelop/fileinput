package fileinput

func (f *FileInput) NotifyStatusChangeAfterClicking() {
	// fmt.Println("notificación click FileInput")
}

func (f *FileInput) SetObjectInDomAfterDelete(data ...map[string]string) (err string) {
	// FileInput es un objeto de depende de otro ("padre")
	// por lo tanto notificamos a el para actualizar la vista u otro elemento si se requiere
	// fmt.Println("notificación SetObjectInDomAfterDelete FileInput")

	if f.notifyDelete != nil {
		f.notifyDelete.NotifyFileDeletion()
	}

	return
}
