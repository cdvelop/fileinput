
function Camera(container_files) {
    let camera_is_closed = true;

    const form = container_files.parentNode;

    let video = createVideoTag(container_files);
    
    const canvas = container_files.querySelector('canvas[name="canvas"')
    // photo default config
    canvas.width = parseInt(video.dataset.width);
    canvas.height = parseInt(video.dataset.height);


    const button = form.querySelector('button[name="capture"]')

    const compute_style = window.getComputedStyle(form);
    let width = parseInt(compute_style.getPropertyValue('width'));
    let height = 0;

    function canplayListener() {
        height = video.videoHeight / (video.videoWidth / width);
        if (isNaN(height)) {
            height = width / (4 / 3);
        }
        video.setAttribute("width", width);
        video.setAttribute("height", height);
    }

    function enableCameraCapture(file_container) {
        container_files = file_container;
        let stay_viewer_open = false;

        // console.log("enableCameraCapture viewer_status", file_container.dataset.display_status)
        if (file_container.dataset.display_status === "on") {
            stay_viewer_open = true
        }
        
        MediaViewer(file_container, video, stay_viewer_open)
        camera_is_closed = false;
        navigator.mediaDevices
        .getUserMedia({ video: true, audio: false })
        .then((stream) => {
            video.srcObject = stream;
            video.play();
        })
        .catch((err) => {
            console.error("An error occurred:", err);
        });
        
        video.addEventListener("canplay", canplayListener, false);
        
        button.classList.add('icon-selected');
    }
    
    function disableCameraCapture() {
        // console.log("disableCameraCapture viewer_status", viewer_status)
        if (!camera_is_closed) {
            MediaViewer(container_files, video, false)
            video.srcObject.getTracks().forEach((track) => track.stop());
            camera_is_closed = true;
            video.removeEventListener("canplay", canplayListener, false);

            // Para ocultar el elemento de video suavemente

            button.classList.remove('icon-selected');
        }
    }

    function formName() {
        return form.name;
    }

    function isCameraClosed() {
        return camera_is_closed;
    }

    const previous_img = container_files.querySelector('[name="previous_img"]')
    function showNewPicture() {
        moveScrollFileImg(previous_img)
    }

    const object_name = getObjectNameFromFileContainer(container_files)

    return {
        Enable: enableCameraCapture,
        Disable: disableCameraCapture,
        ObjectName: object_name,
        FormName: formName,
        IsClosed: isCameraClosed,
        Canvas: canvas,
        Video: video,
        ShowNewPicture: showNewPicture
    };
}


function createVideoTag(container_files) {
    // Crea un nuevo elemento video
    var video = document.createElement('video');

    // Establece los atributos del video con la configuración proporcionada
    video.setAttribute('data-width', container_files.dataset.width);
    video.setAttribute('data-height', container_files.dataset.height);
    video.setAttribute('title', 'toca el video para tomar una captura');
    video.setAttribute('onclick', 'takePicture(this)');
    video.setAttribute('id', 'video_capture');

    return video
}