package ridefactory

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

type Rider struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

var RiderFactory = factory.NewFactory(
	&Rider{},
).Attr("FirstName", func(args factory.Args) (interface{}, error) {
	return randomdata.FirstName(2), nil
}).Attr("LastName", func(args factory.Args) (interface{}, error) {
	return randomdata.LastName(), nil
}).Attr("Phone", func(args factory.Args) (interface{}, error) {
	return randomPhoneNumber(), nil
})

func randomPhoneNumber() string {
	areaCode := fmt.Sprintf("%3d", randomdata.Number(1, 999))
	exchange := fmt.Sprintf("%3d", randomdata.Number(1, 999))
	number := fmt.Sprintf("%4d", randomdata.Number(1, 9999))
	return fmt.Sprintf("%s%s%s", areaCode, exchange, number)
}

func GetRider() *Rider {
	return RiderFactory.MustCreate().(*Rider)
}
