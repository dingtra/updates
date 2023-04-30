package branch


import (
	"net/http"
	"rundb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)


func (let *BranchStruct) CheckAndVerifyBranch ( r *http.Request, w http.ResponseWriter) {
	session, _ := rundb.Store.Get(r, "session")
	getcoll := rundb.Connect("app").Collection("branch")
	
	FindAll := rundb.FindAll(getcoll, bson.M{})
	var getbranch string

	if len(FindAll) > 0 {
		for _, item := range FindAll {
			if item["usersid"].(primitive.ObjectID).Hex() == let.Usersid {
				getbranch +="#"+item["name"].(string)
			}
		}
	}

	if len(getbranch) > 0 {
		session.Values["branchnames"] = getbranch
		session.Save(r, w)
	}else{
		session.Values["branchnames"] = ""
		session.Save(r, w)
	}
}
