# Restroutes

The purpose of restroutes is twofold: to simplify the use of a restful routing system into the gin framework for Golang, and to help organize go web app programs with MVC concepts. The idea is to abstract a lot of the HTTP methods for each model and programmatically add them whenever a new resource is needed or created. We assign each HTTP method, passed in as a string, to a receiver struct that then calls the appropriate method on that route.

### Overview
Rather than have files with functions that lack naming conventions for each route function, we assign the function to the struct defined in each controller. For example, instead of:

```go
func randomGetFunction(){
  //one func
}

func randomPutFunction(){

}
.... etc etc
```
our files look more like
```go
func (r *ResourceStruct) Show(c *gin.Context) {

}

func (r *ResourceStruct) Put(c *gin.Context) {

}
```
We never have to guess anymore what route is mapped to what resource, or what the restful function is supposed to do.

### How it works

Let's start with a basic web application using [github.com/gin-gonic/gin](Gin):

```go
package main

import (
	"github.com/gin-gonic/gin"
	rr "github.com/jeisenberg/restroutes"
	c "path/to/controllers/file"
)

// we create the Routes var to tell our application what routes we want mapped to what controller functions
// for example, lets say we have projects and posts

var Routes = rr.Routes{
	"/api/projects/:id": rr.Route{&c.Projects{}, "Show", "GET"},
	"/api/projects/new": rr.Route{&c.Projects{}, "Create", "POST"},
	"/api/projects":     rr.Route{&c.Projects{}, "Index", "GET"},
	"/api/posts/new":    rr.Route{&c.Posts{}, "Create", "POST"},
	"/api/posts/:id":    rr.Route{&c.Posts{}, "Show", "GET"},
}

func main(){
  r := gin.Default()
  rr.RegisterGin(r, Routes)
}
```

This will automatically tell your gin application to route `/api/projects/:id` to the func defined in your controllers directory:

```go
  type Projects struct {
    EmbeddedSuperStruct
  }

  func (p *Projects) Show(c * gin.Context) {
    //do controller magic here
  }
```
