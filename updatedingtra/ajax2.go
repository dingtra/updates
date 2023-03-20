package updatedingtra


import (
	"net/http"
	"rundb"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"strconv"
	"html/template"
	"encoding/json"
)


func (let *UpdateStruct) RemoveImage(w http.ResponseWriter, r *http.Request ) {
	getcoll := rundb.Connect("app").Collection("register")
	Error := map[string]template.HTML{}
	if strings.ToLower(r.Method) == "post" {
		session, _ := rundb.Store.Get(r, "session")

		if session.Values["pic"] != nil && len(session.Values["pic"].(string)) > 0 {
			rundb.UpdateOne(getcoll, bson.M{"pic":""}, session.Values["usersid"].(string))
			session.Values["pic"] = nil
			session.Save(r, w)

			// in this block it means image has been updated to "" => which means removed as assumed
			// the next step is to update existing images across the page 
			// the pictures id have been grabbed to make it easier as a response to the ajax call

			Error["userxr090img"] = template.HTML(`<img src='/dingtra-web-assets/svg/pr.png' />`)
			Error["headerxr090img"] = template.HTML(`<img src='/dingtra-web-assets/svg/pr.png' />`)
			Error["filexr09780"] = template.HTML(`<label class="xr09bklabel" for="imagefiler"> <img src='/dingtra-web-assets/svg/pr.png' /> </label>`)

			// update all images for branches if available
			for i, _ := range [5]string{} {

				Error["branchxr090img"+strconv.Itoa(i)] = template.HTML(`<span id="userxr090img"><img src='/dingtra-web-assets/svg/pr.png' /></span>`)
			}

			Error["errsklop90"] = template.HTML(`
				<div style="position:fixed;right:20px;top:100px;z-index:150;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1);background:#ccc;padding:10px 20px 0px 20px;border-radius:4px;width:200px;" class="errxrr090">
				    <div style="padding:5px 0px 5px 0px;" class="errxy090"><div style="border-radius:4px;padding:5px 5px 5px 5px;font-size:14px;" class="errxy090-chld">
				        <span> Detected and removed existing picture </span>
				    </div></div>
			    </div>
					
			   `)
			
			Error["imagetrsh09xr"] = template.HTML(`
			<div class="imgtrashmore">
				<div id="imgtrsh090" class="imgtrashmore090"> 
					<div style="pointer-events:none;"><span>Remove image</span></div>
				</div>
				<div id="imgtrsh091" style="border:unset;" class="imgtrashmore090"> 
					<span style="pointer-events:none;">Cancel</span>
				</div>
			</div>
				
			`)
		}else {
			Error["errsklop90"] = template.HTML(`
				<div style="position:fixed;right:20px;top:100px;z-index:150;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1);background:#ccc;padding:10px 20px 0px 20px;border-radius:4px;width:200px;" class="errxrr090">
				    <div style="padding:5px 0px 5px 0px;" class="errxy090"><div style="border-radius:4px;padding:5px 5px 5px 5px;font-size:14px;" class="errxy090-chld">
				        <span> Unable to remove an empty picture </span>
				    </div></div>
			    </div>
					
			`)

			Error["imagetrsh09xr"] = template.HTML(`
			<div class="imgtrashmore">
			    <div id="imgtrsh090" style="padding-right:20px;border-right:1px solid #ccc;" class="imgtrashmore090"> 
				    <span>Remove image</span>
			    </div>
			    <div id="imgtrsh091" style="padding-left:20px;" class="imgtrashmore090"> 
				    <span>Cancel</span>
			    </div>
		    </div>
			`)
		}
	}

	jsn, _ := json.Marshal(Error)
	let.JsonVal = string(jsn)
}