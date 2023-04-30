package branch

import (
	"net/http"
	"rundb"
	"strings"
	"rendertemplates"
	"fmt"
)


func Http (w http.ResponseWriter, r *http.Request)  {
	session, _ := rundb.Store.Get(r, "session")
	w.Header().Set("Content-Type", "text/html")
	let := BranchStruct{}
	let.Route(r)

	if session.Values["username"] != nil {
		let.Username =session.Values["username"].(string)
		let.Usersid = session.Values["usersid"].(string)
		if len(let.Urls) > 0 {
			if let.Urls[0] == "createbrnch" {
				if strings.ToLower(r.Method) == "post" {
					let.CheckAndVerifyBranch(r, w)
					let.CreateBranch(r)
					fmt.Fprintf(w, " %s", let.JsonVal)
				}else{
					let.HTML(r)

					fmt.Fprintf(w, rendertemplates.HTML(let, "http://localhost:7070/dingtra-web-assets/templates/branch/branch.html"))
				}

			}else if let.Urls[0] == "srch" {
				let.CheckBrnachName(r)
				fmt.Fprintf(w, " %s", let.JsonVal)

			}else if let.Urls[0] == "slct"{
				if strings.ToLower(r.Method) == "post" {
					data := strings.ToLower(strings.TrimSpace(r.FormValue("data")))
					fmt.Fprintf(w, " %s",AboutBranchLink(r, data))	
				}
			}
		}
	}else {
		http.Redirect(w, r, "/app/login/", http.StatusSeeOther)
	}

}