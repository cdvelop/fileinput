package fileinput

func (f File) BuildContainerView(id, field_name string, allow_skip_completed bool) string {

	var disabled string
	if !f.DefaultEnableInput {
		disabled = ` disabled`
	}

	tags := `<div class="container-files">`
	tags += `<span class="modal-capture">
	<button id="` + id + `" type="button" name="capture" onclick="openCapture(this)"` + disabled + `><svg aria-hidden="true" focusable="false" class="form-btn"><use xlink:href="#icon-camera"/></svg></button>`

	tags += `<div class="video-container video-hidden"><video data-width="` + f.ImagenWidth + `" data-height="` + f.ImagenHeight + `" title="toca el video para tomar una captura" onclick="takePicture(this)" id="video_capture">Video stream no disponible en este dispositivo.</video></div>`

	// tags += `<div class="output"><img id="photo_capture" alt="The screen capture in this box."/></div>`

	tags += `</span>`

	var tabindex string
	long := len(id)
	if long >= 2 {
		tabindex = id[long-1:]
	}

	tags += `<fieldset tabindex="` + tabindex + `" class="file border"` + disabled + `>
	<legend class="basic-legend"><label for="` + id + `">` + f.Legend + `</label></legend>`

	tags += `<button type="button" name="previous_img" onclick="moveScrollFileImg(this)"><i class="arrow left"></i></button>`
	tags += `<div data-id="` + f.Object.Name + `" name="` + field_name + `" class="scroll-file_img" >`

	// tags += `<div data-id="` + f.Object.Name + `" name="` + field_name + `" class="scroll-file_img" onclick="imgFileSelected(event)">`

	// tags += `<div name="file_img"><img src="file?id=1697817735064087200"></div>`
	// tags += `<div name="file_img"><img src="file?id=1697817735066634900"></div>`
	// tags += `<div name="file_img"><img src="file?id=1"></div>`
	// tags += `<div name="file_img"><img src="file?id=2"></div>`
	// tags += `<div name="file_img"><img src="file?id=3"></div>`
	tags += `<input type="file" accept="` + f.FileType + `/` + f.AllowedExtensions + `">`

	tags += `</div>`

	tags += `<button type="button" name="next_img" onclick="moveScrollFileImg(this)"><i class="arrow right"></i></button>`

	tags += `</fieldset>`

	tags += `<canvas name="canvas" style="display: none;"> </canvas>`
	// tags += `<input type="hidden"  oninput="changeDataFile()"/>`

	tags += `</div>`

	return tags
}

func (f File) BuildItemView(all_data ...map[string]string) (html string) {

	for _, data := range all_data {

		if id, ok := data[f.Id_file]; ok {

			src := `file?id=` + id
			if url, exist := data["url"]; exist {
				// f.App.Log("url file:", url)
				src = url
			}

			html += `<div name="file_img" onmousedown="startTimer(event)" onmouseup="clearTimer(event)">
			<div class="delete-tab"></div>
			<img src="` + src + `" data-id="` + id + `"></div>`

		}
	}

	return
}
