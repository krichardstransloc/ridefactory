package ridefactory

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var PositionFactory = factory.NewFactory(
	&Position{},
).Attr("Latitude", func(args factory.Args) (interface{}, error) {
	// return getRandomLatitide(35.98, 35.99, 5)
	num := float64(randomdata.Number(9800, 9900))
	return float64(35 + (num / 10000)), nil

}).Attr("Longitude", func(args factory.Args) (interface{}, error) {
	// return getRandomLongitude(-78.91, -78.90, 5)
	num := float64(randomdata.Number(9000, 9100))
	return float64((78 + (num / 10000)) * -1), nil
})

// func getRandomLatitide(low, high float64, precision int) (float64, error) {
// 	if low < -90 || high > 90 {
// 		return 0, errors.New("latitude out of bounds")
// 	}

// 	return randomdata.Decimal(low, high, precision), nil
// }

// func getRandomLongitude(low, high float64, precision int) (float64, error) {
// 	if low < -180 || high > 180 {
// 		return 0, errors.New("longitude out of bounds")
// 	}

// 	return randomdata.Decimal(low, high, precision), nil
// }

func GetPosition() *Position {
	return PositionFactory.MustCreate().(*Position)
}
