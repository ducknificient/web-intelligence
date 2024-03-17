package service

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

type CrawlerService interface {
	Crawling(seedurl string, task string) (err error)
	TestCrawling()
	StartCrawling() (err error)
	StopCrawling() (err error)
	CrawlpageList(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListData, err error)
	CrawlpageListParsed(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListParsedData, err error)
}

type BasicCrawling struct {
	logger logger.Logger
	// SeedURL   string
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
		Q      *Queue
	)

	c.Task = task
	// c.SeedURL = seedurl

	// existingQueue, err := c.GetExistingQueue()
	// if err != nil {
	// 	c.logger.CrawlError(errMsg)
	// 	return
	// }

	// err = c.GetLatestSeedUrl()
	// if err != nil {
	// 	c.logger.CrawlError(errMsg)
	// 	return
	// }

	// Q = NewQueueFromExisting(existingQueue)

	Q = NewQueue()

	Q.Enqueue(seedurl)

	line := 0
	for !Q.IsEmpty() {
		line++
		// url.Error
		u := Q.Dequeue() // Get a URL from Q
		c.logger.CrawlLog(fmt.Sprintf("%#v. %#v", line, u))
		// fmt.Printf("%#v. %#v", line, u)

		fr, err := c.Fetch(u) // Fetch results
		if err != nil {
			fmt.Printf("%#v", err)
			errMsg = fmt.Sprintf("Unable to fetch. task:{%#v} .seedurl:{%#v} .err: %#v .u: %#v", seedurl, task, err.Error(), u)
			c.logger.CrawlError(errMsg)
			Q.Enqueue(u)
			continue
		}

		// fmt.Printf("fetch res: %#v\n")
		// fmt.Printf(" header: %#v\n", fr.Header)
		c.logger.CrawlLog(fmt.Sprintf(" type: %#v .content type: %#v\n", fr.DocumentType, fr.DocumentContentType))

		switch fr.DocumentType {
		case "html":
			du = string(fr.DocumentFile)
			if strings.TrimSpace(du) != "" { // If the HTML document is not empty
				err = c.StoreD(du, u, fr) // Store it in D
				if err != nil {
					errMsg = fmt.Sprintf("Unable to storeD. task:{%#v} .seedurl:{%#v} .err: %#v .u:{%#v} .du:{%#v} \n", seedurl, task, err.Error(), u, du)
					err = errors.New(errMsg)
					c.logger.CrawlError(errMsg)
					return err
				}

				// check apakah .pdf atau bukan

				L, err := c.ExtractURL(u, du) // Extract all "clean" hrefs from d(u)

				c.logger.CrawlLog(fmt.Sprintf(" total href: %#v\n", len(L)))
				// fmt.Printf(" total href: %#v\n", len(L))

				if err != nil {
					errMsg = fmt.Sprintf("Unable to extract url. task:{%#v} .seedurl:{%#v} .err: %#v .url:{%#v} .content: {%#v} \n", seedurl, task, err.Error(), u, du)
					err = errors.New(errMsg)
					c.logger.CrawlError(errMsg)
					return err
				}

				for _, v := range L {
					c.StoreE(u, v)

					isContainsD, err := c.ContainsD(v)
					if err != nil {
						errMsg = fmt.Sprintf("Unable to check ContainsD. task:{%#v} .seedurl:{%#v} .err: %#v .u:{%#v} . \n", seedurl, task, err.Error(), u)
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
		default:
			err = c.StoreDocument(u, fr.DocumentType, fr.DocumentFile, fr.DocumentContentType) // Store it in D
			if err != nil {
				errMsg = fmt.Sprintf("Unable to storeD. task:{%#v} .seedurl:{%#v} .err: %#v .u:{%#v} .du:{%#v} \n", seedurl, task, err.Error(), u, du)
				err = errors.New(errMsg)
				c.logger.CrawlError(errMsg)
				return err
			}
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

	fetchres.DocumentContentType = resp.Header.Get("Content-Type")
	fetchres.DocumentFile = respbody

	if strings.Contains(fetchres.DocumentContentType, "html") {

		fetchres.DocumentType = `html`
		fetchres.DocumentContentType = resp.Header.Get("Content-Type")
		fetchres.DocumentFile = respbody

		c.logger.CrawlLog(fmt.Sprintf("url: %#v . header: %#v document type : %#v file: ", url, fetchres.DocumentContentType, fetchres.DocumentType))
	} else if strings.Contains(fetchres.DocumentContentType, "text") {

		// sementara buat semua text jadi html
		fetchres.DocumentType = `html`
		fetchres.DocumentContentType = resp.Header.Get("Content-Type")
		fetchres.DocumentFile = respbody

		c.logger.CrawlLog(fmt.Sprintf("url: %#v . header: %#v document type : %#v file: ", url, fetchres.DocumentContentType, fetchres.DocumentType))
	} else if strings.Contains(fetchres.DocumentContentType, "pdf") {

		fetchres.DocumentType = `pdf`
		fetchres.DocumentContentType = resp.Header.Get("Content-Type")
		fetchres.DocumentFile = respbody

		c.logger.CrawlLog(fmt.Sprintf("url: %#v . header: %#v document type : %#v file: ", url, fetchres.DocumentContentType, fetchres.DocumentType))
	} else if strings.Contains(fetchres.DocumentContentType, "image") {

		fetchres.DocumentType = `image`
		fetchres.DocumentContentType = resp.Header.Get("Content-Type")
		fetchres.DocumentFile = respbody

		c.logger.CrawlLog(fmt.Sprintf("url: %#v . header: %#v document type : %#v file: ", url, fetchres.DocumentContentType, fetchres.DocumentType))
	} else if strings.Contains(fetchres.DocumentContentType, "video") {
		fetchres.DocumentType = `video`
		fetchres.DocumentContentType = resp.Header.Get("Content-Type")
		fetchres.DocumentFile = respbody

		c.logger.CrawlLog(fmt.Sprintf("url: %#v . header: %#v document type : %#v file: ", url, fetchres.DocumentContentType, fetchres.DocumentType))

	} else {
		fetchres.DocumentType = `unknown`
		fetchres.DocumentContentType = resp.Header.Get("Content-Type")
		fetchres.DocumentFile = respbody

		c.logger.CrawlLog(fmt.Sprintf("url: %#v . header: %#v document type : %#v file: ", url, fetchres.DocumentContentType, fetchres.DocumentType))
	}

	return fetchres, err
}

func (c *BasicCrawling) ExtractURL(inputurl string, html string) (filteredHrefs []string, err error) {
	// Parse base URL
	base, err := url.Parse(inputurl)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to parse url. %#v. err: %#v \n", inputurl, err.Error())
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
		c.logger.CrawlLog(fmt.Sprintf("%#v. href : %#v \n", a, href))

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
					c.logger.CrawlLog(fmt.Sprintf("%#v. host != 0, href : %#v \n", a, href))
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
				c.logger.CrawlLog(fmt.Sprintf("%#v. host == 0, href : %#v \n", a, href))
			}

		}
	}

	return filteredHrefs, err
}

func (c *BasicCrawling) StoreD(pagesource string, link string, fr entity.FetchResult) (err error) {

	err = c.Datastore.StoreD(pagesource, link, c.Task, fr.DocumentType, fr.DocumentContentType)
	if err != nil {
		return err
	}

	return err
}

func (c *BasicCrawling) StoreDocument(link string, documentype string, document []byte, documentcontenttype string) (err error) {

	err = c.Datastore.StoreDocument(link, c.Task, documentype, document, documentcontenttype)
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

func (c *BasicCrawling) CrawlpageListParsed(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListParsedData, err error) {

	fmt.Println(param)
	// var dataList []entity.CrawlhrefListData
	dataList, err = c.Datastore.CrawlpageListParsed(param)
	if err != nil {
		return dataList, err
	}

	return dataList, err
}

func (c *BasicCrawling) GetExistingQueue() (queue []string, err error) {

	queue, err = c.Datastore.GetExistingQueue(c.Task)
	if err != nil {
		return queue, err
	}

	return queue, err
}

func (c *BasicCrawling) GetLatestSeedUrl(param_seedurl string) (seedurl string, err error) {

	// var seedurl string
	seedurl, err = c.Datastore.GetLatestSeedUrl(c.Task, param_seedurl)
	if err != nil {
		return param_seedurl, err
	}

	// c.SeedURL = seedurl

	return seedurl, err
}
