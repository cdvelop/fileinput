package fileinput

func (f File) CameraTurn(on bool) {
	var turn = "on"

	if !on {
		turn = "off"
	}

	err := f.App.Clicking(f.Object, turn)
	if err != nil {
		f.App.Log(err)
		return
	}

}
