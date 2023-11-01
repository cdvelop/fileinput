

function Camera(form) {
    let camera_is_closed = true;
    const container_files = form.querySelector('.container-files')
    const canvas = container_files.querySelector('canvas[name="canvas"')
    const span = container_files.querySelector('span.modal-capture')

    const video_cont = form.querySelector('.video-container')

    let video = form.querySelector('#video_capture');

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

    function enableCameraCapture() {
        shiftContainer(form, container_files, span, "on")
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
        // video.style.display = "block";
        // video.classList.remove("video-hidden");
        video_cont.classList.remove("video-hidden");
        button.classList.add('icon-selected');
    }

    function disableCameraCapture() {
        if (!camera_is_closed) {
            shiftContainer(form, container_files, span, "off")
            video.srcObject.getTracks().forEach((track) => track.stop());
            camera_is_closed = true;
            video.removeEventListener("canplay", canplayListener, false);
            // video.style.display = "none"; // Ocultar el elemento de video
            // Para ocultar el elemento de video suavemente
            video_cont.classList.add("video-hidden");
            button.classList.remove('icon-selected');
        }
    }

    function formName() {
        return form.name;
    }

    function isCameraClosed() {
        return camera_is_closed;
    }

    // photo default config
    canvas.width = parseInt(video.dataset.width);
    canvas.height = parseInt(video.dataset.height);

    const object_id = getObjectIdFromForm(form)


    return {
        Enable: enableCameraCapture,
        Disable: disableCameraCapture,
        ObjectID: object_id,
        FormName: formName,
        IsClosed: isCameraClosed,
        Canvas: canvas,
        Video:video
    };
}