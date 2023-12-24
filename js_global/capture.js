
let camera;
function openCapture(t) {
    
    let span = t.closest('span')
    if (span != undefined) {
        const file_container = span.parentNode;
        // const form = file_container.parentNode;        
        
        const object_name = getObjectNameFromFileContainer(file_container)

        // console.log("openCapture object_name:",object_name)
        
        if (camera === undefined || camera.ObjectName != object_name) {
            camera = Camera(file_container);
            // console.log("NUEVA CÁMARA o IDS DIFERENTES")
        } 
        
        // console.log("openCapture file_container:",file_container)


        if (camera.IsClosed()) {
            camera.Enable(file_container.dataset.display_status);
        } else {
            camera.Disable();
        }
     
    }
    
}



function getObjectNameFromFileContainer(file_container) {
    // const obj = file_container.querySelector(`div[name="files"]`).dataset.id

    // let name = obj.dataset.id

    // console.log("object name",name)

    return file_container.querySelector(`div[name="files"]`).dataset.id

}
