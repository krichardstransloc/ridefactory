package ridefactory

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

type Ride struct {
	ServiceID  string    `json:"service_id"`
	Rider      *Rider    `json:"rider"`
	Capacity   int       `json:"capacity"`
	Wheelchair bool      `json:"wheelchair"`
	Pickup     *RideStop `json:"pickup"`
	Dropoff    *RideStop `json:"dropoff"`
	Source     string    `json:"source"`
}

var RideFactory = factory.NewFactory(
	&Ride{ServiceID: "717", Source: "dispatcher", Wheelchair: false},
).Attr("Rider", func(args factory.Args) (interface{}, error) {
	return GetRider(), nil
}).Attr("Capacity", func(args factory.Args) (interface{}, error) {
	return randomdata.Number(1, 3), nil
}).Attr("Pickup", func(args factory.Args) (interface{}, error) {
	return GetRideStop(), nil
}).Attr("Dropoff", func(args factory.Args) (interface{}, error) {
	return GetRideStop(), nil
})

func GetRides(n int) []*Ride {
	rides := make([]*Ride, n)
	for i := 0; i < n; i++ {
		rides[i] = RideFactory.MustCreate().(*Ride)
	}
	return rides
}
