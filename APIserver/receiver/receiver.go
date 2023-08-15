package receiver

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Basileus1990/EasyFileTransfer.git/APIserver/dataStructs"
	"github.com/Basileus1990/EasyFileTransfer.git/APIserver/database"
)

const (
	keyLength = 64
)

// Takes the JSON file, gets it converted and adds to TODO: database using randomly generated key
func ShareFiles(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	mainDir, err := convertJSONToDirTree(body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key, err := generateUniqueKey()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	database.AddDir(mainDir, key)

	w.WriteHeader(http.StatusOK)
	response := []byte(fmt.Sprintf("{\"key\":\"%s\"}", key))
	w.Write(response)
}

// converts given JSON and checks its validity
func convertJSONToDirTree(sentDir []byte) (*dataStructs.Directory, error) {
	var mainDir dataStructs.Directory
	err := json.Unmarshal(sentDir, &mainDir)
	if err != nil {
		return nil, errors.New("sent data format is invalid")
	}

	err = mainDir.IsValidDirectory()
	if err != nil {
		return nil, err
	}

	return &mainDir, nil
}

// generates a unique key for the directory
func generateUniqueKey() (string, error) {
	key := make([]byte, keyLength)
	for {
		_, err := rand.Read(key)
		if err != nil {
			return "", err
		}

		encodedKey := base64.StdEncoding.EncodeToString(key)
		if !database.ContainsDir(string(encodedKey)) {
			return encodedKey, nil
		}
	}
}
