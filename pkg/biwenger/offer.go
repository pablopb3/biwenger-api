package biwenger

type ActionToOffer struct {
	Status string `json:"status"`
}

type AcceptOfferResponse struct {
	Data struct {
		Amount   int    `json:"amount"`
		Created  int    `json:"created"`
		FromID   int    `json:"fromID"`
		ID       int    `json:"id"`
		Modified int    `json:"modified"`
		Status   string `json:"status"`
		ToID     int    `json:"toID"`
		Type     string `json:"type"`
	} `json:"data"`
	Status int `json:"status"`
}

type PlaceOfferBody struct {
	Amount           int         `json:"amount"`
	RequestedPlayers []int       `json:"requestedPlayers"`
	To               interface{} `json:"to"`
	Type             string      `json:"type"`
}

type PlaceOfferResponse struct {
	Data struct {
		Amount  int    `json:"amount"`
		Created int    `json:"created"`
		FromID  int    `json:"fromID"`
		ID      int    `json:"id"`
		Status  string `json:"status"`
		ToID    int    `json:"toID"`
		Type    string `json:"type"`
		Until   int    `json:"until"`
	} `json:"data"`
	Status int `json:"status"`
}
