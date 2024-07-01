package main

import (
	"fmt"
	"net/http"
)
 
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://xxsfdfsd.com",
	}

	c := make(chan string)


	for _, link := range links {
		go checkLink(link, c)
	}

	// for i:=0; i < len(links); i++ {
	// 	fmt.Println(<-c, "finished checking")
	// }

	for l := range c {
		fmt.Println(l, "finished checking")
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- link 
		fmt.Println(link, "might be down!")
		return
	}
	c <- link 
	fmt.Println(link, " is up")
}