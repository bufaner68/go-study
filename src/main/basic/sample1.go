package main

import "github.com/emicklei/go-restful"

func main() {
  ws := new(restful.WebService)
  ws.Path("/users")
  ws.Route(ws.GET("/").To(u.findAllUsers).
    Doc("get all users").
    Metadata(restfulspec.KeyOpenAPITags, tags).
    Writes([]User{}).
    Returns(200, "OK", []User{}))
 
 container := restful.NewContainer().Add(ws)
 http.ListenAndServe(":8080", container)
}