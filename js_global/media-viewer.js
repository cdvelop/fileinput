

function MediaViewer(file_container, new_content, stay_viewer_open = false) {

    const media_cont = file_container.querySelector('.media-container')

    // console.log("media_cont", media_cont)

    // Verifica si ya hay un contenido en el contenedor de medios
    if (media_cont.firstChild) {
        // Si hay un hijo, reemplázalo con el nuevo contenido
        media_cont.replaceChild(new_content, media_cont.firstChild);
    } else {
        // Si no hay hijos, simplemente agrega el nuevo contenido al contenedor de medios
        media_cont.appendChild(new_content);
    }

    let display_status = file_container.dataset.display_status

    if (stay_viewer_open && display_status === "on") {
        return
    }

    const modal = file_container.querySelector('span.modal-file-viewer')

    const form = file_container.parentNode;

    // DESPLAZAMIENTO DEL CONTENEDOR
    let shiftInY = 0;


    if (display_status === undefined || display_status === "") {// no definido, lo encendemos
        display_status = "on"
        // shiftInY = file_container.offsetTop - 36;
        shiftInY = form.clientHeight - file_container.offsetTop - 35;
        // shiftInY = file_container.offsetTop - form.clientHeight;
        // console.log("file_container.offsetTop - 35",file_container.offsetTop - 35)
                
        media_cont.classList.remove("media-hidden");
    } else {
        display_status = ""// lo apagamos
        media_cont.classList.add("media-hidden");
    }
    
    file_container.dataset.display_status = display_status;
    
    file_container.style.transform = `translateY(` + shiftInY + `px)`;
    
 
    const transitionEndHandler = () => {
        // Realizar scroll
        // modal.scrollIntoView({ block: "start", behavior: "smooth" });
        modal.scrollIntoView({ block: "start", behavior: "smooth" });
        // file_container.scrollIntoView({ block: "end", behavior: "smooth" });
                
        // Eliminar el evento transitionend
        modal.removeEventListener('transitionend', transitionEndHandler);
    };
    // Agrega el evento transitionend
    modal.addEventListener('transitionend', transitionEndHandler);
   
    
    // expandir span //tarda 0.4s
    if (modal.style.height != '') {
        modal.style.height = '';

    }else{
        let expand = (modal.clientHeight + form.clientHeight) - (field_height + 10 )
        // console.log("expand:", expand)
        modal.style.height = expand + 'px';

    }
}

