package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/ducknificient/web-intelligence/go/datastore"
	"github.com/ducknificient/web-intelligence/go/entity"
	"github.com/ducknificient/web-intelligence/go/logger"
)

type BasicCrawling struct {
	logger    logger.Logger
	SeedURL   string
	Task      string
	Datastore datastore.Datastore
	IsStop    bool
}

func NewCrawler(datastore datastore.Datastore, logger logger.Logger) (c *BasicCrawling) {
	return &BasicCrawling{
		Datastore: datastore,
		IsStop:    false,
		logger:    logger,
	}

}

// func (c *BasicCrawling) CustomCrawlLogic() (isstop bool) {
// 	isstop = false

// 	return isstop
// }

func (c *BasicCrawling) Crawling(seedurl string, task string) (err error) {

	var (
		errMsg string
		du     string
	)

	Q := NewQueue()
	Q.Enqueue(seedurl)

	line := 0
	for !Q.IsEmpty() {
		line++

		u := Q.Dequeue() // Get a URL from Q
		c.logger.CrawlLog(fmt.Sprintf("%#v. %#v\n", line, u))

		fr, err := c.Fetch(u) // Fetch results
		if err != nil {
			errMsg = fmt.Sprintf("Unable to fetch. task:{%#v} .seedurl:{%#v} .err: %#v .u: %#v", seedurl, task, err.Error(), u)
			err = errors.New(errMsg)
			c.logger.CrawlError(errMsg)
			return err
		}

		if fr.Header == `pdf` {

			// store

			err = c.StorePdf(fr.PdfFilename, u, fr.PdfFile) // Store it in D
			if err != nil {
				errMsg = fmt.Sprintf("Unable to storeD. task:{%#v} .seedurl:{%#v} .err: %#v .u:{%#v} .du:{%#v} ", seedurl, task, err.Error(), u, du)
				err = errors.New(errMsg)
				c.logger.CrawlError(errMsg)
				return err
			}

		} else if fr.Header == `html` {

			du = fr.HTMLText

			if strings.TrimSpace(du) != "" { // If the HTML document is not empty
				err = c.StoreD(du, u) // Store it in D
				if err != nil {
					errMsg = fmt.Sprintf("Unable to storeD. task:{%#v} .seedurl:{%#v} .err: %#v .u:{%#v} .du:{%#v} ", seedurl, task, err.Error(), u, du)
					err = errors.New(errMsg)
					c.logger.CrawlError(errMsg)
					return err
				}

				// check apakah .pdf atau bukan

				L, err := c.ExtractURL(u, du) // Extract all "clean" hrefs from d(u)
				if err != nil {
					errMsg = fmt.Sprintf("Unable to extract url. task:{%#v} .seedurl:{%#v} .err: %#v .url:{%#v} .content: {%#v} ", seedurl, task, err.Error(), u, du)
					err = errors.New(errMsg)
					c.logger.CrawlError(errMsg)
					return err
				}

				for _, v := range L {
					c.StoreE(u, v)

					isContainsD, err := c.ContainsD(v)
					if err != nil {
						errMsg = fmt.Sprintf("Unable to check ContainsD. task:{%#v} .seedurl:{%#v} .err: %#v .u:{%#v} . ", seedurl, task, err.Error(), u)
						err = errors.New(errMsg)
						c.logger.CrawlError(errMsg)
						return err
					}

					if !Q.Contains(v) && !isContainsD {
						// fmt.Printf("enqued. %#v, %#v,%#v\n ", v, !Q.Contains(v), !isContainsD)
						msg := fmt.Sprintf("enqued. %#v, %#v,%#v\n ", v, !Q.Contains(v), !isContainsD)
						c.logger.CrawlLog(msg)
						Q.Enqueue(v)
					} else {
						msg := fmt.Sprintf("not enqued. %#v, %#v,%#v\n ", v, !Q.Contains(v), !isContainsD)
						c.logger.CrawlLog(msg)
					}
				}
			}
		}

		if line == 1000 {
			break
		}

		if c.IsStop {
			fmt.Println("stopping from controller")
			c.logger.CrawlLog("stopping crawl")
			break
		}

		// if line == 1 {
		// 	fmt.Println("testing cloudflare")
		// 	break
		// }

	}

	return err
}

func (c *BasicCrawling) StartCrawling() (err error) {
	c.IsStop = true

	return err
}

func (c *BasicCrawling) StopCrawling() (err error) {
	c.IsStop = false

	return err
}

func (c *BasicCrawling) TestCrawling() {

	for {

		fmt.Println("controller starting")

		if !c.IsStop {
			fmt.Println("stopping crawl")
			break
		}

		time.Sleep(1000 * time.Microsecond)
	}

}

