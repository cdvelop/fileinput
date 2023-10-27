function takePicture(e) {
    

    console.log("tomando foto:",e)
}

function captureImage(canvasSelector, video, canvasWidth = 800, canvasHeight = 600) {
    const canvas = document.querySelector(canvasSelector);
    canvas.style.display = 'none';
  
    const context = canvas.getContext("2d");
  
    // Aumentar la resolución del canvas
    canvas.width = canvasWidth;
    canvas.height = canvasHeight;
  
    context.drawImage(video, 0, 0, canvas.width, canvas.height);
  
    return new Promise((resolve, reject) => {
      canvas.toBlob((blob) => {
        if (blob) {
          resolve(blob);
        } else {
          reject("Error al capturar la imagen.");
        }
      }, "image/jpeg", 0.9);
    });
  }