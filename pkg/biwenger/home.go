package biwenger

type Home struct {
	Status int  `json:"status"`
	Data   Data `json:"data"`
}

type Data struct {
	Events []Event `json:"events"`
}

type Event struct {
	Type  string `json:"type"`
	Date  int    `json:"date"`
	Round Round  `json:"round"`
}

type Round struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Short string `json:"short"`
}
