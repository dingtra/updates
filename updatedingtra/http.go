package updatedingtra


import (
	"net/http"
	"rundb"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
)


func Http (w http.ResponseWriter, r *http.Request ){
	session, _ := rundb.Store.Get(r, "session")
	w.Header().Set("Content-Type", "text/html")
	let := UpdateStruct{}
	let.Route(r)

	if session.Values["username"] != nil && session.Values["usersid"] != nil {
		collection := rundb.Connect("app").Collection("register")

		if len(let.Urls) == 0 {
			finduser := rundb.FindOne(collection, bson.M{"username":session.Values["username"].(string)})
			let.ProfileForm(finduser)
			fmt.Fprintf(w, "%s",let.Details)
		}else{

			if let.Urls[0] == "xkr090" {
				let.UpdateUserProfile(w,r)
				fmt.Fprintf(w, "%s",let.JsonVal)

			}else if let.Urls[0] == "xkr091" {
				let.RemoveImage(w, r)
				fmt.Fprintf(w, "%s",let.JsonVal)
			}
		}

	}

}