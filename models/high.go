package models

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAll() (string, error) {
	url := "https://www.data.jma.go.jp/developer/xml/feed/regular.xml"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to get disaster information")
		return "", err
	}
	defer resp.Body.Close()

	xmlStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read disaster information")
		return "", err
	}

	f := Feed{}
	xml.Unmarshal(xmlStr, &f)

	jsonBytes, err := json.Marshal(f)
	if err != nil {
		fmt.Println("failed to marshal disaster information")
		return "", err
	}
	fmt.Println(string(jsonBytes))
	return string(jsonBytes), nil
}

type Feed struct {
	MLName   xml.Name `xml:"feed"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
	Updated  string   `xml:"updated"`
	Link     []Link   `xml:"link"`
	Id       string   `xml:"id"`
	Rights   string   `xml:"rights"`
	Entry    []Entry  `xml:"entry"`
}

type Entry struct {
	MLName  xml.Name `xml:"entry"`
	Title   string   `xml:"title"`
	Id      string   `xml:"id"`
	Updated string   `xml:"updated"`
	Author  string   `xml:"author"`
	Link    Link     `xml:"link"`
	Content string   `xml:"content"`
}

type Author struct {
	MLName xml.Name `xml:"author"`
	Name   string   `xml:"name"`
}

type Link struct {
	MLName    xml.Name `xml:"link"`
	XmlnsRel  string   `xml:"rel,attr"`
	XmlnsHref string   `xml:"href,attr"`
	XmlnsType string   `xml:"type,attr"`
}
