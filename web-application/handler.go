package main

import (
	"fmt"
	"net/http"
)

var htmlContent = `
<!DOCTYPE html>
<html>
	<head><title>%s</title></head>
	<body>%s</body>
</html>
`

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s", r.Method, r.URL)
	homeContent := fmt.Sprintf(htmlContent, "Home", "<h1>This is home</h1>")

	_, _ = w.Write([]byte(homeContent))
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	aboutContent := `<h1>This is the about page</h1><p>This is small example of the about page</p>`
	aboutCon := fmt.Sprintf(htmlContent, "About", aboutContent)

	_, _ = w.Write([]byte(aboutCon))
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	aboutContent := `<h1>This is the contact page</h1><p>This is small example of the contact page</p>`
	aboutCon := fmt.Sprintf(htmlContent, "contact", aboutContent)

	_, _ = w.Write([]byte(aboutCon))
}
