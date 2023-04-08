// Filename: gohttp/client.
//go-http/client
package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error){

	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	//add the custom/client-provided jheaders
	request.Header = c.getHeaders(headers)

	if err != nil{
		return nil, errors.New("ubable to create a new request")
	}

	return client.Do(request)
}

func (c* httpClient) getHeaders(requestHeaders http.Header) http.Header{

	//create map of headers
	result := make(http.Header)

	//add the client headers
	for key,value := range c.Headers{
		if len(value)>0{
			result.Set(key,value[0])
		}
	}

	//add the request headers
	for key,value := range requestHeaders{
		if len(value)>0{
			result.Set(key,value[0])
		}
	}

	return result;
}