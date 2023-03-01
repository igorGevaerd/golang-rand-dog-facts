package main

type DogFact struct {
	Fact string `json:"body"`
}

type DogInfo struct {
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	Attributes DogFact `json:"attributes"`
}

type DogData struct {
	Data []DogInfo `json:"data"`
}

// response example:
// {
// 	"data":[
// 		 {
// 				"id":"591a1d34-475e-4a8d-862a-baeca4f2b19c",
// 				"type":"fact",
// 				"attributes":{
// 					 "body":"Most dogs have 18 or more muscles to tilt, rotate, and move their ears."
// 				}
// 		 }
// 	]
// }
