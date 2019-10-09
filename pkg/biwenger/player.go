package biwenger

type Player struct {
	Status int `json:"status"`
	Data   struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		Position       int    `json:"position"`
		Price          int    `json:"price"`
		FantasyPrice   int    `json:"fantasyPrice"`
		Status         string `json:"status"`
		PriceIncrement int    `json:"priceIncrement"`
		Competition    struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Slug     string `json:"slug"`
			Sport    string `json:"sport"`
			Currency string `json:"currency"`
			Country  string `json:"country"`
			Enabled  bool   `json:"enabled"`
			Type     string `json:"type"`
		} `json:"competition"`
		Team struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Slug      string `json:"slug"`
			NextMatch struct {
				ID     int    `json:"id"`
				Date   int    `json:"date"`
				Status string `json:"status"`
				Round  struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Short  string `json:"short"`
					Season struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Slug string `json:"slug"`
					} `json:"season"`
				} `json:"round"`
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
				Home struct {
					ID      int           `json:"id"`
					Name    string        `json:"name"`
					Slug    string        `json:"slug"`
					Score   interface{}   `json:"score"`
					Reports []interface{} `json:"reports"`
				} `json:"home"`
				Away struct {
					ID      int           `json:"id"`
					Name    string        `json:"name"`
					Slug    string        `json:"slug"`
					Score   interface{}   `json:"score"`
					Reports []interface{} `json:"reports"`
				} `json:"away"`
				Preview interface{} `json:"preview"`
				Summary interface{} `json:"summary"`
			} `json:"nextMatch"`
		} `json:"team"`
		Reports []struct {
			Match struct {
				ID     int    `json:"id"`
				Date   int    `json:"date"`
				Status string `json:"status"`
				Round  struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Short string `json:"short"`
				} `json:"round"`
				Home struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Slug  string `json:"slug"`
					Score int    `json:"score"`
				} `json:"home"`
				Away struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Slug  string `json:"slug"`
					Score int    `json:"score"`
				} `json:"away"`
			} `json:"match"`
			Events []struct {
				Type     int `json:"type"`
				Metadata int `json:"metadata"`
			} `json:"events"`
			Points int  `json:"points"`
			Star   bool `json:"star,omitempty"`
		} `json:"reports"`
		Prices  [][]int `json:"prices"`
		Seasons []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Competition struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Slug     string `json:"slug"`
				Sport    string `json:"sport"`
				Currency string `json:"currency"`
				Country  string `json:"country"`
				Enabled  bool   `json:"enabled"`
				Type     string `json:"type"`
			} `json:"competition,omitempty"`
			Player struct {
				ID   int    `json:"id"`
				Slug string `json:"slug"`
			} `json:"player,omitempty"`
		} `json:"seasons"`
		News []struct {
			ID       int    `json:"id"`
			Source   string `json:"source"`
			Title    string `json:"title"`
			URL      string `json:"url"`
			Date     int    `json:"date"`
			Photo    string `json:"photo"`
			Video    string `json:"video"`
			Comments int    `json:"comments"`
		} `json:"news"`
		Threads []struct {
			ID       int    `json:"id"`
			Source   string `json:"source"`
			Title    string `json:"title"`
			URL      string `json:"url"`
			Date     int    `json:"date"`
			Photo    string `json:"photo"`
			Video    string `json:"video"`
			Comments int    `json:"comments"`
		} `json:"threads"`
		Fitness          []int `json:"fitness"`
		Points           int   `json:"points"`
		PlayedHome       int   `json:"playedHome"`
		PlayedAway       int   `json:"playedAway"`
		PointsHome       int   `json:"pointsHome"`
		PointsAway       int   `json:"pointsAway"`
		PointsLastSeason int   `json:"pointsLastSeason"`
		ScoreID          int   `json:"scoreID"`
	} `json:"data"`
}

type PlayersAliasInfo struct {
	NumID struct {
		ID               int           `json:"id"`
		Name             string        `json:"name"`
		Slug             string        `json:"slug"`
		TeamID           int           `json:"teamID"`
		Position         int           `json:"position"`
		Price            int           `json:"price"`
		FantasyPrice     int           `json:"fantasyPrice"`
		Status           string        `json:"status"`
		PriceIncrement   int           `json:"priceIncrement"`
		Fitness          []interface{} `json:"fitness"`
		Points           int           `json:"points"`
		PlayedHome       int           `json:"playedHome"`
		PlayedAway       int           `json:"playedAway"`
		PointsHome       int           `json:"pointsHome"`
		PointsAway       int           `json:"pointsAway"`
		PointsLastSeason int           `json:"pointsLastSeason"`
	} `json:"id"`
}
