package meander

type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

type googleResponse struct {
	Results []*Place `json:"results"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}
