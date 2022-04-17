package pkg

type Token struct {
	Value string `json:"token"`
}

type RealTimeEmissionsIndex struct {
	Freq      string
	Ba        string
	Percent   string
	Moer      string
	PointTime string `json:"point_time"`
}

type IndexOptions struct {
	Ba 		  string
	Latitude  float64
	Longitude float64
	Style     string
}
