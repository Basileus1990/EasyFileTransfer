// API to access and edit database's kept directories

package database

import "github.com/Basileus1990/EasyFileTransfer.git/APIserver/dataStructs"

func AddDir(mainDir *dataStructs.Directory, key string) {
	db[key] = *mainDir
}

func ContainsDir(key string) bool {
	_, ok := db[key]
	return ok
}
