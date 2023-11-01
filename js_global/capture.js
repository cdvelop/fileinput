
let camera;
function openCapture(t) {
    let span = t.closest('span')
    if (span != undefined) {
        const container = span.parentNode;
        const form = container.parentNode;


        if (camera === undefined) {
            camera = Camera(form);
        } else if (camera.FormName() != form.name) {
            camera = Camera(form);
        }

        if (camera.IsClosed()) {
            camera.Enable();
        } else {
            camera.Disable();
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




