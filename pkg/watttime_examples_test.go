package pkg

import (
	"fmt"
	"testing"
)

func ExampleWattTime_Login(){
	wattTime := New()
	token, err := wattTime.Login("<username>", "<password>")
	if err != nil {
		fmt.Println(err.Error())
	}
}
