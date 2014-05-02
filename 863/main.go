package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type country struct {
	name     string
	from     int
	to       int
	overlaps []*country
}

func main() {
	fileName := os.Args[1]
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var countries []*country
	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			tokens := strings.Split(line, " ")
			fromTo := strings.Split(tokens[1], "-")
			fromMD := strings.Split(fromTo[0], "/")
			toMD := strings.Split(fromTo[1], "/")
			if len(fromMD[1]) == 1 {
				fromMD[1] = "0" + fromMD[1]
			}
			if len(toMD[1]) == 1 {
				toMD[1] = "0" + toMD[1]
			}
			from, err := strconv.Atoi(fromMD[0] + fromMD[1])
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(toMD[0] + toMD[1])
			if err != nil {
				panic(err)
			}
			countries = append(countries, &country{name: tokens[0], from: from, to: to})
		}
	}
	for i, l := 0, len(countries); i < l; i++ {
		for j := i + 1; j < l; j++ {
			c1 := countries[i]
			c2 := countries[j]
			if c1.to < c2.from || c2.to < c1.from {
				continue
			}
			c1.overlaps = append(c1.overlaps, c2)
			c2.overlaps = append(c2.overlaps, c1)
		}
	}
	for {
		var mostOverlappedCountry *country
		for _, c := range countries {
			if mostOverlappedCountry == nil {
				mostOverlappedCountry = c
				continue
			}
			if len(mostOverlappedCountry.overlaps) < len(c.overlaps) {
				mostOverlappedCountry = c
			}
		}
		if len(mostOverlappedCountry.overlaps) == 0 {
			break
		}
		for _, c := range mostOverlappedCountry.overlaps {
			for i, oc := range c.overlaps {
				if oc == mostOverlappedCountry {
					c.overlaps = append(c.overlaps[:i], c.overlaps[i+1:]...)
					break
				}
			}
		}
		for i, c := range countries {
			if c == mostOverlappedCountry {
				countries = append(countries[:i], countries[i+1:]...)
				break
			}
		}
	}
	for i, c := range countries {
		if i == 0 {
			fmt.Printf("%d", len(countries))
		}
		fmt.Printf(" %s", c.name)
	}
	fmt.Print("\n")
	//for _, c := range countries {
	//	fmt.Printf("%s %d %d\n", c.name, c.from, c.to)
	//}
}
