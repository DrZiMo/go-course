package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Print(websites)
	fmt.Print(websites["Google"])

	websites["Linkedin"] = "https://linkedin.com"

	fmt.Print(websites["Linkedin"])

	// using make function
	courseRating := make(floatMap, 3)

	courseRating["Go"] = 4.7
	courseRating["React"] = 4.6
	courseRating["JavaScript"] = 5

	courseRating.output()
}
