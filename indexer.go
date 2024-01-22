package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Email_structure map[string]string
type Payload_structure map[string]interface{}

var wg sync.WaitGroup

func get_content_in_path(folder string) []string {
	var content_list []string
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) < 1 {
		return []string{}
	}

	for _, file := range files {
		content_list = append(content_list, file.Name())
	}

	return content_list
}

func generate_email_json(content *bufio.Scanner) Email_structure {
	email_json := Email_structure{}
	line_index := 0
	email_json["id"] = generateId()
	for content.Scan() {
		if line_index <= 14 { // The line #14 contains the last email header
			split_content := strings.Split(content.Text(), ":")
			if len(split_content) > 1 {
				email_json[split_content[0]] = strings.Trim(split_content[1], " ")
			}
			line_index++
		} else {
			email_json["Body"] += content.Text()
		}
	}

	return email_json
}

func insert_data_batch(data Payload_structure) {
	user := "admin"
	password := "admin"
	auth := user + ":" + password
	bas64encoded_creds := base64.StdEncoding.EncodeToString([]byte(auth))
	zinc_url := "http://localhost:4080/api/_bulkv2"
	payload, _ := json.MarshalIndent(data, "", "   ")
	req_body := bytes.NewBuffer(payload)
	req, err := http.NewRequest("POST", zinc_url, req_body)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+bas64encoded_creds)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

}

func scan_folders_list(folder_list []string, database_path string) {

	records := []Email_structure{}

	for _, folder_name := range folder_list {
		folder_path := database_path + folder_name
		folder_content := get_content_in_path(folder_path)
		for _, sub_folder_name := range folder_content {
			sub_folder_path := folder_path + "/" + sub_folder_name
			sub_folder_content := get_content_in_path(sub_folder_path)

			for _, mail_item := range sub_folder_content {
				file_path := sub_folder_path + "/" + mail_item
				mail_file, _ := os.Open(file_path)
				mail_content := bufio.NewScanner(mail_file)
				records = append(records, generate_email_json(mail_content))
			}
		}
	}

	payload := Payload_structure{}
	payload["index"] = "enron"
	payload["records"] = records
	insert_data_batch(payload)
	wg.Done()
}

func generateId() string {
	randomNumber := rand.Int63()

	base36String := strconv.FormatInt(randomNumber, 36)

	if len(base36String) > 6 {
		return base36String[2:8]
	}
	return base36String
}

/*
This is the first approach I made to try and index all the emails
I'm leaving this function just to talk about it during the interview
*/
func first_try_to_index(database_path string) {
	records := []Email_structure{}

	current_path := database_path
	maildir_content := get_content_in_path(current_path)

	for _, name_folder := range maildir_content {
		current_path = database_path + name_folder
		mail_categories := get_content_in_path(current_path)

		for _, category := range mail_categories {
			current_path = database_path + name_folder + "/" + category
			mail_list := get_content_in_path(current_path)

			for _, mail_item := range mail_list {
				file_path := current_path + "/" + mail_item
				mail_file, _ := os.Open(file_path)
				mail_content := bufio.NewScanner(mail_file)
				records = append(records, generate_email_json(mail_content))

			}
		}
	}
	payload := Payload_structure{}
	payload["index"] = "enron"
	payload["records"] = records
	insert_data_batch(payload)
}

func main() {
	if len(os.Args) < 2 {
		panic("Para ejecutar el indexer, ingrese la dirección de la base de datos a indexar")
	}

	database_path := os.Args[1] + "/maildir/"
	current_path := database_path
	maildir_content := get_content_in_path(current_path)
	magic_number := 15 // Divides all the first level folders to make the index process quicker
	pointer_start := 0
	pointer_end := 0
	step := len(maildir_content) / magic_number

	log.Println("Incio de la indexacion")
	for i := 1; i <= magic_number; i++ {
		pointer_end = i * step
		wg.Add(1)
		go scan_folders_list(maildir_content[pointer_start:pointer_end], current_path)
		pointer_start = pointer_end
	}
	wg.Wait()
	log.Println("Fin de la indexacion")

}
