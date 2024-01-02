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
     
     
     const button = fieldset.parentNode.querySelector('button[name="capture"]');
   
    if (enable) {
        fieldset.disabled = false
        button.disabled = false
    } else {
        fieldset.disabled = true
        button.disabled = true
    }


}