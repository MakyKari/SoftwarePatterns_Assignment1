package main

import "fmt"

type Links struct {
	links map[string]string
}

var LinkMap *Links

func (l *Links) displayAllLinks() {
	for k, v := range l.links {
		fmt.Println("Path of " + k + ": " + v)
	}
}

func getInstance() *Links {
	if LinkMap == nil {
		LinkMap = &Links{}
	}

	return LinkMap
}

func main(){
	getInstance()
	LinkMap.links = make(map[string]string)
	LinkMap.links["google"] = "https://www.google.com"
	LinkMap.links["twitter"] = "https://twitter.com"
	LinkMap.links["singleton"] = "https://refactoring.guru/ru/design-patterns/singleton/"
	LinkMap.links["moodle"] = "https://moodle.astanait.edu.kz/"
	LinkMap.links["spotify"] = "https://open.spotify.com/"

	LinkMap.displayAllLinks()
	
	getInstance()
	fmt.Println()
	fmt.Println("Trying to create new instance")
	LinkMap.displayAllLinks()
}