package fileinput

func (f FileInput) ResetInput() {

	f.callJsFunc("resetInputFile", f.DefaultEnableInput)
}

func (f FileInput) InputEnable() {
	f.callJsFunc("enableFileInput", true)
}

func (f FileInput) callJsFunc(func_name string, enable bool) {

	err := f.App.CallFunction(func_name, map[string]interface{}{
		"enable": enable,
		"query":  f.App.QuerySelectorObject(f.Object.ModuleName, f.Object.Name),
	})
	if err != nil {
		f.App.Log(err)
	}

}
