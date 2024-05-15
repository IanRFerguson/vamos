package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "Get the current time in different timezones",
	Long: `Takes a given American timezone and returns the current time in that area.
	
	For example:
	go run main.go timezone --tz Eastern
	go run main.go timezone --tz Pacific
	go run main.go timezone --tz Hawaiian`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get value from command line
		_timezone_, _ := cmd.Flags().GetString("tz")

		// Get timestamp value
		_timeOutput_, _error_ := getTimestamp(_timezone_)
		if _error_ != nil {
			fmt.Printf("** ERROR ** %s", _error_)
		}
		fmt.Println((_timeOutput_))
	},
}

func getTimestamp(timeInput string) (time.Time, error) {
	if timeInput == "UTC" {
		return time.Now().UTC(), nil
	} else {
		var _location_ = ""
		switch timeInput {
		case "Eastern":
			_location_ = "America/New_York"
		case "Central":
			_location_ = "America/Chicago"
		case "Mountain":
			_location_ = "America/Denver"
		case "Pacific":
			_location_ = "America/Los_Angeles"
		case "Hawaiian":
			_location_ = "Pacific/Honolulu"
		case "Alaskan":
			_location_ = "America/Anchorage"
		default:
			return time.Now().UTC(), errors.New("Invalid input " + timeInput + "!")
		}
		loc, _ := time.LoadLocation(_location_)
		return time.Now().In(loc), nil
	}
}

func init() {
	rootCmd.AddCommand(timezoneCmd)
	timezoneCmd.PersistentFlags().String(
		"tz",
		"UTC",
		`Set your desired timezone
Must be one of Eastern, Central, Mountain, Pacific, Hawaiian, Alaskan`)
}
