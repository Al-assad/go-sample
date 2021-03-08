package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// @desc go http rest server 示例 (httprouter)

func main() {
	router := httprouter.New()
	router.GET("/user/:uid", getUserHandler)
	router.POST("/adduser/:uid", addUserHandler)
	router.DELETE("/deluser/:uid", deleteUserHandler)
	router.PUT("/modifyuser/:uid", modifyUserHandler)
	http.ListenAndServe(":8080", router)
}

func getUserHandler(reps http.ResponseWriter, req *http.Request, param httprouter.Params) {
	uid := param.ByName("uid")
	reps.Write([]byte("get user " + uid))
}

func addUserHandler(reps http.ResponseWriter, req *http.Request, param httprouter.Params) {
	uid := param.ByName("uid")
	reps.Write([]byte("add user " + uid))
}

func deleteUserHandler(reps http.ResponseWriter, req *http.Request, param httprouter.Params) {
	uid := param.ByName("uid")
	reps.Write([]byte("delete user " + uid))
}

func modifyUserHandler(reps http.ResponseWriter, req *http.Request, param httprouter.Params) {
	uid := param.ByName("uid")
	reps.Write([]byte("modify user " + uid))
}
