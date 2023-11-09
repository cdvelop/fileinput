package fileinput

//criterion ej: video,imagen,pdf
//add true, false remove only
func FilterDataBy(add bool, criterion string, data_in ...map[string]string) (data_out []map[string]string) {

	var filter []string

	switch criterion {
	case "pdf":
		filter = append(filter, ".pdf")

	case "imagen":
		filter = append(filter, ".jpg", ".png", ".jpeg")

	case "video":
		filter = append(filter, ".avi", ".mkv", ".mp4")

	}

	for _, data := range data_in {
		extension := data["extension"]
		matchesFilter := false

		for _, ext := range filter {
			if extension == ext {
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
