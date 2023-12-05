package models


type TestStruct struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

var MockAlbum = []TestStruct {
	{ 
		Name: "1",
		Age: 1,
	},
	{ 
		Name: "2",
		Age: 2,
	},
	{ 
		Name: "3",
		Age: 3,
	},
 } 