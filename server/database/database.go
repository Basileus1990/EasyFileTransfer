package database

import "github.com/Basileus1990/EasyFileTransfer.git/server/dataStructs"

// temporary
var db map[string]dataStructs.Directory

func init() {
	db = make(map[string]dataStructs.Directory)
}
