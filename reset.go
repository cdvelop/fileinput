package fileinput

func (f FileInput) ResetAdapterView() {

	f.Log("ResetViewAdapter FileInput", f.Object.ObjectName)

	f.Object.CallJsFunctionObject("resetInputFile", f.conf.DefaultEnableInput)

}

func (f FileInput) InputEnable() {
	f.Object.CallJsFunctionObject("enableFileInput", true)
}
