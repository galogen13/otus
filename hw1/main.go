package main

import (
	"fmt"
	"os"
	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot get the time: %s\n", err)
        os.Exit(1)
	}
	layout := "02.01.2006 15:04:05 (MST)"
	fmt.Printf("Current Local Time: %s", time.Local().Format(layout))
}
