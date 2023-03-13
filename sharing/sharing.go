package sharing

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

var serverURL = "http://localhost:8080/share"

type dirElement struct {
	Name string
	Type string
}

// args[0] is the path
func Share(args []string) error {
	data, err := getDirContentAsJSON(args)
	if err != nil {
		return err
	}

	status, err := sendDirContent(data)
	if err != nil {
		return err
	}
	log.Println("Status: ", status)

	return nil
}

func getDirContentAsJSON(args []string) ([]byte, error) {
	dirContent, err := os.ReadDir(args[0])
	if err != nil {
		return nil, errors.New("a wrong path was given! Please try again")
	}

	var data []dirElement
	for _, element := range dirContent {
		data = append(data, dirElement{element.Name(), element.Type().String()})
	}

	return json.Marshal(data)
}

func sendDirContent(data []byte) (string, error) {
	request, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	return response.Status, nil
}
