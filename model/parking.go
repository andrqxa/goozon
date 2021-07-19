package model

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

const (
	timeTemplate = "02.01.2006 15:04:05"
)

type Parking struct {
	Start time.Time
	End   time.Time
}

func (p Parking) Equal(other Parking) bool {
	return p.Start.Equal(other.Start) && p.End.Equal(other.End)
}

func (p Parking) String() string {
	return fmt.Sprintf("[from: %s to: %s]", p.Start, p.End)
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

func (a ParkingSeq) String() string {
	res := make([]string, 0)
	for _, p := range a {
		res = append(res, p.String())
	}
	return strings.Join(res, "\n")
}

func (a ParkingSeq) Equal(o ParkingSeq) bool {
	if a == nil || o == nil || a.Len() != o.Len() {
		return false
	}
	sort.Sort(a)
	sort.Sort(o)
	for i := 0; i < a.Len(); i++ {
		if !a[i].Equal(o[i]) {
			return false
		}
	}
	return true
}

func GetParkingSequence(i io.Reader) (ParkingSeq, error) {
	result := make(ParkingSeq, 0)
	scanner := bufio.NewScanner(i)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		nextString := scanner.Text()
		x := strings.Split(nextString, ";")
		first, second := strings.TrimSpace(x[0]), strings.TrimSpace(x[1])
		startTime, err := time.Parse(timeTemplate, first)
		if err != nil {
			panic(err)
		}
		endTime, err := time.Parse(timeTemplate, second)
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

func (parkSeq ParkingSeq) GetMaxParkingPlace() int {
	if parkSeq.Len() == 0 {
		return 0
	}
	maxPlace := 0
	start := parkSeq[0].Start
	end := parkSeq[len(parkSeq)-1].End
	duration := end.Sub(start).Nanoseconds() / 1e9
	parkingDuration := make([]int, duration, duration)
	for _, p := range parkSeq {
		currentDuration := p.End.Sub(p.Start).Nanoseconds() / 1e9
		offcet := p.Start.Sub(start).Nanoseconds() / 1e9
		fmt.Printf("currentDuration=%d\n", currentDuration)
		fmt.Printf("offcet=%d\n", offcet)
		for i := 0; int64(i) < currentDuration-1; i++ {
			parkingDuration[offcet+int64(i)]++
			currentElement := parkingDuration[offcet+int64(i)]
			if currentElement > maxPlace {
				maxPlace = currentElement
			}
		}
	}
	return maxPlace
}
