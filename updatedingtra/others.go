package updatedingtra


import (
	"net/http"
	"strings"
	"html/template"
)

type UpdateStruct struct {
	Username string
	Fullname string
	Website string
	Work string
	Location string
	Bio string
	Details template.HTML
	Usersid string
	JsonVal string
	Urls []string
}


func (let *UpdateStruct ) Route (r *http.Request ) {

	url := strings.Split(r.URL.Path[len("/xklop/xr09/"):], "/")

	for _, k := range url {

		if len(k) > 0 {

			let.Urls = append(let.Urls, k)
		}
	}

}