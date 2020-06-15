package commands

import (
	"encoding/json"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/spf13/cobra"

	"github.com/transloc/ridefactory/internal/api"
	"github.com/transloc/ridefactory/pkg/ridefactory"
)

var apiURL string
var apiToken string
var agencyName string
var nPosts int
var timeHorizon int

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "POST Rides to OnDemand",
	Run: func(cmd *cobra.Command, args []string) {
		rides := ridefactory.GetRides(nPosts)

		if timeHorizon == 0 {
			postRides(rides)
		} else {
			postRidesWithDelay(rides, timeHorizon)
		}
	},
}

func init() {
	localOnDemand := "http://localhost.transloc.com:8080"
	rootCmd.AddCommand(postCmd)
	postCmd.Flags().StringVarP(&apiURL, "url", "u", localOnDemand, "API URL to use")
	postCmd.Flags().StringVarP(&apiToken, "token", "t", "", "auth token to use")

	postCmd.Flags().IntVarP(&nPosts, "nrides", "n", 1, "number of rides to create")
	postCmd.Flags().IntVarP(&timeHorizon, "duration", "d", 0, "duration of ride submission time window, 0 to post all at once")

	postCmd.MarkFlagRequired("token")
}

func getAPI() *api.API {
	return &api.API{URL: apiURL, Token: apiToken, Agency: "imperialdemo"}
}

func postRides(rides []*ridefactory.Ride) {
	var wg sync.WaitGroup
	api := getAPI()

	for _, ride := range rides {
		wg.Add(1)
		go func(ride *ridefactory.Ride) {
			defer wg.Done()
			rideStr, _ := json.Marshal(ride)
			api.PostRide(rideStr)
		}(ride)
	}
	wg.Wait()
}

func getSleeps(n, t int) []int {
	rand.Seed(time.Now().UnixNano())

	step := t / n
	sleeps := make([]int, 0)
	min := step * -1
	max := step

	for i := 0; i < t; i += step {
		r := rand.Intn(max-min+1) + min
		sleep := i + r
		if sleep < 0 {
			sleep = 0
		}
		sleeps = append(sleeps, sleep)
	}

	sort.Ints(sleeps)
	return sleeps
}

func postRidesWithDelay(rides []*ridefactory.Ride, horizon int) {
	api := getAPI()
	sleeps := getSleeps(len(rides), horizon)

	for i, ride := range rides {
		time.Sleep(time.Duration(sleeps[i]) * time.Second)
		rideStr, _ := json.Marshal(ride)
		api.PostRide(rideStr)
	}
}
