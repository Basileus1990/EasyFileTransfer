package models

import "encoding/json"

// model for creating a structure for sent directiories and files

type file struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int    `json: "type"`
}

type Directory struct {
	Name    string      `json:"name"`
	SubDirs []Directory `json:"subdirs`
	Files   []file      `json:"files"`
}

func ConvertJSONToDirTree(sentDir []byte) (*Directory, error) {
	var mainDir Directory
	err := json.Unmarshal(sentDir, &mainDir)
	if err != nil {
		return nil, err
	}

	return &mainDir, nil
}
