package main

import (
	"os"

	"github.com/andrqxa/goozon/model"
)

func main() {
	parkSeq, err := model.GetParkingSequence(os.Stdin)
	if err != nil {
		panic(err)
	}
	maxPlace := 0
	start := parkSeq[0].Start
	end := parkSeq[len(parkSeq)-1].End
	duration := end.Sub(start).Nanoseconds() / 1e9
	parkingDuration := make([]int, duration, duration)
	for _, p := range parkSeq {
		currentDuration := p.End.Sub(p.Start).Nanoseconds() / 1e9
		offcet := p.Start.Sub(start).Nanoseconds() / 1e9
		for i := 0; int64(i) < currentDuration; i++ {
			parkingDuration[offcet+int64(i)]++
			currentElement := parkingDuration[offcet+int64(i)]
			if currentElement > maxPlace {
				maxPlace = currentElement
			}
		}
	}
}
