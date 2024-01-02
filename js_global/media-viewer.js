

function MediaViewer(file_container, new_content, stay_viewer_open = false) {

    const media_cont = file_container.querySelector('.media-container')

    // console.log("file_container", file_container)

    if (new_content) {
        // Verifica si ya hay un contenido en el contenedor de medios
        if (media_cont.firstChild) {
            // Si hay un hijo, reemplázalo con el nuevo contenido
            media_cont.replaceChild(new_content, media_cont.firstChild);
        } else {
            // Si no hay hijos, simplemente agrega el nuevo contenido al contenedor de medios
            media_cont.appendChild(new_content);
        }
    }


    if (stay_viewer_open && modal_displayed == true) {
        return
    }

    const form = file_container.parentNode;

    // console.log("stay_viewer_open:", stay_viewer_open, "modal_displayed now:", modal_displayed)
    // DESPLAZAMIENTO DEL CONTENEDOR
    let shiftInY = 0;

    if (modal_displayed == false) {// no definido, lo encendemos
        modal_displayed = true
        shiftInY = form.clientHeight - file_container.offsetTop - 35;

        media_cont.classList.remove("media-hidden");
    } else {
        modal_displayed = false// lo apagamos
        media_cont.classList.add("media-hidden");
    }


    file_container.style.transform = `translateY(` + shiftInY + `px)`;

    const modal = file_container.querySelector('span.modal-file-viewer')

    const transitionEndHandler = () => {
        // Realizar scroll
        // modal.scrollIntoView({ block: "start", behavior: "smooth" });
        if (modal_displayed) {
            modal.scrollIntoView({ block: "start", behavior: "smooth" });
        } else {
            form.scrollIntoView({ block: "end", behavior: "smooth" });
        }

        // Eliminar el evento transitionend
        modal.removeEventListener('transitionend', transitionEndHandler);
    };
    // Agrega el evento transitionend
    modal.addEventListener('transitionend', transitionEndHandler);


    // expandir span //tarda 0.4s
    if (modal.style.height != '') {
        modal.style.height = '';

    } else {
        let expand = (modal.clientHeight + form.clientHeight) - (field_height + 10)
        // console.log("expand:", expand)
        modal.style.height = expand + 'px';

    }
}

