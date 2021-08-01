package Services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goLangRestApiMuxGorm/DbContext_Migration"
	models "goLangRestApiMuxGorm/Models"
	"net/http"
)



func GetUsers(httpWriter http.ResponseWriter, httpRequest *http.Request)  {
	httpWriter.Header().Set("Content-Type", "application/json")
	var users[] models.User //creating slice object to storedata
	dbContext.DB.Find(&users) //passing pointer of slice to get all users
	json.NewEncoder(httpWriter).Encode(users) //returning slice of all users

}


func GetUser(httpWriter http.ResponseWriter, httpRequest *http.Request)  {
	httpWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(httpRequest)
	var user models.User
	dbContext.DB.Find(&user, params["id"])
	json.NewEncoder(httpWriter).Encode(user)
}


func CreateUser(httpWriter http.ResponseWriter, httpRequest *http.Request)  {
	httpWriter.Header().Set("Content-Type", "application/json")
	var user models.User //creating user object
	json.NewDecoder(httpRequest.Body).Decode(&user) //passing pointer of user object and getting all data
													//assigned to it from body
	dbContext.DB.Create(&user) //passing pointer to db for adding to database
	json.NewEncoder(httpWriter).Encode(user) //returning created db user object
}


func UpdateUser(httpWriter http.ResponseWriter, httpRequest *http.Request)  {
	httpWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(httpRequest) //getting all parameters
	var user models.User //creating user property
	dbContext.DB.First(&user, params["id"]) //getting user from db against id from params
	json.NewDecoder(httpRequest.Body).Decode(&user) //updating user from request body
	dbContext.DB.Save(&user) //saving user info
	json.NewEncoder(httpWriter).Encode(user) //returning updated user info
}


func DeleteUser(httpWriter http.ResponseWriter, httpRequest *http.Request)  {
	params := mux.Vars(httpRequest) //getting all parameters
	var user models.User //creating user property
	dbContext.DB.Delete(&user, params["id"]) //getting user from db against id from params
	json.NewEncoder(httpWriter).Encode("The user is deleted successfully ") //returning created db user object
}
