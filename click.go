package fileinput

func (f FileInput) CameraTurn(on bool) {
	var turn = "on"

	if !on {
		turn = "off"
	}

	err := f.Object.ClickingID(turn)
	if err != "" {
		f.Log(err)
		return
	}

}
