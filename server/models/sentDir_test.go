package models

import (
	"log"
	"testing"
)

func TestRecivingJSONDir(t *testing.T) {
	var validJSONs []string = []string{
		`{
			"name":"main1",
			"subdirs":[
				{
					"name":"main2",
					"subdirs":[],
					"files":[]
				}
			],
			"files":[
				{
					"name":"test.txt",
					"type":"txt",
					"size":5000
				}
			]
		}`,
		`{
			"name":"main2",
			"subdirs":[
				{
					"name":"submain2",
					"subdirs":[
						{
							"name":"subsubmain2",
							"subdirs":[
								{
									"name":"subsubsubmain2",
									"subdirs":[],
									"files":[
										{
											"name":"test.txt",
											"type":"txt",
											"size":50020
										}
									]
								}
							],
							"files":[]
						}
					],
					"files":[
						{
							"name":"test0.txt",
							"type":"txt",
							"size":50005
						},
						{
							"name":"test1.txt",
							"type":"txt",
							"size":50008
						},
						{
							"name":"test2.txt",
							"type":"txt",
							"size":500095
						}
					]
				}
			],
			"files":[
				{
					"name":"test.txt",
					"type":"txt",
					"size":5000
				}
			]
		}`,
	}
	for _, mainDir := range validJSONs {
		dir, err := ConvertJSONToDirTree([]byte(mainDir))
		if err != nil {
			t.Fatal(err, " for input: ", mainDir)
		}
		log.Println(dir.Name)
		log.Println(dir.Files[0].Size)
		log.Println(dir.SubDirs[0].Name)
	}
}
