package aurpc

type Error struct {
	Version     int           `json:"version"`
	Type        string        `json:"type"`
	Resultcount int           `json:"resultcount"`
	Results     []interface{} `json:"results"`
	Error       string        `json:"error"`
}
