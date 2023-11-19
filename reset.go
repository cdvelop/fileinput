package fileinput

func (f FileInput) ResetInput() {

	f.callJsFunc("resetInputFile", f.conf.DefaultEnableInput)
}

func (f FileInput) InputEnable() {
	f.callJsFunc("enableFileInput", true)
}

func (f FileInput) callJsFunc(func_name string, enable bool) {

	err := f.dom.CallFunction(func_name, map[string]interface{}{
		"enable": enable,
		"query":  f.theme.QuerySelectorObject(f.Object.ModuleName, f.Object.ObjectName),
	})
	if err != nil {
		f.Log(err)
	}

}
