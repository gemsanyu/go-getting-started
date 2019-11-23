package lib

import (
	"math"
)

/**
 * Created by M. Muflih Kholidin
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

func CalculateHaversineDistance(origin, destination Coordinate) float64 {
	DY := math.Abs(origin.Lat-destination.Lat) / 180 * math.Pi
	DX := math.Abs(origin.Lng-destination.Lng) / 180 * math.Pi
	Y1 := origin.Lat / 180 * math.Pi
	Y2 := destination.Lat / 180 * math.Pi
	R := 6372800.00000000 //in Meter ! it is the average great-elliptic or great-circle radius
	a := math.Sin(DY/2)*math.Sin(DY/2) + math.Cos(Y1)*math.Cos(Y2)*math.Sin(DX/2)*math.Sin(DX/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

type Coordinate struct {
	Lat float64
	Lng float64
}
