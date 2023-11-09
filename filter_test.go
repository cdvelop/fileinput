package fileinput_test

import (
	"testing"

	"github.com/cdvelop/fileinput"
)

func TestFilterDataBy(t *testing.T) {
	data := []map[string]string{
		{"extension": ".pdf"},
		{"extension": ".jpg"},
		{"extension": ".mp4"},
		{"extension": ".docx"},
		{"extension": ".txt"},
	}

	// Prueba con criterio "pdf" (add = true)
	filteredPDF := fileinput.FilterDataBy(true, "pdf", data...)

	// Comprueba que solo se retienen los archivos PDF
	for _, item := range filteredPDF {
		if item["extension"] != ".pdf" {
			t.Errorf("La extensión no es PDF: %s", item["extension"])
		}
	}

	// Prueba con criterio "imagen" (add = true)
	filteredImage := fileinput.FilterDataBy(true, "imagen", data...)

	// Comprueba que solo se retienen las imágenes
	for _, item := range filteredImage {
		if item["extension"] != ".jpg" && item["extension"] != ".jpeg" && item["extension"] != ".png" {
			t.Errorf("La extensión no es una imagen: %s", item["extension"])
		}
	}

	// Prueba con criterio "video" (add = true)
	filteredVideo := fileinput.FilterDataBy(true, "video", data...)

	// Comprueba que solo se retienen los videos
	for _, item := range filteredVideo {
		if item["extension"] != ".avi" && item["extension"] != ".mkv" && item["extension"] != ".mp4" {
			t.Errorf("La extensión no es un video: %s", item["extension"])
		}
	}

	// Prueba con criterio "pdf" (add = false)
	filteredNonPDF := fileinput.FilterDataBy(false, "pdf", data...)

	// Comprueba que se eliminan los archivos PDF
	for _, item := range filteredNonPDF {
		if item["extension"] == ".pdf" {
			t.Errorf("Se esperaba que se eliminara el archivo PDF: %s", item["extension"])
		}
	}
}
