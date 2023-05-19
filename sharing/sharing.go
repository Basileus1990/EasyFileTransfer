package sharing

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Basileus1990/EasyFileTransfer.git/server/dataStructs"
)

const serverURL = "http://localhost:8080/share"
const maxDepth = 2

func Share(path string) error {
	var mainDir dataStructs.Directory
	mainDir.Name = filepath.Base(path)
	err := getDirContent(path, &mainDir, 0)
	if err != nil {
		return err
	}
	marshaled, err := json.Marshal(mainDir)
	if err != nil {
		return err
	}

	_, err = sendDirContent(marshaled)
	if err != nil {
		return err
	}

	return nil
}

func getDirContent(path string, parent *dataStructs.Directory, depth int) error {
	if path[len(path)-1] != '/' {
		path += "/"
	}
	dirContent, err := os.ReadDir(path)
	if err != nil {
		return errors.New("a wrong path was given! Please try again")
	}

	for _, element := range dirContent {
		fInfo, err := os.Stat(path + element.Name())
		if err != nil {
			return errors.New("a wrong path was given! Please try again")
		}

		if fInfo.IsDir() {
			// adding directories and recusively getting subdirs
			var newDir dataStructs.Directory
			newDir.Name = element.Name()
			getSubDirContent(path+element.Name()+"/", &newDir, depth+1)
			parent.SubDirs = append(parent.SubDirs, newDir)

		} else {
			// adding files
			parent.Files = append(parent.Files, dataStructs.File{Name: element.Name(), Size: int(fInfo.Size())})
		}
	}

	return nil
}

func getSubDirContent(path string, parent *dataStructs.Directory, depth int) error {
	if depth > maxDepth {
		return nil
	}
	err := getDirContent(path, parent, depth)
	if err != nil {
		return err
	}
	return nil
}

func sendDirContent(data []byte) (*http.Response, error) {
	request, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	log.Println("Status: ", response.Status)
	body, _ := ioutil.ReadAll(response.Body)
	log.Println("Response: ", string(body))

	return response, nil
}
