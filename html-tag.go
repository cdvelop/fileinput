package fileinput

func (f File) HtmlTag(id, field_name string, allow_skip_completed bool) string {

	var tags = `<div class="container-files">`

	// tags += `<input id="` + id + `" type="file" id="avatar" name="` + field_name + `" accept="image/,.png,.jpeg,.jpg" />`

	tags += `<div class="col-img">
	<div class="checkbox-container">
	  <label for="1697817735064087200">
		<img src="file?id=1697817735064087200" class="img-list">
		<input type="checkbox" id="1697817735064087200">
	  </label>
	</div>
  </div>`

	tags += `</div>`
	return tags
}
