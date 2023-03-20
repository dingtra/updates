module github.com/dingtra/updates/updatedingtra

go 1.19

replace hashtag => ../hashtagandtags/hashtag

require (
	github.com/cloudinary/cloudinary-go v1.7.0
	go.mongodb.org/mongo-driver v1.11.3
	hashtag v0.0.0-00010101000000-000000000000
)

require (
	github.com/creasty/defaults v1.5.1 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/schema v1.2.0 // indirect
)
