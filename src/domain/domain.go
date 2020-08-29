package domain

import (
	"html/template"
	"time"
)

// ObjectsData is a array of objects from Cosmic API
type ObjectsData struct {
	Objects []Post
}

// ObjectData is a object from Cosmic API
type ObjectData struct {
	Object Post
}

// Post is a representation of post object
type Post struct {
	Title     string
	Slug      string
	Content   template.HTML
	CreatedAt time.Time `json:"created_at"`
	Metadata  Metadata
}

// Metadata is a representation of metadata object
type Metadata struct {
	Hero        Image
	Description string
}

// Image is a object of URL & ImgixURL
type Image struct {
	URL      string
	ImgixURL string `json:"imgix_url"`
}
