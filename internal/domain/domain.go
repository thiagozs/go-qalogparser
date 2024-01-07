package domain

type Game struct {
	TotalKills   int            `json:"total_kills"`
	Players      []string       `json:"players"`
	Kills        map[string]int `json:"kills"`
	KillsbyMeans map[string]int `json:"kills_by_means"`
}

type Matches struct {
	Games []map[string]interface{} `json:"games"`
}
