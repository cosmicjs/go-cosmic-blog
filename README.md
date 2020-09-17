# Go + Cosmic
![Image](https://imgix.cosmicjs.com/5de4e840-ec6b-11ea-a61c-9d54bbbbeb4a-go-cosmic-blog.png?w=1200&auto=format)

> This repo contains an example blog starter that is built with [Go](https://golang.org/), and [Cosmic](https://www.cosmicjs.com).

> [See live demo hosted on Heroku](https://go-cosmic-blog.herokuapp.com/)

## Prerequisites

- Go (I recommend using v1.15 or higher)

## Install

``` bash
go mod download
go get github.com/cespare/reflex
npm install
```

## Configuration
In `.env` you need to add configuration for your Cosmic Bucket

``` bash
# copy .env.example to .env
cp .env.example .env
```

add `BUCKET_SLUG` & `READ_KEY` if required

## Run

``` bash
# Then you can run it by
# Note: this command will only run on OSX and Linux. As reflex utility is only compatible on these platforms.
# On windows, you can run the application using `go run app.go` 
npm start
```
> Or you can build your application using
```bash
npm run build
```

then navigate to [http://localhost:8080](http://localhost:8080) to access your application
