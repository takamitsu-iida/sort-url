package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/takamitsu-iida/radix"
)

type domainList struct {
	XMLName  xml.Name  `xml:"domainlist"`
	Services []service `xml:"service"`
}

func NewDomainList() *domainList {
	return &domainList{}
}

func (dl *domainList) LoadFromXmlFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(data, dl)
	if err != nil {
		return err
	}

	return nil
}

type service struct {
	XMLName xml.Name `xml:"service"`
	Name    string   `xml:"name,attr"`
	Dst     dst      `xml:"dst"`
}

type dst struct {
	XMLName xml.Name `xml:"dst"`
	Url     []string `xml:"url"`
}

/*
<?xml version="1.0" encoding="utf-8"?>
<domainlist>
  <service name="MS-Apps">
    <dst>
      <url>outlook.office.com</url>
      <url>outlook.office365.com</url>
	  ...
    </dst>
  </service>
</domainlist>
*/

func main() {
	DOMAINS_XML_PATH := filepath.Join("./data", "domains.xml")
	CUSTOM_XML_PATH := filepath.Join("./data", "custom.xml")

	domainList := NewDomainList()
	err := domainList.LoadFromXmlFile(DOMAINS_XML_PATH)
	if err != nil {
		log.Fatal(err)
	}

	// extract url in "MS-Apps"
	urls := []string{}
	for _, service := range domainList.Services {
		name := service.Name
		if name == "MS-Apps" {
			urls = append(urls, service.Dst.Url...)
		}
	}

	customList := NewDomainList()
	err = customList.LoadFromXmlFile(CUSTOM_XML_PATH)
	if err != nil {
		log.Fatal(err)
	}

	// extract url in "MS-Apps-Custom"
	for _, service := range customList.Services {
		name := service.Name
		if name == "MS-Apps-Custom" {
			urls = append(urls, service.Dst.Url...)
		}
	}

	// create radix tree
	r := radix.New()

	duplicated := []string{}

	for i, url := range urls {
		u := reverse(url)
		inserted := r.Insert(u, i)
		if inserted == false {
			duplicated = append(duplicated, url)
		}
	}

	// dump the tree
	r.Walk(func(s string, v interface{}) bool {
		fmt.Println(s, v)
		return false
	})

	// check duplicated
	if len(duplicated) > 0 {
		fmt.Println("===duplicate===")
		for _, dup := range duplicated {
			fmt.Println(dup, " is duplicated.")
		}
	}
}

// reverse url string
// for example: outlook.office.com -> com.office.outlook
func reverse(url string) string {
	arr := strings.Split(url, ".")

	// reverse the arr
	n := len(arr)
	for i := 0; i < (n / 2); i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}

	return strings.Join(arr, ".")
}
