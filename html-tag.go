package fileinput

func (f File) HtmlTag(id, field_name string, allow_skip_completed bool) string {

	tags := `<div class="container-files">`
	tags += `<span class="modal-capture" onclick="openCapture(event)">
	<legend class="basic-legend"><svg aria-hidden="true" focusable="false" class="form-btn">
	<use xlink:href="#icon-camera" /></svg></legend></span>`
	tags += `<button class="btn-prev-img" onclick="prevButtonFileImg(event)"><i class="arrow left"></i></button>`

	tags += `<div class="scroll-file-img" onclick="imgFileSelected(event)">`

	tags += `<img src="file?id=1697817735064087200" class="file-img">`
	tags += `<img src="file?id=1697817735066634900" class="file-img">`
	tags += `<img src="file?id=1" class="file-img">`
	tags += `<img src="file?id=2" class="file-img">`
	tags += `<img src="file?id=3" class="file-img">`

	tags += `</div>`
	tags += `<button class="btn-next-img" onclick="nextButtonFileImg(event)"><i class="arrow right"></i></button>`
	tags += `</div>`

	return tags
}
