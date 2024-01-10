package fileinput

func (f *FileInput) GamepadConnected() {
	// f.Log("ENCENDER ICONO GAME PAD")
	f.gamepad.NameJsFunc = "joystickEnable"
	f.gamepad.Enable = true

	f.Object.UserMessage("Joystick Conectado")

	_, err := f.gamepad.CallWithEnableAndQueryParams(f.Object)
	f.Log(err)

}

func (f *FileInput) GamepadDisconnected() {
	f.gamepad.NameJsFunc = "joystickEnable"
	f.gamepad.Enable = false
	// f.gamepadState = ` data-gamepad="off"`
	// f.Log("APAGAR ICONO GAME PAD")

	f.Object.UserMessage("Joystick Desconectado")

	_, err := f.gamepad.CallWithEnableAndQueryParams(f.Object)
	f.Log(err)
}

func (f FileInput) GamepadPressAnyButton() {
	// f.Log("TOMANDO FOTO CON GAME PAD")
	f.gamepad.NameJsFunc = "gamepadPhoto"
	f.gamepad.Enable = true
	_, err := f.gamepad.CallWithEnableAndQueryParams(f.Object)
	f.Log(err)
}
