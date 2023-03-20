package updatedingtra


import (
	"net/http"
	"rundb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"encoding/json"
	"strings"
	"strconv"
	"html/template"
	"regexp"
	"github.com/dingtra/hashtagandtags/hashtag"
	// "fmt"
)


func (let *UpdateStruct ) UpdateUserProfile (w http.ResponseWriter, r *http.Request ) {
	session, _ := rundb.Store.Get(r, "session")
	getcoll := rundb.Connect("app").Collection("register")
	theid, _ := primitive.ObjectIDFromHex(session.Values["usersid"].(string))
	finduser := rundb.FindOne(getcoll, bson.M{"_id":theid})
	Error := map[string]template.HTML{}

	if strings.ToLower(r.Method) == "post" {

		username := strings.ToLower(strings.TrimSpace(r.FormValue("username")))
		fullname := strings.ToLower(strings.TrimSpace(r.FormValue("fullname")))
		location := strings.ToLower(strings.TrimSpace(r.FormValue("location")))
		website := strings.ToLower(strings.TrimSpace(r.FormValue("website")))
		work := strings.ToLower(strings.TrimSpace(r.FormValue("work")))
		bio := strings.TrimSpace(r.FormValue("bio"))
		graberror := []string{}

		UpdateBson := bson.M{}

		if username != session.Values["username"].(string) {
			matched, _ := regexp.MatchString("^[a-z0-9_.]+$", username)

			if len(username) > 25 {

				if len(username) > 25 {
					graberror = append(graberror, `
					  <div style="padding:10px 10px 10px 10px;" class="errxy090"><div style="background:#ccc;width:200px;border-radius:4px;padding:5px 5px 5px 5px;color:black;font-size:14px;" class="errxy090-chld">
					    <span>Username can only be 25 characters or less</span>
					  </div></div>
					`)
				}

			}else{
				if matched == false {
					graberror = append(graberror, `
					  <div style="padding:10px 10px 10px 10px;" class="errxy090"><div style="background:#ccc;width:200px;border-radius:4px;padding:5px 5px 5px 5px;color:black;font-size:14px;" class="errxy090-chld">
					    <span>Username format not supported</span>
					  </div></div>
					`)
				}else{
					checkcredentials := rundb.FindOne(getcoll, bson.M{"username":username})
		
					if len(checkcredentials) > 0 {
						graberror = append(graberror, `
					     <div style="padding:10px 10px 10px 10px;" class="errxy090"><div style="background:#ccc;width:200px;border-radius:4px;padding:5px 5px 5px 5px;color:black;font-size:14px;" class="errxy090-chld">
					        <span>Username taken already</span>
					     </div></div>
					    `)
					}else{
						UpdateBson["username"] = username
					}
				}
			}
		}

		if len(fullname) > 25 {
			graberror = append(graberror, `
			    <div style="padding:10px 10px 10px 10px;" class="errxy090"><div style="background:#ccc;width:200px;border-radius:4px;padding:5px 5px 5px 5px;color:black;font-size:14px;" class="errxy090-chld">
			       <span> Fullname can only be 30 characters or less </span>
			    </div></div>
		    `)
			
		}else{

			if finduser["fullname"] != nil && len(finduser["fullname"].(string)) > 0 {
				if finduser["fullname"].(string) != fullname{
					UpdateBson["fullname"] = fullname
				}
			}else{
				UpdateBson["fullname"] = fullname
			}
		}

		if len(location) > 45 {
			graberror = append(graberror, `
			    <div style="padding:10px 10px 10px 10px;" class="errxy090"><div style="background:#ccc;width:200px;border-radius:4px;padding:5px 5px 5px 5px;color:black;font-size:14px;" class="errxy090-chld">
			       <span> Location can only be 40 characters or less </span>
			    </div></div>
		    `)
		}else{

			if finduser["location"] != nil && len(finduser["location"].(string)) > 0 {
				if finduser["location"].(string) != location{
					UpdateBson["location"] = location
				}
			}else{
				UpdateBson["location"] = location
			}
		}

		if len(work) > 25 {
			graberror = append(graberror, `
			    <div style="padding:10px 10px 10px 10px;" class="errxy090"><div style="background:#ccc;width:200px;border-radius:4px;padding:5px 5px 5px 5px;color:black;font-size:14px;" class="errxy090-chld">
			       <span>exceeding work limit, atleast 25 characters </span>
			    </div></div>
		    `)
		}else{
			UpdateBson["work"] = work
		}

		if len(bio) > 120 {
			graberror = append(graberror, `
			    <div style="padding:10px 10px 10px 10px;" class="errxy090"><div style="background:#ccc;width:200px;border-radius:4px;padding:5px 5px 5px 5px;color:black;font-size:14px;" class="errxy090-chld">
			       <span>exceeding bio limit, atleast 120 characters </span>
			    </div></div>
		    `)
		}else{
			if finduser["about"] != nil && len(finduser["about"].(string)) > 0 {
				if finduser["about"].(string) != bio{
					UpdateBson["about"] = bio
				}
			}else{
				UpdateBson["about"] = bio
			}
		}

		if len(website) > 0 {
			if finduser["website"] != nil && len(finduser["website"].(string)) > 0 {
				if finduser["website"].(string) != website {
					UpdateBson["website"] = website
				}
			}else{
				UpdateBson["website"] = website
			}
		}

		if len(graberror) > 0 {
			if len(graberror) > 2 {
				Error["errsklop90"] = template.HTML(`
			        <div style="position:fixed;right:20px;top:100px;z-index:150;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1);border:1px solid #ccc;background:white;padding:10px 10px 0px 10px;border-radius:4px;" class="errxrr090">
			        `+strings.Join(graberror, "")+`
					  <div style="width:fit-content;margin:auto;padding:5px 0px 10px 0px"><span style="color:red;font-size:14px;">Too many errors</span></div>
			        </div>
			    `)
			}else{
				Error["errsklop90"] = template.HTML(`
				    <div style="position:fixed;right:20px;top:100px;z-index:150;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1);border:1px solid #ccc;background:white;padding:10px 10px 0px 10px;border-radius:4px;" class="errxrr090">
			          `+strings.Join(graberror, "")+`
			        </div>
			   `)
			}
		}else{
			img := CheckSingleImage(r)
			var imageresult string

			if img == true {
				imageresult = CloudSingleImage(r)
				UpdateBson["pic"] = imageresult
			}

			if len(UpdateBson) > 0 {
				// proceed
				rundb.UpdateOne(getcoll, UpdateBson, session.Values["usersid"].(string))
				session.Values["username"] = username
				var UpdateUserOnUrl string

				if UpdateBson["username"] != nil {
					UpdateUserOnUrl = `<script>window.history.pushState("object or string", "Title", "/`+username+`");</script>` 
				}

				if img == true {
					session.Values["pic"] = imageresult 
					Error["userxr090img"] = template.HTML(`<img src="`+imageresult[1:]+`" />`)
					Error["headerxr090img"] = template.HTML(`<img src="`+imageresult[1:]+`" />`)

					// we are showing first five branch in overview so to automatically update the images on branch
					// we will assume for the images to be 5

					for i, _ := range [15]string{} {

						Error["branchxr090img"+strconv.Itoa(i)] = template.HTML(`<img src="`+imageresult[1:]+`" />`)
					}
				}
				session.Save(r, w)

				Error["errsklop90"] = template.HTML(`
				    <div style="position:fixed;right:20px;top:100px;z-index:150;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1);background:#ccc;padding:10px 20px 0px 20px;border-radius:4px;width:200px;" class="errxrr090">
					    <div style="padding:5px 0px 5px 0px;" class="errxy090"><div style="border-radius:4px;padding:5px 5px 5px 5px;color:green;font-size:14px;" class="errxy090-chld">
					      <span> Updated successfully </span>
				        </div></div>
			        </div>
			   `+UpdateUserOnUrl)

			   Error["bdy"+session.Values["usersid"].(string)] = template.HTML(`<div style="padding-top:10px;" class="bdyxr091x">`+hashtag.EmbedTagsInBio(bio)+`<span></span></div>`)
			   Error["username"+session.Values["usersid"].(string)] = template.HTML("@"+username)
			   Error["fullname"+session.Values["usersid"].(string)] = template.HTML(fullname)
			   Error["website"+session.Values["usersid"].(string)] = template.HTML(`<div style="position:relative;top:2px;" class="ovx6789xl"><i data-feather='link-2' color='black'></i></div> <div class="ovx6789x"> `+website+` 
			   </div>
			   <script>
                    feather.replace()
                </script>
			   `)
			   Error["location"+session.Values["usersid"].(string)] = template.HTML(`<div class="ovx6789xl"><i data-feather='map-pin' color='black'></i></div> <div class="ovx6789x">`+location+`
			   </div>
			   <script>
                    feather.replace()
                </script>
			   `)
			}else{
				Error["errsklop90"] = template.HTML(`
				<div style="position:fixed;right:20px;top:100px;z-index:150;box-shadow: 0 3px 10px rgb(0 0 0 / 0.1);background:#ccc;padding:10px 20px 0px 20px;border-radius:4px;width:200px;" class="errxrr090">
				    <div style="padding:5px 0px 5px 0px;" class="errxy090"><div style="border-radius:4px;padding:5px 5px 5px 5px;font-size:14px;" class="errxy090-chld">
				        <span> Nothing changed here. </span>
				    </div></div>
			    </div>
					
			   `)
			}
		}
	}

	jsn, _ := json.Marshal(Error)
	let.JsonVal = string(jsn)
}