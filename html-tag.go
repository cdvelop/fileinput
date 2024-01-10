package fileinput

func (f FileInput) BuildContainerView(id, field_name string, allow_skip_completed bool) string {

	var disabled string
	if !f.conf.DefaultEnableInput {
		disabled = ` disabled`
	}

	tags := `<div class="file-input-container" data-width="` + f.conf.ImagenWidth + `" data-height="` + f.conf.ImagenHeight + `">`
	tags += `<span class="modal-file-viewer">
	<button id="` + id + `" type="button" name="capture" onclick="openCapture(this)" title="iniciar captura de imagen"` + disabled + `><svg aria-hidden="true" focusable="false" class="form-btn"><use xlink:href="#icon-camera"/></svg></button>`

	tags += `<button type="button" name="joystick" title="joystick desconectado"` + disabled + `><svg aria-hidden="true" focusable="false" class="form-btn"><use xlink:href="#icon-joystick"/></svg></button>`

	tags += `<div class="media-container media-hidden"></div>`

	// tags += `<div class="output"><img id="photo_capture" alt="The screen capture in this box."/></div>`

	tags += `</span>`

	var tabindex string
	long := len(id)
	if long >= 2 {
		tabindex = id[long-1:]
	}

	tags += `<fieldset tabindex="` + tabindex + `" class="file border"` + disabled + `>
	<legend class="basic-legend"><label for="` + id + `">` + f.conf.Legend + `</label></legend>`

	// tags += `<button type="button" name="previous_img" onclick="moveScrollFileImg(this)"><i class="arrow left"></i></button>`
	tags += `<div data-id="` + f.Object.ObjectName + `" name="` + field_name + `" class="scroll-file_img" onmousedown="targetFileHandler(event)" ontouchstart="targetFileHandler(event)">`

	// tags += `<div data-id="` + f.Object.Name + `" name="` + field_name + `" class="scroll-file_img" onclick="imgFileSelected(event)">`

	// tags += `<div name="file_img"><img src="file?id=1697817735064087200"></div>`
	// tags += `<div name="file_img"><img src="file?id=1697817735066634900"></div>`
	// tags += `<div name="file_img"><img src="file?id=1"></div>`
	// tags += `<div name="file_img"><img src="file?id=2"></div>`
	// tags += `<div name="file_img"><img src="file?id=3"></div>`

	// extensions := strings.Join(f.conf.AllowedExtensions, ",")
	// tags += `<input type="file" accept=".` + extensions + `">`

	tags += `</div>`

	// tags += `<button type="button" name="next_img" onclick="moveScrollFileImg(this)"><i class="arrow right"></i></button>`

	tags += `</fieldset>`

	tags += `<canvas name="canvas" style="display: none;"> </canvas>`

	tags += `</div>`

	return tags
}

func (f *FileInput) BuildItemsView(all_data ...map[string]string) (html string) {

	// f.Log("BuildItemsView:", all_data)

	for _, data := range all_data {

		if id, ok := data["id_file"]; ok {

			src := `file?id=` + id
			if url, exist := data["url"]; exist {
				// f.App.Log("url file:", url)
				src = url
			}

			// open
			if typeImagen(data) {
				html += `<div name="file_img" data-id="` + id + `">`
				html += `<div class="delete-tab"></div>`
				html += `<img src="` + src + `">`
			} else if typePdf(data) {
				html += `<div name="file_doc" data-id="` + id + `">`
				html += `<iframe src="` + src + `"></iframe>`

			} else if typeVideo(data) {
				html += `<video src="` + src + `"></div>`

			}

			html += `</div>` // close
		}
	}

	return
}
