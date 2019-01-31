package gosoap

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
)

var tokens []xml.Token

// MarshalXML envelope the body and encode to xml
func (c Client) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {

	tokens = []xml.Token{}

	//start envelope
	if c.Definitions == nil {
		return fmt.Errorf("definitions is nil")
	}

	startEnvelope()
	if len(c.HeaderParams) > 0 {
		startHeader(c.HeaderName, c.Definitions.Types[0].XsdSchema[0].TargetNamespace)
		for k, v := range c.HeaderParams {
			t := xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: k,
				},
			}

			tokens = append(tokens, t, xml.CharData(v), xml.EndElement{Name: t.Name})
		}

		endHeader(c.HeaderName)
	}

	err := startBody(c.Method, c.Definitions.Types[0].XsdSchema[0].TargetNamespace)
	if err != nil {
		return err
	}

	recursiveEncode(c.Params)

	//end envelope
	endBody(c.Method)
	endEnvelope()

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	return e.Flush()
}

func recursiveEncode(hm interface{}) {
	v := reflect.ValueOf(hm)

	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			t := xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: key.String(),
				},
			}

			tokens = append(tokens, t)
			recursiveEncode(v.MapIndex(key).Interface())
			tokens = append(tokens, xml.EndElement{Name: t.Name})
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			recursiveEncode(v.Index(i).Interface())
		}
	case reflect.String:
		content := xml.CharData(v.String())
		tokens = append(tokens, content)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := strconv.FormatInt(v.Int(), 10)
		tokens = append(tokens, xml.CharData(val))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		val := strconv.FormatUint(v.Uint(), 10)
		tokens = append(tokens, xml.CharData(val))
	case reflect.Float32, reflect.Float64:
		val := strconv.FormatFloat(v.Float(), 'g', -1, v.Type().Bits())
		tokens = append(tokens, xml.CharData(val))
	case reflect.Bool:
		val := strconv.FormatBool(v.Bool())
		tokens = append(tokens, xml.CharData(val))
	case reflect.Struct:
		s := indirect(v)
		sType := s.Type()

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)

			t := xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: sType.Field(i).Name,
				},
			}

			tokens = append(tokens, t)
			recursiveEncode(f.Interface())
			tokens = append(tokens, xml.EndElement{Name: t.Name})
		}
	}
}

// indirect drills into interfaces and pointers, returning the pointed-at value.
// If it encounters a nil interface or pointer, indirect returns that nil value.
// This can turn into an infinite loop given a cyclic chain,
// but it matches the Go 1 behavior.
func indirect(vf reflect.Value) reflect.Value {
	for vf.Kind() == reflect.Interface || vf.Kind() == reflect.Ptr {
		if vf.IsNil() {
			return vf
		}
		vf = vf.Elem()
	}
	return vf
}

func startEnvelope() {
	e := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
			{Name: xml.Name{Space: "", Local: "xmlns:xsd"}, Value: "http://www.w3.org/2001/XMLSchema"},
			{Name: xml.Name{Space: "", Local: "xmlns:soap"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		},
	}

	tokens = append(tokens, e)
}

func endEnvelope() {
	e := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
	}

	tokens = append(tokens, e)
}

func startHeader(m, n string) {
	h := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Header",
		},
	}

	if m == "" || n == "" {
		tokens = append(tokens, h)
		return
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns"}, Value: n},
		},
	}

	tokens = append(tokens, h, r)

	return
}

func endHeader(m string) {
	h := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Header",
		},
	}

	if m == "" {
		tokens = append(tokens, h)
		return
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens = append(tokens, r, h)
}

// startToken initiate body of the envelope
func startBody(m, n string) error {
	b := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	if m == "" || n == "" {
		return fmt.Errorf("method or namespace is empty")
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns"}, Value: n},
		},
	}

	tokens = append(tokens, b, r)

	return nil
}

// endToken close body of the envelope
func endBody(m string) {
	b := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens = append(tokens, r, b)
}
