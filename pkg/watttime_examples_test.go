package pkg

import (
	"fmt"
)

func ExampleWattTime_Login(){
	wattTime := New()
	token, err := wattTime.Login("<username>", "<password>")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(token)
}

func ExampleWattTime_Index_withBa() {
	wattTime := New()
	options := IndexOptions{Ba: "SE"}
	realTimeEmissionsIndex, err := wattTime.Index("123abc", options)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Marginal Operating Emissions Rate (MOER)", realTimeEmissionsIndex.Moer)
}

func ExampleWattTime_Index_withLatitudeLongitude() {
	wattTime := New()
	options := IndexOptions{Latitude: 42.372, Longitude: -72.519}
	realTimeEmissionsIndex, err := wattTime.Index("123abc", options)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Marginal Operating Emissions Rate (MOER)", realTimeEmissionsIndex.Moer)
}

func ExampleWattTime_Index_withStyle() {
	wattTime := New()
	options := IndexOptions{Ba: "SE", Style: "Percent"}
	realTimeEmissionsIndex, err := wattTime.Index("123abc", options)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Marginal Operating Emissions Rate (MOER)", realTimeEmissionsIndex.Percent)
}
