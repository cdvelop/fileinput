function takePicture(e) {
  console.log("tomando foto:")
  if (!camera.IsClosed()) {
    // console.log("object:", camera.ObjectName)

    const context = camera.Canvas.getContext("2d");

    context.drawImage(camera.Video, 0, 0, camera.Canvas.width, camera.Canvas.height)

    camera.Canvas.toBlob(function (blob) {

      // const url = URL.createObjectURL(blob);

      saveBlobFile(camera.ObjectName, blob);

      camera.ShowNewPicture();

    }, "image/jpeg", 0.9);
  }
}

