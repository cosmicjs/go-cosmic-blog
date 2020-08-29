package handlers

import (
	"encoding/json"
	"go-cosmic-blog/src/domain"
	"go-cosmic-blog/src/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

// Index handler to serve / route
func Index(w http.ResponseWriter, r *http.Request) {

	if ok := utils.CheckIfEnvExists("BUCKET_SLUG"); !ok {
		http.Error(w, "BUCKET_SLUG is not present in the .env", http.StatusInternalServerError)
		return
	}

	var readKey string
	if ok := utils.CheckIfEnvExists("READ_KEY"); ok {
		readKey = "&read_key=" + os.Getenv("READ_KEY")
	}

	bucketSlug := os.Getenv("BUCKET_SLUG")

	url := utils.APIURL + bucketSlug + "/objects?&hide_metafields=true&type=posts&props=slug,title,content,metadata,created_at" + readKey

	res, err := http.Get(url)
	var data domain.ObjectsData

	if err != nil {
		log.Println(err)
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
		} else {
			json.Unmarshal(body, &data)
		}
	}

	t, _ := template.ParseFiles(
		"src/templates/index.html",
		"src/templates/head.html",
		"src/templates/header.html",
		"src/templates/footer.html",
	)

	t.Execute(w, data)
}
