function takePicture(e) {

  console.log("tomando foto:")
  console.log("object:", camera.FormName())

  const context = camera.Canvas.getContext("2d");

  context.drawImage(camera.Video, 0, 0, camera.Canvas.width, camera.Canvas.height)

  camera.Canvas.toBlob(function (blob) {


    addBlobFileInObject(blob, camera.ObjectID);
    // Puedes hacer lo que necesites con el Blob, como enviarlo a un servidor o mostrarlo en la página.
  }, "image/jpeg", 0.9);

}

