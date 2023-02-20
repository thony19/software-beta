package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type records struct {
	Message string `json:"message"`
	Date    string `json:"date"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
}

var saveJSON = "data.json"

func CreateStruct(message string, date string, from string, to string) records {
	a := records{
		Message: message,
		Date:    date,
		From:    from,
		To:      to,
	}

	return a
}

// func search(directory string) string {
// 	dir, err := os.ReadDir(directory)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// for _, folders := range dir {
// 	// 	if folders.IsDir()
// 	// }

// }

func main() {
	listJson := []records{}
	var subjectRe string
	dir, err := os.ReadDir("../../tracing/enron_mail_20110402/maildir")

	if err != nil {
		log.Fatal(err)
	}

	for _, folders := range dir {
		if folders.IsDir() {
			fmt.Println("Folder:", folders.Name())
			dir_inf, err := os.ReadDir("../../tracing/enron_mail_20110402/maildir/" + folders.Name())

			if err != nil {
				log.Fatal(err)
			}

			for _, finders := range dir_inf {
				if finders.IsDir() && len(listJson) < 5 {
					fmt.Println("info:", finders.Name())

					archive, err := os.ReadDir("../../tracing/enron_mail_20110402/maildir/" + folders.Name() + "/" + finders.Name())

					if err != nil {
						log.Fatal(err)
					}

					for _, info := range archive {
						if !info.IsDir() {
							// fmt.Println(info.Name(), "=======================================")
							dataBytes, err := os.ReadFile("../../tracing/enron_mail_20110402/maildir/" + folders.Name() + "/" + finders.Name() + "/" + info.Name())
							if err != nil {
								log.Fatal(err)
							}

							dataString := string(dataBytes)

							dataQuotes := strings.Replace(dataString, `"`, `-`, -1)

							dataSplit := strings.Split(dataQuotes, "\n")

							dataSplit_m := strings.Split(dataSplit[0], ":")
							dataSplit_d := strings.Split(dataSplit[1], ":")
							dataSplit_f := strings.Split(dataSplit[2], ":")
							dataSplit_t := strings.Split(dataSplit[3], ":")

							for index, aux := range dataSplit {
								if index > 4 {
									subjectRe += aux + "\n"
								}
							}

							datosJson := records{
								Message: dataSplit_m[1],
								Date:    dataSplit_d[1],
								From:    dataSplit_f[1],
								To:      dataSplit_t[1],
								Subject: subjectRe,
							}

							listJson = append(listJson, datosJson)
							fmt.Println(len(listJson))

						}
						// fmt.Println("-----------------------------------------")

					}
				}
			}
		}

	}

	dataJson := map[string]interface{}{
		"index":   "data.json",
		"records": listJson,
	}

	jsonData, err := json.MarshalIndent(dataJson, "", "  ")
	if err != nil {
		log.Fatalf("could not convert struct data into json file: %v", err)
	}

	if err = os.WriteFile(saveJSON, jsonData, 0644); err != nil {
		log.Fatalf("could not saveJSON file: %v", err)
	}

	tt, _ := os.ReadFile("./data.json")
	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulkv2", bytes.NewBuffer(tt))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
