package fileinput

func (f *FileInput) GamePadConnected() {
	// f.Log("ENCENDER ICONO GAME PAD")
	f.gamepad_icon.Enable = true

	_, err := f.gamepad_icon.CallWithEnableAndQueryParams(f.Object)
	f.Log(err)

}

func (f *FileInput) GamePadDisconnected() {
	f.gamepad_icon.Enable = false
	// f.gamepadState = ` data-gamepad="off"`
	// f.Log("APAGAR ICONO GAME PAD")
	_, err := f.gamepad_icon.CallWithEnableAndQueryParams(f.Object)
	f.Log(err)
}

func (f FileInput) TakePhotoWithGamePad() {
	// f.Log("TOMANDO FOTO CON GAME PAD")
	f.gamepad_photo.Enable = true
	_, err := f.gamepad_photo.CallWithEnableAndQueryParams(f.Object)
	f.Log(err)
}
