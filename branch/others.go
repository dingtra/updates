package branch


import (
	"net/http"
	"rundb"
	"strings"
	"html/template"
	"encoding/json"
	"regexp"
	// "fmt"
)


type BranchStruct struct {
	Urls []string
	Details template.HTML
	JsonVal string
	Usersid string
	Username string
}



func (d *BranchStruct ) Route(r *http.Request) {

	url := strings.Split(r.URL.Path[len("/x910/nwbrnch/"):], "/")

	for _, k := range url {

		if k != ""{
			d.Urls = append(d.Urls, strings.ToLower(k))
		}
	}

}


func (let *BranchStruct ) CheckBrnachName ( r *http.Request)  {

	session, _ := rundb.Store.Get(r, "session")
	var checkbranch string
	Error := map[string]template.HTML{}


	if strings.ToLower(r.Method) == "post" {
		key := strings.ToLower(strings.TrimSpace(r.FormValue("data")))

		if len(key) > 0 {
			if session.Values["branchnames"] != nil {
				checkbranch = session.Values["branchnames"].(string)
			}

			checkname := regexp.MustCompile(`[^A-Za-z0-9]`)

			resultname := checkname.ReplaceAllString(key, "-")

			mainresult := []string{}

			for _, item := range strings.Split(resultname, "-") {

				if len(strings.TrimSpace(item)) > 0 {
					mainresult = append(mainresult, item)
				}
			}

			if len(checkbranch) > 0 {
				moldbranch := make(map[string]bool)

				br := strings.Split(checkbranch, "#")

				for _, item := range br {
					if len(strings.TrimSpace(item)) > 0 {
						moldbranch[item] = true
					}
				}

				if moldbranch[strings.Join(mainresult, "-")] == true {
					Error["#resname-909sk"] =template.HTML("<div class='err-show-div'><div class='err-show-div-ch'><span style='color:red;'>Branch name <b>"+strings.Join(mainresult, "-")+"</b> exists already</span></div></div>")
					Error["#fnd-908res"] = "<svg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24' fill='none' stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round' class='feather feather-x' style='color:red;height:13px;width:13px;' color='black'><line x1='18' y1='6' x2='6' y2='18'></line><line x1='6' y1='6' x2='18' y2='18'></line></svg>"
				}else{
					Error["#resname-909sk"] =template.HTML("<div class='err-show-div'><div class='err-show-div-ch'><span>Branch name <b>"+strings.Join(mainresult, "-")+"</b> available</span></div></div>")
					Error["#fnd-908res"] = "<svg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24' fill='none' stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round' class='feather feather-check' style='color:green;height:13px;width:13px;' color='black'><polyline points='20 6 9 17 4 12'></polyline></svg>"
				}
			}else{
				Error["#resname-909sk"] =template.HTML("<div class='err-show-div'><div class='err-show-div-ch'><span>Branch name <b>"+strings.Join(mainresult, "-")+"</b> available</span></div></div>")
				Error["#fnd-908res"] = "<svg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24' fill='none' stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round' class='feather feather-check' style='color:green;height:13px;width:13px;' color='black'><polyline points='20 6 9 17 4 12'></polyline></svg>"
					
			}
			
	
			jsn, _ := json.Marshal(Error)
			
			let.JsonVal = string(jsn)

		}else{
			Error["#resname-909sk"] = ""
		}
		

	}
}
