package fileinput

func (f File) ResetInput() {
	f.callJsFunc("resetInputFile", f.DefaultEnableInput)
}

func (f File) InputEnable() {
	f.callJsFunc("enableFileInput", true)
}

func (f File) callJsFunc(func_name string, enable bool) {

	err := f.App.CallFunction(func_name, map[string]interface{}{
		"enable": enable,
		"query":  f.App.QuerySelectorObject(f.Object.ModuleName, f.Object.Name),
	})
	if err != nil {
		f.App.Log(err)
	}

}
