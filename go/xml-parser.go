package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc        string `xml:"loc"`
	Lastmod    string `xml:"lastmod"`
	Changefreq string `xml:"changefreq"`
	Priority   string `xml:"priority"`
}

func xmlparser() {
	// Open the XML file
	file, err := os.Open("ayosehat.xml")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the XML
	var urlset URLSet
	err = xml.NewDecoder(file).Decode(&urlset)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(len(urlset.URLs))
}

/*
func run_xml_parser() {
	// Open the XML file
	file, err := os.Open("ayosehat.xml")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the XML
	var urlset URLSet
	err = xml.NewDecoder(file).Decode(&urlset)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var (
		seedurl string
		task    string
	)

	conn, err := connect_pgpool()
	if err != nil {
		fmt.Println(err)
	}

	// Print the loc element
	for _, urld := range urlset.URLs {
		fmt.Println("Loc:", urld.Loc)

		task = `AYOSEHAT-XML`
		// seedurl = "https://www.azlyrics.com/a.html/"
		seedurl = urld.Loc

		go func() {
			Q := NewQueue()
			Q.Enqueue(seedurl)

			line := 0
			for !Q.IsEmpty() {
				line++
				fmt.Printf("%#v.\n", line)
				u := Q.Dequeue()    // Get a URL from Q
				du, err := Fetch(u) // Fetch its HTML text
				if err != nil {
					panic(err)
				}

				if strings.TrimSpace(du) != "" { // If the HTML document is not empty
					err = StoreD(du, u, task, conn) // Store it in D
					if err != nil {
						// panic(err)
						fmt.Println(err)
					}

					L := ExtractURL(u, du) // Extract all "clean" hrefs from d(u)

					for _, v := range L {
						StoreE(u, v, task, conn)

						isContainsD, err := ContainsD(v, conn)
						if err != nil {
							// panic(err)
							fmt.Println(err)
						}

						if !Q.Contains(v) && !isContainsD {
							Q.Enqueue(v)
						}
					}
				}

				if line == 1000 {
					fmt.Println("breaking 1000 lines")
					break
				}

				if line == 1 {
					fmt.Println("single link only")
					break
				}
			}
		}()
	}

	// Wait for kill signal of channel
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("before quit")
	// This blocks until a signal is passed into the quit channel
	<-quit
	fmt.Println("after quit")
}
*/
