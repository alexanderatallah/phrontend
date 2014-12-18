// +build !appengine
// Above is a special build command: https://blog.golang.org/the-app-engine-sdk-and-workspaces-gopath

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func main() {
	// API Subroute
	api := web.New()
	goji.Handle("/api/*", api)
	api.Use(middleware.SubRouter)
	api.Get("/posts", RoutePosts)

	// Static assets
	goji.Get("/*", RouteStatic)

	// goji.NotFound(NotFound)
	goji.Serve()
}

func RouteStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PATH REQUEST:", r.URL.Path)
	http.ServeFile(w, r, "static/"+r.URL.Path)
}

func RoutePosts(w http.ResponseWriter, r *http.Request) {
	posts := []Post{}
	file, err := ioutil.ReadFile("./db/posts.json")
	PanicIf(err)

	fmt.Printf("db file: %s\n", string(file))
	err = json.Unmarshal(file, &posts)
	CheckErr(err, "Error unmarshalling posts.json:")

	bs, err := json.Marshal(posts)
	if err != nil {
		ReturnError(w, err)
		return
	}
	fmt.Fprint(w, string(bs))
}
