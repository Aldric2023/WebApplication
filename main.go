//tasks

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/Aldric2023/webapplication/public/QuoteAPI"
)

type UserData struct {
	PageTitle string
	Body      template.HTML
	IP        string
}

func middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//this is executed on the way down to the handeler
		log.Println("Executing middleware")
		log.Println(r.URL.Path)

		if r.URL.Path != "/" && r.URL.Path != "/random" && r.URL.Path != "/greeting" && r.URL.Path != "/phpmyadmin" && r.URL.Path != "/favicon.ico" {
			//return and dont go any further
			fmt.Println("invalid link")
			return
		}

		next.ServeHTTP(w, r)
		//this is executed on the way up
		log.Println("Executing middlewareB again")
		log.Printf("IP address: %s ", r.RemoteAddr)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	body := "<h2><p>Welcome to the test webapp<br>" +
		"Under Production<br>" +
		"</p></h2>"

	data := UserData{
		PageTitle: "About Me",
		Body:      template.HTML(body),
		IP:        r.Host,
	}

	log.Println("Url user used: " + data.IP)
	ts, _ := template.ParseFiles("public/index.html.tmpl")

	ts.Execute(w, data)

	log.Println("About working")

}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	//get random quote using API
	//call some function
	jsonData := QuoteAPI.RetrieveData("quote")

	// Create a slice to store the JSON data
	var quotes []map[string]string

	// Unmarshal the JSON data into the slice
	err := json.Unmarshal([]byte(jsonData), &quotes)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := "<h2><p>" + quotes[0]["quote"] + "</p></h2> <br>" +
		"<h3> -" + quotes[0]["author"] + "</h3> "

	data := UserData{
		PageTitle: "Here is a random quote",
		Body:      template.HTML(body),
		IP:        r.Host,
	}

	ts, _ := template.ParseFiles("public/index.html.tmpl")

	ts.Execute(w, data)

	log.Println("Quote working")

}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Greeting working")
	//get greeting using API
	jsonData := QuoteAPI.RetrieveData("greeting")

	// Unmarshal the JSON data into a map
	var greetingData map[string]string
	err := json.Unmarshal([]byte(jsonData), &greetingData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the greeting message using the extracted data
	body := "<h2><p> Today is " + greetingData["day_of_week"] + " the " + greetingData["day"] + " of " + greetingData["month"] + ", " +
		greetingData["year"] + " and the time is now " + greetingData["hour"] + ":" + greetingData["minute"] +
		":" + greetingData["second"] + "</p></h2>"

	// Pass the greeting message to the HTML template
	data := UserData{
		PageTitle: "We sometimes get lost during the Week.\n Here is a reminder",
		Body:      template.HTML(body),
		IP:        r.Host,
	}

	ts, _ := template.ParseFiles("public/index.html.tmpl")
	ts.Execute(w, data)
}

func phpmyadminHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("phpmyadmin")
	// fmt.Println("Host for phpmyadmin function: " + strings.Split(r.Host, ":")[0])
	http.Redirect(w, r, "http://"+strings.Split(r.Host, ":")[0]+":80/phpmyadmin/", http.StatusSeeOther)

}

func faviconhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("favicon is being requested")
	http.ServeFile(w, r, "public/favicon.png")
}

func main() {
	// serve static files from the "static" directory
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	//work
	mux.Handle("/", middleware(http.HandlerFunc(homeHandler)))
	mux.Handle("/random", middleware(http.HandlerFunc(randomHandler)))
	mux.Handle("/greeting", middleware(http.HandlerFunc(greetingHandler)))
	mux.Handle("/phpmyadmin", middleware(http.HandlerFunc(phpmyadminHandler)))
	mux.Handle("/favicon.ico", middleware(http.HandlerFunc(faviconhandler)))

	log.Fatal(http.ListenAndServe(":8080", mux))

}
