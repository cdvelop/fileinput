
function deleteFileHandler(target) {
    let obj_container = target.parentNode;
    // console.log('deleteFileHandler:', obj_container);
    const obj_name = obj_container.dataset.id;
    // console.log('obj_name ', obj_name);
    // const img = target.querySelector('img')
    // Realiza la acción deseada cuando el usuario mantiene presionado el botón durante 5 segundos.
    // console.log('ID', target.dataset.id);
    let ids = [target.dataset.id]

    deleteObject(obj_name, ids);
}


function deleteFileInput(o) {
    // console.log("delete file in dom", o)
    const container = document.querySelector(o.query)
    // console.log("container", container)
    if (container != undefined) {
        let item = container.querySelector('[data-id="' + o.id + '"]')
        if (item != undefined) {
            item.remove();
            // cierro el visor solo si esta desplegado y la cámara esta apagada
            if (modal_displayed && camera && camera.IsClosed()) {
                MediaViewer(container, "", false)
            }
        }
    }
}