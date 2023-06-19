package kdtree

import "math"

const R float64 = 6371e3 // Radius of the earth, m

//Point holds geo coordinates: [X, Y] <=> [Longitude, Latitude]
type Point [2]float64

func (point Point) Lon() float64 {
	return point[0] // X
}

func (point Point) Lat() float64 {
	return point[1] // Y
}

func (point Point) Distance(p Point) float64 {
	dLat := convertToRadians(p.Lat() - point.Lat())
	dLon := convertToRadians(p.Lon() - point.Lon())
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(convertToRadians(point.Lat()))*math.Cos(convertToRadians(p.Lat()))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return c * R
}

func convertToRadians(deg float64) float64 {
	return deg * math.Pi / 180
}
