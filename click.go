package fileinput

func (f FileInput) CameraTurn(on bool) {
	var turn = "on"

	if !on {
		turn = "off"
	}

	err := f.Clicking(f.Object, turn)
	if err != nil {
		f.Log(err)
		return
	}

}
