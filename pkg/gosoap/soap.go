// Massively simplified version of https://github.com/tiaguinho/gosoap
package gosoap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

const Doctype = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"

// SoapClient return new *Client to handle the requests with the WSDL
func SoapClient(apiUrl string) (*Client, error) {
	_, err := url.Parse(apiUrl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		URL: strings.TrimSuffix(apiUrl, "/"),
	}

	return c, nil
}

// Client struct hold all the informations about WSDL,
// request and response of the server
type Client struct {
	HttpClient *http.Client
	URL        string
	HeaderName string
	Body       []byte
	payload    []byte
}

// GetLastRequest returns the last request
func (c *Client) GetLastRequest() []byte {
	return c.payload
}

// Call call's the method m with Params p
func (c *Client) Call(p interface{}) (err error) {
	r := RequestSoapEnvelope{
		Xsi:  "http://www.w3.org/2001/XMLSchema-instance",
		Xsd:  "http://www.w3.org/2001/XMLSchema",
		Soap: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: RequestSoapBody{
			Content: p,
		},
	}

	c.payload, err = xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return err
	}
	c.payload = []byte(Doctype + string(c.payload))
	// fmt.Println(string(c.payload))
	b, err := c.doRequest()
	if err != nil {
		return err
	}

	var soap ResponseSoapEnvelope
	// err = xml.Unmarshal(b, &soap)
	// error: xml: encoding "ISO-8859-1" declared but Decoder.CharsetReader is nil
	// https://stackoverflow.com/questions/6002619/unmarshal-an-iso-8859-1-xml-input-in-go
	// https://github.com/golang/go/issues/8937

	decoder := xml.NewDecoder(bytes.NewReader(b))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&soap)

	c.Body = soap.Body.Contents

	return err
}

// Unmarshal get the body and unmarshal into the interface
func (c *Client) Unmarshal(v interface{}) error {
	if len(c.Body) == 0 {
		return fmt.Errorf("Body is empty")
	}
	// fmt.Println(string(c.Body))
	var f ResponseSoapFault
	xml.Unmarshal(c.Body, &f)
	if f.Code != "" {
		return fmt.Errorf("[%s]: %s", f.Code, f.Description)
	}

	return xml.Unmarshal(c.Body, v)
}

// doRequest makes new request to the server using c.URL and the body.
// body is enveloped in Call method
func (c *Client) doRequest() ([]byte, error) {
	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(c.payload))
	if err != nil {
		return nil, err
	}

	if c.HttpClient == nil {
		c.HttpClient = &http.Client{
			Timeout: time.Second * 60,
		}
	}

	req.ContentLength = int64(len(c.payload))

	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("Accept", "text/xml")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
