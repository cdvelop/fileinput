function openCapture(e) {
    let span = getSpanContainer(e)
    if (span != undefined) {
        const container = span.parentNode;
        const state = getState(container);


        let legend = span.querySelector('legend')

        // console.log("span:", span, "legend:", legend)

        legend.classList.toggle('icon-selected');



        // const size = form.getBoundingClientRect();
        // console.log("form:", "size.width", size.width, "size.height", size.height)

        shiftContainer(container, span, state)



        if (state == "on") {



        } else {// off

        }

    }
}

function shiftContainer(container, span, state) {

    const form = container.parentNode;
    let size = form.dataset.size;
    let now_size = form.scrollHeight
    if (size == undefined) {
        form.dataset.size = now_size
        size = now_size
    }

    if (size != now_size) {
        console.log("formulario cambio tamaño: anterior:", size, "actual:", now_size, "solo realizamos scroll")
        form.dataset.size = now_size
        // Hacer scroll
        // container.scrollIntoView({ block: "end", behavior: "smooth" });
    } else {
        console.log("formulario tamaño idéntico", size, "muevo el contenedor")
        let shiftInY = 0;
        if (state == "on") {
            shiftInY = form.clientHeight - container.offsetTop - 16;
        }

        container.style.transform = `translateY(` + shiftInY + `px)`;
    }

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

function getSpanContainer(e) {
    const t = e.target
    const tagname = t.tagName.toLowerCase()
    // console.log("tag name:", tagname)
    let span;
    switch (tagname) {
        case "legend":
            span = e.currentTarget;
            break;
        case "svg":
            span = t.parentNode.parentNode;
            break;
        case "use":
            span = t.parentNode.parentNode.parentNode;
            break;
    }
    return span
}