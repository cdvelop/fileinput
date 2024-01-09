

function openCapture(t) {

    let span = t.closest('span')
    if (span != undefined) {
        file_container = span.parentNode;
        // const form = file_container.parentNode;   

        const object_name = getObjectNameFromFileContainer(file_container)

        // console.log("openCapture object_name:", object_name)

        if (camera === undefined || camera.ObjectName != object_name) {
            camera = Camera(file_container);
            // console.log("NUEVA CÁMARA o IDS DIFERENTES")
        }

        // console.log("openCapture file_container:",file_container)
        const btn_capture = span.querySelector('button[name="capture"]')
        // console.log("btn_capture:", btn_capture)
        const txt = " captura de imagen"

        if (camera.IsClosed()) {
            btn_capture.title = "detener" + txt
            camera.Enable(file_container);
        } else {
            btn_capture.title = "iniciar" + txt
            camera.Disable();

            endFileCapture(object_name)
        }

    }

}



function getObjectNameFromFileContainer(file_container) {
    // const obj = file_container.querySelector(`div[name="files"]`).dataset.id

    // let name = obj.dataset.id

    // console.log("object name",name)

    return file_container.querySelector(`div[name="files"]`).dataset.id

}
