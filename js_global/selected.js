function imgFileSelected(e) {
    e.stopPropagation();
    const tagname = e.target.tagName.toLowerCase()
    if (tagname === "img") {
        // console.log("prev", e.target, "tagname:", tagname)

        console.log(e.target);

        let id = getFileId(e.target)

        console.log("id imagen:", id)

        // Obtiene una referencia al elemento con el ID "modal-container"
        // let modal = document.getElementById('modal-window');

        // Clona la imagen seleccionada
        let clonedImage = e.target.cloneNode(true);
        clonedImage.classList.add('full-screen-img');

        ShowModal(true, clonedImage)

        // console.log("modal:", modal)

        // modal.style.display = 'block';
        // Usa scrollIntoView para desplazarte al elemento
        //   modal.scrollIntoView({ behavior: 'smooth' }); // 'smooth' proporciona un desplazamiento suave
    }
}


function getFileId(img) {
    let srcValue = img.getAttribute('src');
    let parts = srcValue.split('=');
    return parts[1];
}