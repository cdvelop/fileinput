
let cameraVideo;
function openCapture(t) {
    // e.stopPropagation()
    let span = getSpanContainer(t)
    if (span != undefined) {
        const container = span.parentNode;
        const state = getState(container);

        let button = span.querySelector('button')
        button.classList.toggle('icon-selected');
        const form = container.parentNode;

        shiftContainer(form, container, span, state)


        if (cameraVideo === undefined) {
            cameraVideo = CameraVideo(form);
        } else if (cameraVideo.FormName() != form.name) {
            cameraVideo = CameraVideo(form);
        }

        if (state == "on") {
            cameraVideo.Enable();
        } else {// off
            cameraVideo.Disable();
        }
    }
}

function shiftContainer(form, container, span, state) {

    let shiftInY = 0;

    if (state == "on") {
        shiftInY = form.clientHeight - container.offsetTop;
    }
    container.style.transform = `translateY(` + shiftInY + `px)`;

    // Define la función de callback
    const transitionEndHandler = () => {
        // Realizar scroll
        container.scrollIntoView({ block: "end", behavior: "smooth" });
        // Eliminar el evento transitionend
        span.removeEventListener('transitionend', transitionEndHandler);
    };
    // Agrega el evento transitionend
    span.addEventListener('transitionend', transitionEndHandler);
    // expandir span
    span.classList.toggle('expand');//tarda 0.4s
}



function getState(container) {
    let state = container.dataset.state;
    if (state == undefined) {
        container.dataset.state = "on"
        state = "on"
    } else if (state == "on") {
        container.dataset.state = "off"
        state = "off"
    } else {
        container.dataset.state = "on"
        state = "on"
    }
    console.log("state:", state)
    return state
}


function getSpanContainer(t) {
    return t.closest('span')
}