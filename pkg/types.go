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
