package gosoap

import (
	"encoding/xml"
)

// SoapEnvelope struct
type ResponseSoapEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  ResponseSoapHeader
	Body    ResponseSoapBody
}

// SoapHeader struct
type ResponseSoapHeader struct {
	XMLName  xml.Name `xml:"Header"`
	Contents []byte   `xml:",innerxml"`
}

// SoapBody struct
type ResponseSoapBody struct {
	XMLName  xml.Name `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}

// Fault response
type ResponseSoapFault struct {
	Code        string `xml:"faultcode"`
	Description string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}

type RequestSoapEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Body    RequestSoapBody
}

type RequestSoapBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Content interface{}
}
