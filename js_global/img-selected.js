function targetFileHandler(e) {
    e.stopPropagation();

    // console.log("targetFileHandler",e.target)
    // const tagname = e.target.tagName.toLowerCase()
    // console.log("**EVENTO: targetFileHandler",e.target,"tagname:", tagname)
    targetHandler(e.target, imgFileSelected, deleteFileHandler)
}



function imgFileSelected(target_img) {
    // console.log("imgFileSelected:", target_img);

    file_container = target_img.parentNode.parentNode.parentNode;

    // Clona la imagen seleccionada
    const cloned_image = target_img.querySelector('img').cloneNode(true)
    cloned_image.setAttribute('onclick', 'fullScreenImage(this)');
    // console.log("imgFileSelected cloned_image", cloned_image);

    
    // si la cámara esta abierta hay que desactivarla
    if (camera != undefined && !camera.IsClosed()) {
        // como esta abierto el visor enviamos la imagen a traves de la cam
        camera.Disable(cloned_image);
        
    }else{
        
        let same = sameImageID(target_img)
        // si la imagen seleccionada es la misma cerramos el visor
        MediaViewer(file_container, cloned_image, !same)
    }


    target_img.classList.toggle("file-selected");

    if (last_file_selected != undefined){
        last_file_selected.classList.toggle("file-selected");
    }


    last_file_selected = target_img
}


function fullScreenImage(img) {
    // console.log("fullScreenImage",img)
    img.classList.add('full-screen-img');
    ShowModal(true, img, closeScreenImage)
}

function closeScreenImage(img) {
    img.classList.remove('full-screen-img');
    // console.log("closeScreenImage",img);
    MediaViewer(file_container, img, true)
}


function sameImageID(img) {
    let same = false;

    let new_id = getFileId(img)
    if (last_file_selected != undefined) {

        let last_id = getFileId(last_file_selected)

        if (last_id === new_id) {
            same = true
        }
    }

    return same
}

function getFileId(img) {
    return img.dataset.id;
}




