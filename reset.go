package fileinput

func (f FileInput) InputEnable() (err string) {

	_, err = f.input_enable.CallWithEnableAndQueryParams(f.Object)

	return
}

func (f FileInput) ResetFrontendObjectState() (err string) {
	_, err = f.input_reset.CallWithEnableAndQueryParams(f.Object)
	return
}
