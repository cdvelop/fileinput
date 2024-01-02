package fileinput

var (
	filter_pdf   = []string{".pdf"}
	filter_img   = []string{".jpg", ".png", ".jpeg"}
	filter_video = []string{".avi", ".mkv", ".mp4"}
)

// criterion ej: video,imagen,pdf
// add true, false remove only
func FilterDataBy(add bool, criterion string, data_in ...map[string]string) (data_out []map[string]string) {

	var filter []string

	switch criterion {
	case "pdf":
		filter = filter_pdf

	case "imagen":
		filter = filter_img

	case "video":
		filter = filter_video

	}

	for _, data := range data_in {
		matchesFilter := false

		for _, ext := range filter {
			if data["extension"] == ext {
				matchesFilter = true
				break
			}
		}

		if (add && matchesFilter) || (!add && !matchesFilter) {
			data_out = append(data_out, data)
		}
	}

	return
}

func typeImagen(data map[string]string) bool {
	for _, filter := range filter_img {
		if data["extension"] == filter {
			return true
		}
	}
	return false
}
func typePdf(data map[string]string) bool {
	for _, filter := range filter_pdf {
		if data["extension"] == filter {
			return true
		}
	}
	return false
}
func typeVideo(data map[string]string) bool {
	for _, filter := range filter_video {
		if data["extension"] == filter {
			return true
		}
	}
	return false
}
