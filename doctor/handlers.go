package main

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Doctor struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Speciality string `json:"speciality"`
}

var doctors = []Doctor{
	{ID: "1", Name: "John", Speciality: "Heart"},
	{ID:"2", Name: "Lia", Speciality: "Eyes"},
}

func CreateDoctor(){

}