package fileinput

func (f FileInput) ResetAdapterView() (err string) {

	f.Log("info ResetViewAdapter FileInput", f.Object.ObjectName)

	return f.Object.CallJsFunctionObject("resetInputFile", f.conf.DefaultEnableInput)

}

func (f FileInput) InputEnable() (err string) {
	return f.Object.CallJsFunctionObject("enableFileInput", true)
}
