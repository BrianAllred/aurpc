package aurpc

type Search struct {
	Version     int    `json:"version"`
	Type        string `json:"type"`
	Resultcount int    `json:"resultcount"`
	Results     []struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"results"`
}
