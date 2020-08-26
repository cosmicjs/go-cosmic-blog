package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"os"
)

// Data is a array of objects from Cosmic API
type Data struct {
	Objects []Post
}

// Post is a representation of post object
type Post struct {
	Title    string
	Slug     string
	Content  template.HTML
	Metadata Metadata
}

// Metadata is a representation of metadata object
type Metadata struct {
	Hero Image
}

// Image is a object of URL & ImgixURL
type Image struct {
	URL      string
	ImgixURL string `json:"imgix_url"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
  }
  
  if ok := checkIfEnvExists("BUCKET_SLUG"); !ok {
    http.Error(w, "BUCKET_SLUG is not present in the .env", http.StatusInternalServerError)
		return
  }

  var readKey string
  if ok := checkIfEnvExists("READ_KEY"); ok {
    readKey = "&read_key=" + os.Getenv("READ_KEY")
  }

  bucketSlug := os.Getenv("BUCKET_SLUG")

  apiURL := "https://api.cosmicjs.com/v1/"
	url := apiURL + bucketSlug + "/objects?&hide_metafields=true&type=posts&props=slug,title,content,metadata" + readKey

  res, err := http.Get(url)
  var data Data

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

	t, _ := template.ParseFiles("index.html")

	t.Execute(w, data)
}

func getPortEnv() string {
  var port string
	var ok bool

	if port, ok = os.LookupEnv("PORT"); !ok {
		port = "8080"
  }
  
  return ":" + port
}

func checkIfEnvExists(key string) bool {
	var ok bool
	if _, ok = os.LookupEnv(key); !ok {
		return false
  }
  return true
}

func main() {
	http.HandleFunc("/", indexHandler)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
  }

	port := getPortEnv()

	fmt.Println("Starting server at port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
