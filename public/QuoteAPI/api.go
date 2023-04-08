package QuoteAPI

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Aldric2023/SystemsProgrammingTest/public/QuoteAPI/gohttp"
)

type application struct {
	client gohttp.HttpClient
}

func RetrieveData(api string) []byte {
	//create one instance of a client
	app := &application{
		client: gohttp.New(),
	}

	//defing our common headers (cleiet Headers)
	commonHeaders := make(http.Header)
	commonHeaders.Set("X-Api-Key", "XKDaRBFXYL/urWQZXyPYww==RNJF42IDi6rJ5Gfd")
	app.client.SetHeaders(commonHeaders)

	if(api =="quote"){
		bytes, err := app.QuoteAPI()
		
		if err != nil {
			panic(err)
		}
		return bytes;	
	}

	if(api =="greeting"){
		bytes, err := app.greetingAPI()
		
		if err != nil {
			panic(err)
		}
		return bytes;	
	}
	
	return nil;

}

func (app *application) greetingAPI() ([]byte, error) {
	headers := make(http.Header)
	headers.Set("Connection", "Keep-alive")
	url := "https://api.api-ninjas.com/v1/worldtime?city=Belmopan"

	bytes, err := app.getData(url)

	if err != nil {
		return nil,err
	}

	return bytes,nil
}

func (app *application) QuoteAPI() ([]byte, error) {
	headers := make(http.Header)
	headers.Set("Connection", "Keep-alive")
	url := "https://api.api-ninjas.com/v1/quotes?category=inspirational"

	bytes, err := app.getData(url)

	if err != nil {
		return nil,err
	}

	return bytes,nil
}

func (app *application) getData(theUrl string) ([]byte, error) {

	// Add some headers to the get request
	headers := make(http.Header)

	response, err := app.client.Get(theUrl, headers)
	if err != nil {
		return nil,err
	}

	//we need to close the response
	defer response.Body.Close()
	fmt.Println(response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil,err
	}

	return bytes,nil
}
