package branch


import (
	"net/http"
	"rundb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	"regexp"
	"encoding/json"
	"time"
	"strconv"
	"strings"
	// "fmt"
)


func (let *BranchStruct) CreateBranch (r *http.Request ) {
	getcoll := rundb.Connect("app").Collection("branch")

	Error := map[string]template.HTML{}
	var visible string

	if strings.ToLower(r.Method) == "post" {
		t := time.Now()
		dateval :=  t.Month().String()[:3]+" "+strconv.Itoa(t.Day())+", "+strconv.Itoa(t.Year())

		name := strings.ToLower(strings.TrimSpace(r.FormValue("name")))
		about := strings.ToLower(strings.TrimSpace(r.FormValue("about")))
		visibility := strings.ToLower(strings.TrimSpace(r.FormValue("visibility")))

		mainresult := CheckName(name)
		theid, _:= primitive.ObjectIDFromHex(let.Usersid)
		findbranchexist := rundb.FindOne(getcoll, bson.M{"name":mainresult, "usersid":theid})

		if len(visibility) > 0 {
			visible = visibility
		}else{
			visible = "public"
		}

		if len(findbranchexist) > 0 {
			Error["resname-909sk"] = "<div class='err-brnch'><div style='padding:5px 0px 5px 0px;'><span style='font-size:13px; color:red;'> Branch name exists already</span></div></div>"
		}

		if len(mainresult)  > 30 {
			Error["resname-909sk"] = "<div class='err-brnch'><div style='padding:5px 0px 5px 0px;'><span style='font-size:13px; color:red;'> Branch name can't be more than 30 characters </span></div></div>"
		}else {
			if len(mainresult) == 0 {
				Error["resname-909sk"] = "<div class='err-brnch'><div style='padding:5px 0px 5px 0px;'><span style='font-size:13px; color:red;'> Branch name can't be empty</span></div></div>"
			}
		}

		if len(name) < 10 {
			Error["txterr890"] = `
			    <div style="padding:5px 0px 5px 0px;" class="txterr090">
				    <div style="background:#f2f2f2;font-size:14px;color:red;padding:10px;border-radius:4px;" class="txterr091">
					    <span>Add a valid description, atleast 10 characters</span>
					</div>
				</div>

				<script>
				setTimeout(function(){
					$("#txterr890").html("");
				},5000)
				</script>
			`
		}
		if len(Error) == 0 {
			rundb.InsertOne(getcoll, bson.M{"name": mainresult, "about": about, "usersid":theid, "visibility":visible, "date":dateval}).(primitive.ObjectID).Hex()
			Error["br909resxy"] = template.HTML("<div class='branch-created-class'><div class='branch-created-class-ch'><button><div id='branch-created-class01'><span>Branch Created Succesfully</span></div> <div id='branch-created-class02'><span> navigate to <a href='/"+let.Username+"/"+mainresult+"'>branch page</a> to Learn more.</span></div></button></div></div>")
		}else{
			Error["br909resxy"] =""
		}

		jsn, _ := json.Marshal(Error)
		
		let.JsonVal = string(jsn)
	}
}


func CheckName (name string) string {
	mainresult := []string{}
	checkname := regexp.MustCompile(`[^A-Za-z0-9]`)
	resultname := checkname.ReplaceAllString(name, "-")

	for _, item := range strings.Split(resultname, "-") {

		if len(strings.TrimSpace(item)) > 0 {
			mainresult = append(mainresult, item)
		}
	}

	return strings.Join(mainresult, "-")
		
}