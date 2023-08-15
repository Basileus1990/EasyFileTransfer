package receiver

import (
	"testing"
)

func TestRecivingValidJSONDir(t *testing.T) {
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
							"size":50005
						},
						{
							"name":"test1.txt",		
							"size":50008
						},
						{
							"name":"test2.txt",		
							"size":500095
						}
					]
				}
			],
			"files":[
				{
					"name":"test.txt",
					"size":5000
				}
			]
		}`,
		`{
			"name": "main1",
			"subdirs": [
				{
					"name": "main2",
					"subdirs": [
						{
							"name": "subdir1",
							"subdirs": [],
							"files": [
								{
									"name": "file1.txt",					
									"size": 10000
								},
								{
									"name": "file2.pdf",					
									"size": 20000
								}
							]
						},
						{
							"name": "subdir2",
							"subdirs": [
								{
									"name": "subdir3",
									"subdirs": [],
									"files": [
										{
											"name": "file3.jpg",							
											"size": 30000
										}
									]
								}
							],
							"files": [
								{
									"name": "file4.png",					
									"size": 40000
								}
							]
						}
					],
					"files": [
						{
							"name": "test.txt",			
							"size": 5000
						}
					]
				}
			],
			"files": [
				{
					"name": "file5.txt",	
					"size": 60000
				}
			]
		}`,
		`{
			"name":"test",
			"subdirs":[],
			"files":[]
		}`,
	}
	for _, mainDir := range validJSONs {
		_, err := convertJSONToDirTree([]byte(mainDir))
		if err != nil {
			t.Fatal(err, " for input: ", mainDir)
		}
	}
}

func TestRecivingInvalidJSONDir(t *testing.T) {
	var invalidJSONs []string = []string{
		`{
			"subdirs":[],
			"files":[]
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
							"name":"test0.txt"
							"size":50005
						},
						{
							"name":"test1.txt"
							"size":50008
						},
						{
							"name":"test2.txt"
							"size":500095
						}
					]
				}
			],
			"files":[
				{
					"name":"test.txt"
					"size":5000
				}
			]
		}`,
		`{
			"name": "main1",
			"subdirs": [
				{
					"name": "main2",
					"subdirs": [
						{
							"name": "subdir1",
							"subdirs": [],
							"files": [
								{
									"name": "file1.txt",
									"size": 10000
								},
								{
									"name": "file2.pdf",
									"size": 20000
								}
							]
						},
						{
							"name": "subdir2",
							"subdirs": [
								{
									"name": "subdir3",
									"subdirs": [],
									"files": [
										{
											"name": "file3.jpg",
											"size": 30000
										}
									]
								
							],
							"files": [
								{
									"name": "file4.png",
									"size": 40000
								}
							]
						}
					],
					"files": [
						{
							"name": "test.txt",
							"size": 5000
						}
					]
				}
			],
			"files": [
				{
					"name": "file5.txt",
					"size": 60000
				}
			]
		}`,
	}
	for _, mainDir := range invalidJSONs {
		_, err := convertJSONToDirTree([]byte(mainDir))
		if err == nil {
			t.Fatal("Did not raise an error for invalid json!: ", mainDir)
		}
	}
}
