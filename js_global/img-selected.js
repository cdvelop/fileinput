
let last_img_id;
let file_container;
function imgFileSelected(target_img) {

    file_container = target_img.parentNode.parentNode.parentNode;

    // console.log("imgFileSelected file_container", file_container);

    // Clona la imagen seleccionada
    const cloned_image = target_img.querySelector('img').cloneNode(true)
    cloned_image.setAttribute('onclick', 'fullScreenImage(this)');
    // console.log("imgFileSelected cloned_image", cloned_image);
    let same = sameImageID(target_img)

    // si la cámara esta abierta hay que desactivarla
    if (camera != undefined && !camera.IsClosed()){
        camera.Disable();
    }

    // si la imagen seleccionada es la misma cerramos el visor
    MediaViewer(file_container, cloned_image,!same)

}


function fullScreenImage(img) {
    // console.log("fullScreenImage",img)
    img.classList.add('full-screen-img');
    ShowModal(true, img,closeScreenImage)
}

function closeScreenImage(img) {
    img.classList.remove('full-screen-img');
    // console.log("closeScreenImage",img);
    MediaViewer(file_container, img,true)
}


function sameImageID(img) {
    let same = false;

    let new_id = getFileId(img)
    if (last_img_id != undefined) {

        if (last_img_id === new_id) {
            same = true
        }
    }

    last_img_id = new_id

    return same
}

function getFileId(img) {
    return img.dataset.id;
}

function targetFileHandler(e) {
    e.stopPropagation();
    const tagname = e.target.tagName.toLowerCase()
    // console.log("**EVENTO: targetFileHandler",e.target)
    if (tagname === "img") {
        const target = e.target.parentNode;
        // console.log("targetFileHandler", target)
        // console.log("deleteFileON", e.target)
        targetHandler(target, imgFileSelected, deleteFileHandler)

    }
}


function deleteFileHandler(target) {

    console.log('BORRAR ELEMENTO:', target);

    const obj_name = target.parentNode.dataset.id;
    console.log('obj_name ', obj_name);

    // const img = target.querySelector('img')
    // Realiza la acción deseada cuando el usuario mantiene presionado el botón durante 5 segundos.
    console.log('ID', target.dataset.id);

    deleteObject(obj_name, target.dataset.id);
}