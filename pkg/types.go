package pkg

type Token struct {
	// A token for authentication towards the WattTime API
	Value string `json:"token"`
}

type RealTimeEmissionsIndex struct {
	// Duration in seconds for which the data is valid from point_time
	Freq string
	//	Balancing authority abbreviation
	Ba string
	// A percentile value between 0 (minimum MOER in the last month i.e. clean) and 100
	// (maximum MOER in the last month i.e. dirty) representing the relative realtime
	// marginal emissions intensity.
	Percent string
	// Marginal Operating Emissions Rate (MOER) value measured in lbs/MWh
	Moer string
	// ISO8601 UTC date/time format indicating when this data became valid
	PointTime string `json:"point_time"`
}

type IndexOptions struct {
	// Options for the Index function.
	// Provided either Ba or Latitude and Longitude, but not all three.
	Ba        string
	Latitude  float64
	Longitude float64
	// Units in which to provide realtime marginal emissions.
	Style string
}
