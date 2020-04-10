package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*type msg struct {
	username    string `json:"user"`
	password    string
	captcha     string
}*/
type Person struct {
	Id        int    `json:"id" form:"id"`
	Module string `json:"module" form:"module"`
	Content  string `json:"content" form:"content"`
}
func Login(this *gin.Context)  {
	persons := make([]Person, 0)
	var person Person
	person.Id = 123
	person.Module = "123"
	person.Content = "123"
	persons = append(persons, person)
	this.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": persons,
	})
}
