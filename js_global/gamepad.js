
let previousButtonState = [];

window.addEventListener("gamepadconnected", (e) => {
    console.log("gamepad connected");
});

window.addEventListener("gamepaddisconnected", (e) => {
    console.log("gamepad disconnected");
});

function gameLoop() {
    const gamepads = navigator.getGamepads();

    for (let i = 0; i < gamepads.length; i++) {
        const gamepad = gamepads[i];
        
        if (gamepad && gamepad.connected) {
            handleButtons(gamepad.buttons);
        }
    }

    requestAnimationFrame(gameLoop);
}

function handleButtons(buttons) {
    buttons.forEach((button, index) => {
        if (button.pressed && !previousButtonState[index]) {

            // console.log("previousButtonState:",previousButtonState)

            console.log(`TOMANDO FOTO CON MANDO ${index}`);
            // Aquí puedes realizar acciones adicionales cuando se presiona un botón
            // console.log("ESTADO CÁMARA ESTA CERRADA?:", camera_is_closed)
            // turnOnCamera()
            takePicture()
        }
        previousButtonState[index] = button.pressed;
    });
}

gameLoop();



// function turnOnCamera() {
//     if (camera_is_closed) {
//         startupCapture()
//     } else {
//         takePicture();
//     }

// }