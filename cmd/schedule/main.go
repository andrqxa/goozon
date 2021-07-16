package main

import (
	"fmt"
	"os"

	"github.com/andrqxa/goozon/model"
)

var (
	unorderTestDate = `16.07.2021 09:00:00; 16.07.2021 09:30:00
	16.07.2021 08:15:00; 16.07.2021 08:45:00
	16.07.2021 08:00:00; 16.07.2021 08:30:00
	16.07.2021 08:40:00; 16.07.2021 09:00:00
	`
	orderTestDate model.ParkingSeq
)

const (
	timeTemplate = "02.01.2006 15:04:05"
)

func main() {
	parkSeq, err := model.GetParkingSequence(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The maximum number of vehicles at the same time in the parking lot equals %d\n", parkSeq.GetMaxParkingPlace())
}
