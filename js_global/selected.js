
function imgFileSelected(target) {
    // Clona la imagen seleccionada
    const clonedImage = target.querySelector('img').cloneNode(true)
    // console.log("IMG:",clonedImage);

    clonedImage.classList.add('full-screen-img');

    ShowModal(true, clonedImage)
}


function getFileId(img) {
    // let srcValue = img.getAttribute('src');
    // let parts = srcValue.split('=');
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

    deleteObject(obj_name,target.dataset.id);
}