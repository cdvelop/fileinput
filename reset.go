package fileinput

func (f FileInput) InputEnable() (err string) {
	return f.input_enable.ExecuteJsFun(f.Object)
}

func (f FileInput) ResetViewHandlerObject() (err string) {
	return
}
