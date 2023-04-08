// Filename: gohttp/client.go
package gohttp

import (
	"net/http"
)

//create a client type; your own client
//client will do get post, patch, put, delete

// sterp1: create an interface
type HttpClient interface {
	Get(url string, headers http.Header)(*http.Response, error)
	Post(url string, headers http.Header, body interface{})(*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)

	SetHeaders(headers http.Header)
}

//create a concrete base on the interface type
type httpClient struct{
	Headers http.Header;
}

func (c *httpClient) SetHeaders(headers http.Header){
	c.Headers = headers
}

func (c* httpClient) Get(url string, headers http.Header)(*http.Response, error){
	return c.do(http.MethodGet,url,headers,nil)

}

func (c* httpClient) Post(url string, headers http.Header, body interface{})(*http.Response, error){
	return c.do(http.MethodPost,url,headers,body)

}

func (c* httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodPut,url,headers, body)


}

func (c* httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodPatch,url,headers,body)

}

func (c* httpClient) Delete(url string, headers http.Header) (*http.Response, error){
	return c.do(http.MethodDelete,url,headers,nil)

}

//make httpClient implement the HttpClient interface
func New() HttpClient{
	client := &httpClient{}

	return client
	
}