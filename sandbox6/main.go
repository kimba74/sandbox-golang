package main

import (
	"flag"
	"fmt"
)

var (
	blubber string
	blabber string
)

func main() {
	flag.Parse()

	fmt.Println("This is the main() function")
	fmt.Printf("Flag 'blubber': %s\n", blubber)
	fmt.Printf("Flag 'blabber': %s\n", blabber)
}

func init() {
	fmt.Println("This is the init() function")
	flag.StringVar(&blubber, "blubber", "yo", "This is the 'blubber' test flag")
	flag.StringVar(&blabber, "blabber", "ya", "This is the 'blabber' test flag")
}
