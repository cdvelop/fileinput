

function gamepadPhoto(o) {
    const container = joystickEnable(o)
    // console.log("container", container)
    var captureButton = container.querySelector("button[name='capture']");
    // console.log("captureButton", captureButton)
    if (!captureButton.classList.contains("icon-selected")) {
        // console.log("la cámara esta apagada hay que encenderla")
        captureButton.click();
    } else {
        // console.log(" la cámara esta iniciada lista para sacar foto'");
        takePicture();
    }
}


function joystickEnable(o) {
    const object = document.querySelector(o.query);
    // console.log("objecto:",object)

    if (object === null) {
        return
    }

    const container = object.parentNode.parentNode;

    const button = container.querySelector('span.modal-file-viewer button[name="joystick"]')

    if (o.enable) {
        button.classList.add('icon-selected');
        button.title = "joystick conectado"

    } else {
        button.classList.remove('icon-selected');
        button.title = "joystick desconectado"
    }

    return container
}