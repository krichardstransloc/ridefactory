package ridefactory

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

type RideStop struct {
	Address  string    `json:"address"`
	Position *Position `json:"position"`
}

var RideStopFactory = factory.NewFactory(
	&RideStop{},
).Attr("Address", func(args factory.Args) (interface{}, error) {
	return fmt.Sprintf("%d %s", randomdata.Number(100), randomdata.Street()), nil
}).Attr("Position", func(args factory.Args) (interface{}, error) {
	return GetPosition(), nil
})

func GetRideStop() *RideStop {
	return RideStopFactory.MustCreate().(*RideStop)
}
