
function imgFileSelected(e) {
    e.stopPropagation();
    // console.log("**EVENTO: imgFileSelected",e)
    const tagname = e.target.tagName.toLowerCase()
    if (tagname === "img") {
        // console.log("prev", e.target, "tagname:", tagname)
        // console.log("show_img:", e.target.show_img)

        // console.log(e.target);

        let id = getFileId(e.target)

        // console.log("id imagen:", id)
        // Obtiene una referencia al elemento con el ID "modal-container"
        // let modal = document.getElementById('modal-window');

        // Clona la imagen seleccionada
        let clonedImage = e.target.cloneNode(true);
        clonedImage.classList.add('full-screen-img');

        ShowModal(true, clonedImage)

        // modal.style.display = 'block';
        // Usa scrollIntoView para desplazarte al elemento
        //   modal.scrollIntoView({ behavior: 'smooth' }); // 'smooth' proporciona un desplazamiento suave
    }
}


function getFileId(img) {
    // let srcValue = img.getAttribute('src');
    // let parts = srcValue.split('=');
    return img.dataset.id;
}

function deleteFileON(e) {
    e.stopPropagation();
    // console.log("**EVENTO: deleteFileON",e)
    const tagname = e.target.tagName.toLowerCase()
    if (tagname === "img") {
        // console.log("tagname", tagname)
        // console.log("deleteFileON", e.target)
        startDeleteTimer(e.target, deleteFileHandler)
    }
}

function deleteFileOFF(e) {
    e.stopPropagation();
    // console.log("**EVENTO: deleteFileOFF",e)
    // e.stopPropagation()
    const tagname = e.target.tagName.toLowerCase()
    if (tagname === "img") {
        // console.log("tagname", tagname)
        // console.log("deleteFileOFF", e.target)
        clearDeleteTimer(e.target)
    }
}

function deleteFileHandler(target) {

    console.log('BORRAR ELEMENTO:', target);

    const obj_name = target.parentNode.parentNode.dataset.id;
    console.log('obj_name ', obj_name);


    // const img = target.querySelector('img')
    // Realiza la acción deseada cuando el usuario mantiene presionado el botón durante 5 segundos.
    console.log('ID', target.dataset.id);
}