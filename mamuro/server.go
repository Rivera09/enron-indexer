package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/hello", basicHandler)
	router.Get("/email", getEmails)

	fs := http.FileServer(http.Dir("./dist"))
	router.Handle("/*", http.StripPrefix("/", fs))

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("No se pudo levantar el servidor", err)
	} else {
		fmt.Println("Running on port :3000")
	}

}

func getEmails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	type Query struct {
		Term  string `json:"term"`
		Field string `json:"field"`
	}

	type Payload struct {
		SearchType string   `json:"search_type"`
		Query      Query    `json:"query"`
		SortFields []string `json:"sort_fields"`
		From       int      `json:"from"`
		MaxResults int      `json:"max_results"`
		Source     []string `json:"_source"`
	}

	payload := Payload{
		SearchType: "match",
		Query: Query{
			Term:  "forecast",
			Field: "_all",
		},
		SortFields: []string{},
		From:       0,
		MaxResults: 20,
		Source:     []string{},
	}

	type Source struct {
		Body      string `json:"Body"`
		Date      string `json:"Date"`
		From      string `json:"From"`
		To        string `json:"To"`
		Subject   string `json:"Subject"`
		XFrom     string `json:"X-From"`
		XTo       string `json:"X-To"`
		XFilename string `json:"X-Filename"`
		Id        string `json:"id"`
	}

	type Hits_list_item struct {
		Source Source `json:"_source"`
	}

	type Hits struct {
		Hits []Hits_list_item `json:"hits"`
	}

	type ResponseData struct {
		Hits Hits `json:"hits"`
	}

	pretty_json, _ := json.MarshalIndent(payload, "", "   ")
	req_body := bytes.NewBuffer(pretty_json)

	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron/_search", req_body)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	user := "admin"
	password := "admin"
	auth := user + ":" + password
	bas64encoded_creds := base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+bas64encoded_creds)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body. ", err)
	}

	var responseData ResponseData
	json.Unmarshal(resp_body, &responseData) // Simplify the response object
	request_reponse, _ := json.MarshalIndent(responseData.Hits.Hits, "", "   ")
	w.Write(request_reponse)
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
