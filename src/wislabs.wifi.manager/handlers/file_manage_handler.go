package handlers

import (
	"net/http"
	"os"
	"io"

	"log"
)

func UploadFilesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	//GET displays the upload form.
	case "GET":

	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			log.Printf("erro 1")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//get a ref to the parsed multipart form
		m := r.MultipartForm

		//get the *fileheaders
		files := m.File["profile-images"]
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			log.Printf("erro 2")
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Printf("erro 3")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.

			dst, err := os.Create("temp/" + "sandun.jpg")
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.Write([]byte("Profile photo successfully updates"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
