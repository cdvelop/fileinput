function resetInputFile(o) {
    //1 apagar la cámara si esta encendida
    // console.log("RESET INPUT FILE:", o)
    if (camera != undefined) {
        if (!camera.IsClosed()) {
            camera.Disable();
        }
    }
     
    // 2 limpiando input
    const cont_files = o.form.querySelector('[name="' + o.field_name + '"]')
    // const cont_files = document.querySelector(o.query)

    cont_files.innerHTML = "";
    //  o["cont_files"]=cont_files;
    fileInputEnabled(cont_files, o.enable)

    resetFileInputVars()
}

function enableFileInput(o) {
    // console.log("enableFileInput",o)
    const cont_files = document.querySelector(o.query)
    fileInputEnabled(cont_files, o.enable)

}

function fileInputEnabled(cont_files, enable) {

     const fieldset = cont_files.closest('fieldset')
     
     
     const btn_capture = fieldset.parentNode.querySelector('button[name="capture"]');
     const btn_joystick = fieldset.parentNode.querySelector('button[name="joystick"]');

    //  console.log(btn_joystick)
   
    if (enable) {
        fieldset.disabled = false
        btn_capture.disabled = false
        btn_joystick.disabled = false
    } else {
        fieldset.disabled = true
        btn_capture.disabled = true
        btn_joystick.disabled = true
    }


}