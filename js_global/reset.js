function resetInputFile(o) {

    //1 apagar la cámara si esta encendida
    // console.log("RESET INPUT FILE:", o)

    if (cameraVideo != undefined) {
        if (!cameraVideo.IsClosed()) {
            cameraVideo.Disable();
        }
    }
    // 2 limpiando input
    const cont_files = document.querySelector(o.query)

    cont_files.innerHTML = "";

    //  o["cont_files"]=cont_files;
    enableFileInput(o)

}

function enableFileInput(o) {
    // console.log("enableFileInput",o)
   
    const cont_files = document.querySelector(o.query)

    const fieldset = cont_files.closest('fieldset')

    const button = fieldset.parentNode.querySelector('button[name="capture"]');
    if (!o.enable) {
        fieldset.disabled = true
        button.disabled = true
    } else {
        fieldset.disabled = false
        button.disabled = false
    }

}