# Go + Cosmic

> This repo contains an example blog starter that is built with [Go](https://golang.org/), and [Cosmic](https://www.cosmicjs.com).

## Prerequisites

- Go (I recommend using v1.15 or higher)

## Install

``` bash
go mod download
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
go run app.go
```

then navigate to [https://localhost:8080](https://localhost:8080) to access your application
