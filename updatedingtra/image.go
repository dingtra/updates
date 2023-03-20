package updatedingtra



import (
	"net/http"
	"strings"
	"context"
    "github.com/cloudinary/cloudinary-go"
    "github.com/cloudinary/cloudinary-go/api/uploader"
	"io/ioutil"
	"log"
	"fmt"
)


func CloudSingleImage(r *http.Request) string {
	var ctx = context.Background()
	cld, _ := cloudinary.NewFromParams("dings", "558969314732393", "GLhq9JWpWaxSOq7raiX01b-p6IM")

	r.ParseMultipartForm( 10 << 20 )

	file, handler, err := r.FormFile("file")

    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
    }

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    defer file.Close()

	tempfile, err := ioutil.TempFile("", "image-*")

    if err != nil {
        fmt.Println(err)
    }

	resp, err:= cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: "image"+Breakall(tempfile.Name())});
	
	if err != nil {
		log.Fatal("Failed to upload files %v\n", err)
	}

	return "#"+resp.SecureURL
}


func Breakall(ma string) string {
	bat := strings.Split(ma, "-")
	fresh := []string{bat[len(bat)-1]}
	return strings.Join(fresh, "")
}