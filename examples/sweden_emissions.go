package main

import (
	"fmt"
	watttime "github.com/IuryAlves/watttime-go/pkg"
	"os"
)

func main() {
	token, err := watttime.Login(os.Getenv("WATTTIME_USER"), os.Getenv("WATTTIME_PASSWORD"))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	realTimeEmissionsIndex, err := watttime.Index(token, "SE")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Clean energy emissions for %s is %s%% \n", realTimeEmissionsIndex.Ba, realTimeEmissionsIndex.Percent)
	os.Exit(0)
}
