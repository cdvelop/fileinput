

function CameraVideo(form) {
    let camera_is_closed = true;
    let video_cont = form.querySelector('.video-container')
    let video = form.querySelector('#video_capture');
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
    }

    console.log("camera_is_closed:", camera_is_closed)

    function disableCameraCapture() {
        if (!camera_is_closed) {
            video.srcObject.getTracks().forEach((track) => track.stop());
            camera_is_closed = true;
            video.removeEventListener("canplay", canplayListener, false);
            // video.style.display = "none"; // Ocultar el elemento de video
            // Para ocultar el elemento de video suavemente
            video_cont.classList.add("video-hidden");
        }
    }

    function formName() {
        return form.name;
    }

    // disableCameraCapture();

    return {
        Enable: enableCameraCapture,
        Disable: disableCameraCapture,
        FormName: formName
    };
}