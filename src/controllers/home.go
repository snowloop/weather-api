package controllers

import "net/http"

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the main page of the weather api!</h1><a href=\"https://www.linkedin.com/in/alexandrejosse/\" target=\"_blank\">My Linkedin</a>"))
	w.Header().Set("Content-Type", "text/html")
}
