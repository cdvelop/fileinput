package fileinput

func (f File) HtmlTag(id, field_name string, allow_skip_completed bool) string {

	tags := `<div class="container-files">`
	tags += `<span class="modal-capture"><button type="button" name="capture" onclick="openCapture(this)"><svg aria-hidden="true" focusable="false" class="form-btn"><use xlink:href="#icon-camera"/></svg></button></span>`

	tags += `<button type="button" name="previous_img" onclick="moveScrollFileImg(this)"><i class="arrow left"></i></button>`
	tags += `<div class="scroll-file_img" onclick="imgFileSelected(event)">`

	tags += `<div name="file_img"><img src="file?id=1697817735064087200"></div>`
	tags += `<div name="file_img"><img src="file?id=1697817735066634900"></div>`
	tags += `<div name="file_img"><img src="file?id=1"></div>`
	tags += `<div name="file_img"><img src="file?id=2"></div>`
	tags += `<div name="file_img"><img src="file?id=3"></div>`
	tags += `</div>`

	tags += `<button type="button" name="next_img" onclick="moveScrollFileImg(this)"><i class="arrow right"></i></button>`

	tags += `</div>`

	return tags
}
