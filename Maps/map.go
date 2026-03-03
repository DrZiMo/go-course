package main

import "fmt"

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
	courseRating := make(map[string]float64, 3)

	courseRating["Go"] = 4.7
	courseRating["React"] = 4.6
	courseRating["JavaScript"] = 5
}
