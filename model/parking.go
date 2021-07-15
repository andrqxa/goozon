package model

import (
	"bufio"
	"io"
	"sort"
	"strings"
	"time"
)

type Parking struct {
	Start time.Time
	End   time.Time
}

type ParkingSeq []Parking

func (a ParkingSeq) Len() int {
	return len(a)
}
func (a ParkingSeq) Less(i, j int) bool {
	return a[i].Start.Before(a[j].Start)
}
func (a ParkingSeq) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func GetParkingSequence(i io.Reader) (ParkingSeq, error) {
	result := make(ParkingSeq, 0)
	for {
		s, err := bufio.NewReader(i).ReadString('\n')
		if err != nil {
			panic(err)
		}
		if err != io.EOF {
			break
		}
		x := strings.Split(s, " ")
		first, second := x[0], x[1]

		startTime, err := time.Parse("02.01.2006 15:04:05", first)
		if err != nil {
			panic(err)
		}
		endTime, err := time.Parse("02.01.2006 15:04:05", second)
		if err != nil {
			panic(err)
		}
		if endTime.Before(startTime) {
			panic(err)
		}
		result = append(result, Parking{startTime, endTime})
	}
	sort.Sort(result)
	return result, nil
}
