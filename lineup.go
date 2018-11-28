package main

type StartingEleven struct {
	L LineUp `json:"lineup"`
}

type LineUp struct {
		Formation   string `json:"type"`
		PlayerIds 	[]int  `json:"playersID"`
	} 