func (c *BasicCrawling) Fetch2(url string) (htmltext entity.FetchResult, err error) {

	// Get content from URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error http get")
		return htmltext, err
	}
	defer resp.Body.Close()

	// Read HTML content
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error read all")
		return htmltext, err
	}

	// check apakah pdf atau bukan
	// if strings.Contains(resp.Header.Get("Content-Type"), "pdf") {

	// 	fetchres.Header = `pdf`
	// 	fetchres.PdfFile = respbody
	// 	fetchres.PdfFilename = url

	// } else {

	// 	fetchres.Header = resp.Header.Get("Content-Type")
	// 	fetchres.HTMLText = string(respbody)
	// }

	// fetchres.HTMLText = string(respbody)

	htmltext.HTMLText = string(respbody)

	return htmltext, err
}

func (c *BasicCrawling) Fetch(url string) (fetchres entity.FetchResult, err error) {

	// Get content from URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error http get")
		return fetchres, err
	}
	defer resp.Body.Close()

	// Read HTML content
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error read all")
		return fetchres, err
	}

	// check apakah pdf atau bukan
	if strings.Contains(resp.Header.Get("Content-Type"), "pdf") {

		fetchres.Header = `pdf`
		fetchres.PdfFile = respbody
		fetchres.PdfFilename = url

	} else if strings.Contains(resp.Header.Get("Content-Type"), "html") ||
		strings.Contains(resp.Header.Get("Content-Type"), "text") {

		fetchres.Header = `html`
		fetchres.HTMLText = string(respbody)
	}

	fetchres.HTMLText = string(respbody)

	return fetchres, err
}

func (c *BasicCrawling) ExtractURL(inputurl string, html string) (filteredHrefs []string, err error) {
	// Parse base URL
	base, err := url.Parse(inputurl)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to parse url. %#v. err: %#v", inputurl, err.Error())
		err = errors.New(errMsg)
		c.logger.Error(errMsg)
		return filteredHrefs, err
	}
	baseHost := base.Host
	// baseHost += `/`

	// Regex pattern to find href attributes
	hrefPattern := `href=["'](.*?)["']`

	// Find all matches of href attributes
	matches := regexp.MustCompile(hrefPattern).FindAllStringSubmatch(html, -1)

	// Array to store filtered href attributes
	filteredHrefs = []string{}

	c.logger.CrawlLog(fmt.Sprintf("found %#v href \n", len(matches)))

	// Iterate through all matches
	for a, match := range matches {
		href := match[1]
		c.logger.CrawlLog(fmt.Sprintf("%#v. href : %#v ,", a, href))

		// Parse href URL
		parsedHref, err := url.Parse(href)
		if err != nil {
			continue
		}
		c.logger.CrawlLog(fmt.Sprintf("parsed href : %#v \n", parsedHref.Path))

		// If the found URL has a host and the host is a subdomain of the base host
		if parsedHref.Host != "" {
			if strings.Contains(parsedHref.Host, baseHost) {
				if !strings.HasPrefix(href, "http") {
					href = "https:" + href
				}

				if !strings.Contains("/#", href) || !strings.Contains("mailto", href) {
					c.logger.CrawlLog(fmt.Sprintf("exists: %#v\n", href))
					filteredHrefs = append(filteredHrefs, href)
				} else {
					c.logger.CrawlLog(fmt.Sprintf("%#v. host != 0, href : %#v ,", a, href))
				}
			}
		} else {
			// If the found URL does not have a subdomain, add the subdomain

			if len(href) > 0 {
				if href[0:1] != `/` {
					href = `/` + href
				}
			}

			newHref := "https://" + baseHost + href

			if !strings.Contains("/#", newHref) || !strings.Contains("mailto", newHref) {
				c.logger.CrawlLog(fmt.Sprintf("exists new: %#v\n", newHref))
				filteredHrefs = append(filteredHrefs, newHref)
			} else {
				// fmt.Printf("href: %#v\n", newHref)
				c.logger.CrawlLog(fmt.Sprintf("%#v. host == 0, href : %#v ,", a, href))
			}

		}
	}

	return filteredHrefs, err
}

func (c *BasicCrawling) StoreD(pagesource string, link string) (err error) {

	err = c.Datastore.StoreD(pagesource, link, c.Task)
	if err != nil {
		return err
	}

	return err
}

func (c *BasicCrawling) StorePdf(filename string, link string, pdf []byte) (err error) {

	err = c.Datastore.StorePdf(filename, link, c.Task, pdf)
	if err != nil {
		return err
	}

	return err
}

func (c *BasicCrawling) StoreE(link string, href string) (err error) {

	err = c.Datastore.StoreE(link, href, c.Task)
	if err != nil {
		return err
	}

	return err
}

func (c *BasicCrawling) ContainsD(link string) (contains bool, err error) {

	contains, err = c.Datastore.ContainsD(link)
	if err != nil {
		return false, err
	}

	return contains, err
}

func (c *BasicCrawling) CrawlpageList(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListData, err error) {

	fmt.Println(param)
	// var dataList []entity.CrawlhrefListData
	dataList, err = c.Datastore.CrawlpageList(param)
	if err != nil {
		return dataList, err
	}

	return dataList, err
}
