package biwenger

type PlayerInMarket struct {
	IdPlayer int `json:"idPlayer"`
	Price    int `json:"price"`
	IdUser   int `json:"idUser"`
}

type ReceivedOffer struct {
	IdOffer  int `json:"idOffer"`
	IdPlayer int `json:"idPlayer"`
	Amount   int `json:"amount"`
	IdUser   int `json:"idUser"`
}

type MarketResponse struct {
	Data struct {
		Loans  []interface{} `json:"loans"`
		Offers []struct {
			Amount  int `json:"amount"`
			Created int `json:"created"`
			From    struct {
				Icon string `json:"icon"`
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"from"`
			ID               int    `json:"id"`
			RequestedPlayers []int  `json:"requestedPlayers"`
			Status           string `json:"status"`
			To               struct {
				Icon string `json:"icon"`
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"to"`
			Type  string `json:"type"`
			Until int    `json:"until"`
		} `json:"offers"`
		Sales []struct {
			Date   int `json:"date"`
			Player struct {
				ID int `json:"id"`
			} `json:"player"`
			Price int `json:"price"`
			Until int `json:"until"`
			User  struct {
				ID int `json:"id"`
			} `json:"user"`
		} `json:"sales"`
		Status struct {
			Balance    int `json:"balance"`
			MaximumBid int `json:"maximumBid"`
		} `json:"status"`
	} `json:"data"`
	Status int `json:"status"`
}

type MarketStatsResponse struct {
	Status int `json:"status"`
	Data   struct {
		Competition struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Slug     string `json:"slug"`
			Sport    string `json:"sport"`
			Currency string `json:"currency"`
			Country  string `json:"country"`
			Enabled  bool   `json:"enabled"`
			Type     string `json:"type"`
		} `json:"competition"`
		Values [][]int `json:"values"`
		Ups    []struct {
			OldPrice       int    `json:"oldPrice"`
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Slug           string `json:"slug"`
			Position       int    `json:"position"`
			Price          int    `json:"price"`
			FantasyPrice   int    `json:"fantasyPrice"`
			Country        string `json:"country"`
			Birthday       int    `json:"birthday"`
			Status         string `json:"status"`
			PriceIncrement int    `json:"priceIncrement"`
			Difference     int    `json:"difference"`
			Team           struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"team"`
			Fitness          []int  `json:"fitness"`
			Points           int    `json:"points"`
			PlayedHome       int    `json:"playedHome"`
			PlayedAway       int    `json:"playedAway"`
			PointsHome       int    `json:"pointsHome"`
			PointsAway       int    `json:"pointsAway"`
			PointsLastSeason int    `json:"pointsLastSeason"`
			StatusText       string `json:"statusText,omitempty"`
		} `json:"ups"`
		Downs []struct {
			OldPrice       int    `json:"oldPrice"`
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Slug           string `json:"slug"`
			Position       int    `json:"position"`
			Price          int    `json:"price"`
			FantasyPrice   int    `json:"fantasyPrice"`
			Country        string `json:"country"`
			Birthday       int    `json:"birthday"`
			Status         string `json:"status"`
			PriceIncrement int    `json:"priceIncrement"`
			Difference     int    `json:"difference"`
			Team           struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"team"`
			Fitness          []interface{} `json:"fitness"`
			Points           int           `json:"points"`
			PlayedHome       int           `json:"playedHome"`
			PlayedAway       int           `json:"playedAway"`
			PointsHome       int           `json:"pointsHome"`
			PointsAway       int           `json:"pointsAway"`
			PointsLastSeason int           `json:"pointsLastSeason"`
			StatusText       string        `json:"statusText,omitempty"`
		} `json:"downs"`
	} `json:"data"`
}

type SendToMarket struct {
	Type  string `json:"type"`
	Price string `json:"price"`
}

type StatusResponse struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}
