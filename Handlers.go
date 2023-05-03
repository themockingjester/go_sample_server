package main

import (
	"encoding/json"
	"fmt"

	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

type responseForm struct {
	Success bool
	Message string
	Code    int
	Data    any
}

func AboutPage(res http.ResponseWriter, req *http.Request) {
	RenderTemplate(res, "about.page.tmpl")
}

// JobReq represents a job request.
type PostBodyParamsStructure struct {
	ClassName int `json:"class"`
}

type UserDataBodyParams struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func AddUser(res http.ResponseWriter, req *http.Request) {
	bodyParams := UserDataBodyParams{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&bodyParams); err != nil {
		fmt.Printf("error parsing request JSON: %v", err)
		res.Header().Set("Content-Type", "application/json")

		result := responseForm{Success: true, Message: "error parsing request JSON", Code: 200, Data: struct{}{}}
		responseJson, _ := json.Marshal(result)
		res.Write(responseJson)
		return
	}
	fmt.Println(bodyParams)
	insertionRes := bodyParams.AddUserInDB()
	result := responseForm{Success: insertionRes.Success, Message: insertionRes.Message, Code: insertionRes.Code, Data: insertionRes.Data}
	res.Header().Set("Content-Type", "application/json")
	responseJson, _ := json.Marshal(result)
	res.Write(responseJson)
	return
}
func PostRequest(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	source := chi.URLParam(req, "source")
	//Extracting body parameters from request
	bodyParams := PostBodyParamsStructure{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&bodyParams); err != nil {
		fmt.Printf("error parsing request JSON: %v", err)
		res.Header().Set("Content-Type", "application/json")

		result := responseForm{Success: true, Message: "error parsing request JSON", Code: 200, Data: struct{}{}}
		responseJson, _ := json.Marshal(result)
		res.Write(responseJson)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	result := responseForm{Success: true, Message: "Successfully received post request", Code: 200, Data: struct {
		QueryParameter  string
		BodyParameter   int
		ParamsParameter string
	}{
		QueryParameter:  id,
		BodyParameter:   bodyParams.ClassName,
		ParamsParameter: source,
	}}

	responseJson, _ := json.Marshal(result)
	res.Write(responseJson)

}
func GetRequest(res http.ResponseWriter, req *http.Request) {
	result := struct {
		Result string
	}{
		Result: fmt.Sprintf("Request Received"),
	}
	response := responseForm{Success: true, Message: "Successfully performed action", Code: 200, Data: result}

	res.Header().Set("Content-Type", "application/json")

	responseJson, _ := json.Marshal(response)
	res.Write(responseJson)
}

func RenderTemplate(res http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templateName)
	err := parsedTemplate.Execute(res, parsedTemplate)
	if err != nil {
		fmt.Sprintln("Getting error while rendering template: %v", err)
		return
	}
}
