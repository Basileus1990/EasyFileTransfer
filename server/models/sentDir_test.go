package models

import (
	"log"
	"testing"
)

func TestRecivingJSONDir(t *testing.T) {
	sampleJSONs := `{"name":"main","subdirs":[{"name":"main2","subdirs":[]}],"files":[{"name":"test.txt","type":"txt","size":5000}]}`
	mainDir, err := ConvertJSONToDirTree([]byte(sampleJSONs))
	if err != nil {
		t.Fatal(err, " for input: ", sampleJSONs)
	}
	log.Println(mainDir.Name)
	log.Println(mainDir.Files[0].Size)
	log.Println(mainDir.SubDirs[0].Name)
}
