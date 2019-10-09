package biwenger

type CurrentLineUpResponse struct {
	Status int        `json:"status"`
	Data   LineUpData `json:"data"`
}

type LineUpResponse struct {
	Status int            `json:"status"`
	Data   LineUpBaseData `json:"data"`
}

type StartingEleven struct {
	L LineUp `json:"lineup"`
}

type LineUp struct {
	Formation string `json:"type"`
	PlayerIds []int  `json:"playersID"`
}

type Owner struct {
	Date  int `json:"date"`
	Price int `json:"price"`
}

type PlayerBase struct {
	ID    int   `json:"id"`
	Owner Owner `json:"owner"`
}

type LineUpData struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Group      string        `json:"group"`
	Icon       string        `json:"icon"`
	Points     int           `json:"points"`
	Balance    int           `json:"balance"`
	JoinDate   int           `json:"joinDate"`
	LineupDate int           `json:"lineupDate"`
	LineUp     LineUp        `json:"lineup"`
	Market     []interface{} `json:"market"`
	Players    []PlayerBase  `json:"players"`
	Offers     []interface{} `json:"offers"`
}

type LineUpBaseData struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Group      string `json:"group"`
	Icon       string `json:"icon"`
	Points     int    `json:"points"`
	Balance    int    `json:"balance"`
	JoinDate   int    `json:"joinDate"`
	LineupDate int    `json:"lineupDate"`
}
