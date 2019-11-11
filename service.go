package service

import (
    "github.com/go-martini/martini" 
)

func NewServer(port string) {   
    app := martini.Classic()
    app.Get("/hello/:name", func(params martini.Params) string {
        return "Hello " + params["name"] + " !"
    })
    app.RunOnAddr(":"+port)   
}