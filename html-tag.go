package fileinput

import "github.com/cdvelop/strings"

func (f File) HtmlTag(id, field_name string, allow_skip_completed bool) string {

	tags := `<div class="container-files">`
	tags += `<span class="modal-capture">
	<button id="` + id + `" type="button" name="capture" onclick="openCapture(this)"><svg aria-hidden="true" focusable="false" class="form-btn"><use xlink:href="#icon-camera"/></svg></button>`

	tags += `<div class="video-container"><video title="toca el video para tomar una captura" onclick="takePicture(this)" id="video_capture">Video stream no disponible en este dispositivo.</video></div>`

	tags += `<div class="output"><img id="photo_capture" alt="The screen capture in this box."/></div>`

	tags += `</span>`

	long := len(id)
	if long >= 2 {
		f.TabIndexNumber = id[long-1:]
	}

	tags += `<fieldset tabindex="` + f.TabIndexNumber + `" class="file border">
	<legend class="basic-legend"><label for="` + id + `">` + f.Legend + `</label></legend>`

	tags += `<button type="button" name="previous_img" onclick="moveScrollFileImg(this)"><i class="arrow left"></i></button>`
	tags += `<div name="` + field_name + `" class="scroll-file_img" onclick="imgFileSelected(event)">`

	// tags += `<div name="file_img"><img src="file?id=1697817735064087200"></div>`
	// tags += `<div name="file_img"><img src="file?id=1697817735066634900"></div>`
	// tags += `<div name="file_img"><img src="file?id=1"></div>`
	// tags += `<div name="file_img"><img src="file?id=2"></div>`
	// tags += `<div name="file_img"><img src="file?id=3"></div>`
	tags += `<input type="file" accept="` + f.FileType + `/` + f.AllowedExtensions + `">`

	tags += `</div>`

	tags += `<button type="button" name="next_img" onclick="moveScrollFileImg(this)"><i class="arrow right"></i></button>`

	tags += `</fieldset>`
	// tags += `<input type="hidden"  oninput="changeDataFile()"/>`
	tags += `</div>`

	return tags
}

func buildTag(data []string) (html string) {
	for _, id := range data {
		if id != "" {
			html += `<div name="file_img"><img src="file?id=` + id + `"></div>`
		}
	}
	return
}

func (f File) BuildNewView(values string) (html string) {
	data := strings.Split(values, ",")
	return buildTag(data)
}
