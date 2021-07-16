package main

import (
	"fmt"
	"io"
	"os"

	"github.com/andrqxa/goozon/model"
)

var (
	inData io.Reader
	err1   error
)

func main() {
	switch len(os.Args) {
	case 1:
		inData = os.Stdin
	case 2:
		{
			inData, err1 = os.Open(os.Args[1])
			if err1 != nil {
				panic(err1)
			}
		}
	default:
		panic("Too many arguments. You chould set only 1 - path to file with data")
	}
	parkSeq, err := model.GetParkingSequence(inData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The maximum number of vehicles at the same time in the parking lot equals %d\n", parkSeq.GetMaxParkingPlace())
}
