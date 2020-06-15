package commands

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/transloc/ridefactory/pkg/ridefactory"
)

var nRides int

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Ride as JSON",
	Run: func(cmd *cobra.Command, args []string) {
		rides := ridefactory.GetRides(nRides)
		writeRides(rides)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().IntVarP(&nRides, "nrides", "n", 1, "number of rides to create")
}

func writeRides(rides []*ridefactory.Ride) {
	for _, ride := range rides {
		rideStr, _ := json.Marshal(ride)
		fmt.Println(string(rideStr))
	}
}
