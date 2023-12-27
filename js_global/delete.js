

function deleteFileHandler(target) {

    file_container = target.parentNode;

    // console.log('BORRAR ELEMENTO:', target);
    const obj_name = file_container.dataset.id;
    // console.log('obj_name ', obj_name);

    // const img = target.querySelector('img')
    // Realiza la acción deseada cuando el usuario mantiene presionado el botón durante 5 segundos.
    // console.log('ID', target.dataset.id);
    let ids = [target.dataset.id]

    deleteObject(obj_name, ids);
}


function deleteFile(id) {
    // console.log("delete file in dom", id)
    if (file_container != undefined) {
        let file = file_container.querySelector('[data-id="' + id + '"]')
        if (file != undefined) {
            // console.log("delete file ", file)
            file.remove();
        }
    }
}