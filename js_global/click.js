

function FileInputObjectClicking(module, opt) {
    const container = module.querySelector(`.file-input-container`);
    // console.log("click en:", container)
    if (opt == "on" || opt == "off") {
        clickButtonCapture(container);
        console.log(opt," cámara")
        
    }else{
        console.log("click en id?", opt)
    }
}




function clickButtonCapture(container) {
    var captureButton = container.querySelector("button[name='capture']");
    // Dispara el evento de clic en el botón
    captureButton.click();
}