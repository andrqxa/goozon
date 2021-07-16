package main

import (
	"fmt"
	"os"

	"github.com/andrqxa/goozon/model"
)

func main() {
	parkSeq, err := model.GetParkingSequence(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The maximum number of vehicles at the same time in the parking lot equals %d\n", parkSeq.GetMaxParkingPlace())
}
