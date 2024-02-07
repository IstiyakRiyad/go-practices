package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"errors"
)

// Define directly
var n = flag.Int("n", 0, "The value of n")
var species = flag.String("species", "gopher", "The species we are studying")
var bool = flag.Bool("bool", true, "This is a boolean value")

// Multiple flag acts on same value 
var gopherType string
func initMultiFlagSameValue() {
	const (
		defaultGopher = "pocket"
		usage		= "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage + " (shorthand)")

}


// Custom Flag types, Multiple array like input to flag
// Custom Variable must have two user defined function called, String() & Set(string)
type interval []time.Duration

func (i *interval) String() string {
	return fmt.Sprint(*i)
}

func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}

	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

var intervalFlag interval

func initIntervalFlag() {
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}


func main() {
	initMultiFlagSameValue()
	initIntervalFlag()

	fmt.Println("This will be printed before --help configurations because it comes before flag.Parse() function")
	flag.Parse()


	fmt.Println(
		"The value of n: ", *n, 
		"\nSpecies: ", *species, 
		"\ngopherType: ", gopherType, 
		"\nInterval flag: ", intervalFlag,
		"\nBoolean flag: ", *bool,
	)
}

