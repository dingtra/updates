package updatedingtra


import (
	"net/http"
	"strings"
	// "fmt"
)

var (
	img = []string{".jpeg", ".png", ".gif", ".jpg", "webp"}
)


func CheckSingleImage (r *http.Request) (bool) {
	r.ParseMultipartForm(10 << 20)

	_, handle, err := r.FormFile("file")

	var check bool

	if err == nil {

		for _, k := range img {

			if strings.HasSuffix(handle.Filename, k) == true {
				check = true
			}
		}
	}

	return check
}