function takePicture(e) {
  if (!camera.IsClosed()) {
    // console.log("tomando foto:")
    // console.log("object:", camera.FormName())

    const context = camera.Canvas.getContext("2d");

    context.drawImage(camera.Video, 0, 0, camera.Canvas.width, camera.Canvas.height)

    camera.Canvas.toBlob(function (blob) {

      // const url = URL.createObjectURL(blob);

      saveBlobFile(camera.ObjectName, camera.ObjectID, blob);

      camera.ShowNewPicture();

    }, "image/jpeg", 0.9);
  }
}

