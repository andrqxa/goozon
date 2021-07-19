package model

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	time11, _ := time.Parse(timeTemplate, "16.07.2021 08:00:00")
	time12, _ := time.Parse(timeTemplate, "16.07.2021 08:30:00")
	time21, _ := time.Parse(timeTemplate, "16.07.2021 08:15:00")
	time22, _ := time.Parse(timeTemplate, "16.07.2021 08:45:00")
	time31, _ := time.Parse(timeTemplate, "16.07.2021 08:40:00")
	time32, _ := time.Parse(timeTemplate, "16.07.2021 09:00:00")
	time41, _ := time.Parse(timeTemplate, "16.07.2021 09:00:00")
	time42, _ := time.Parse(timeTemplate, "16.07.2021 09:30:00")
	orderTestDate = ParkingSeq{
		Parking{time11, time12},
		Parking{time21, time22},
		Parking{time31, time32},
		Parking{time41, time42},
	}
}

var (
	unorderTestDate = `16.07.2021 09:00:00; 16.07.2021 09:30:00
	16.07.2021 08:15:00; 16.07.2021 08:45:00
	16.07.2021 08:00:00; 16.07.2021 08:30:00
	16.07.2021 08:40:00; 16.07.2021 09:00:00`

	orderTestDate ParkingSeq
)

func TestParkingEqual(t *testing.T) {
	assert := assert.New(t)
	time11, _ := time.Parse(timeTemplate, "16.07.2021 08:00:00")
	time12, _ := time.Parse(timeTemplate, "16.07.2021 08:30:00")
	time21, _ := time.Parse(timeTemplate, "16.07.2021 08:01:00")
	time22, _ := time.Parse(timeTemplate, "16.07.2021 08:30:00")
	time31, _ := time.Parse(timeTemplate, "16.07.2021 08:00:00")
	time32, _ := time.Parse(timeTemplate, "16.07.2021 08:31:00")
	park1 := Parking{time11, time12}
	park2 := Parking{time21, time22}
	park3 := Parking{time31, time32}
	assert.Equal(park1.Equal(park1), true, "Parking time should be equal")
	assert.Equal(park1.Equal(park2), false, "Parking time should not be equal")
	assert.Equal(park1.Equal(park3), false, "Parking time should not be equal")

}

func TestParkingSequenceEqual(t *testing.T) {
	assert := assert.New(t)
	time11, _ := time.Parse(timeTemplate, "16.07.2021 08:00:00")
	time12, _ := time.Parse(timeTemplate, "16.07.2021 08:30:00")
	time21, _ := time.Parse(timeTemplate, "16.07.2021 08:01:00")
	time22, _ := time.Parse(timeTemplate, "16.07.2021 08:30:00")
	time31, _ := time.Parse(timeTemplate, "16.07.2021 08:00:00")
	time32, _ := time.Parse(timeTemplate, "16.07.2021 08:31:00")
	park1 := Parking{time11, time12}
	park2 := Parking{time21, time22}
	park3 := Parking{time31, time32}

	seq1 := ParkingSeq{park1, park2}
	seq2 := ParkingSeq{park1, park2, park3}
	seq3 := ParkingSeq{park2, park3}

	assert.Equal(seq1.Equal(seq1), true, "ParkingSeq time should be equal")
	assert.Equal(seq1.Equal(seq2), false, "ParkingSeq time should not be equal")
	assert.Equal(seq1.Equal(seq3), false, "ParkingSeq time should not be equal")

}

func TestGetParkingSequence(t *testing.T) {
	assert := assert.New(t)
	inData, err := os.Open("../data/test1.txt")
	if err != nil {
		panic(err)
	}
	ps, _ := GetParkingSequence(inData)
	assert.Equal(ps.Equal(orderTestDate), true, "ParkingSeq time should be equal")
}

func TestTwoMaxPlace(t *testing.T) {
	assert := assert.New(t)
	inData, err := os.Open("../data/test1.txt")
	if err != nil {
		panic(err)
	}
	ps, _ := GetParkingSequence(inData)
	assert.Equal(ps.GetMaxParkingPlace(), 2, "Max plase should be equal 2")
}

func TestOnePlace(t *testing.T) {
	assert := assert.New(t)
	inData, err := os.Open("../data/test2.txt")
	if err != nil {
		panic(err)
	}
	ps, _ := GetParkingSequence(inData)
	assert.Equal(ps.GetMaxParkingPlace(), 1, "Max plase should be equal 1")
}

func TestOneSeparatePlace(t *testing.T) {
	assert := assert.New(t)
	inData, err := os.Open("../data/test4.txt")
	if err != nil {
		panic(err)
	}
	ps, _ := GetParkingSequence(inData)
	assert.Equal(ps.GetMaxParkingPlace(), 1, "Max plase should be equal 1")
}

func TestOneIntersectingPlace(t *testing.T) {
	assert := assert.New(t)
	inData, err := os.Open("../data/test5.txt")
	if err != nil {
		panic(err)
	}
	ps, _ := GetParkingSequence(inData)
	assert.Equal(ps.GetMaxParkingPlace(), 1, "Max plase should be equal 1")
}

func TestZeroPlace(t *testing.T) {
	assert := assert.New(t)
	inData, err := os.Open("../data/test3.txt")
	if err != nil {
		panic(err)
	}
	ps, _ := GetParkingSequence(inData)
	assert.Equal(ps.GetMaxParkingPlace(), 0, "Max plase should be equal 0")
}

func TestThreePlace(t *testing.T) {
	assert := assert.New(t)
	inData, err := os.Open("../data/test0.txt")
	if err != nil {
		panic(err)
	}
	ps, _ := GetParkingSequence(inData)
	assert.Equal(ps.GetMaxParkingPlace(), 3, "Max plase should be equal 3")
}

// func TestPanic(t *testing.T) {
// 	assert := assert.New(t)
// 	inData, err := os.Open("../data/test_ozon1.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	ps, _ := GetParkingSequence(inData)
// 	assert.Panics(func() { ps.GetMaxParkingPlace() }, "invalid data")
// }

func TestMemory(t *testing.T) {
	assert := assert.New(t)
	inData, err := os.Open("../data/test_ozon2.txt")
	if err != nil {
		panic(err)
	}
	ps, _ := GetParkingSequence(inData)
	assert.Equal(ps.GetMaxParkingPlace(), 2, "Max plase should be equal 2")
}

func TestMemory2(t *testing.T) {
	assert := assert.New(t)
	inData, err := os.Open("../data/test6.txt")
	if err != nil {
		panic(err)
	}
	ps, _ := GetParkingSequence(inData)
	assert.Equal(ps.GetMaxParkingPlace(), 2, "Max plase should be equal 2")
}
