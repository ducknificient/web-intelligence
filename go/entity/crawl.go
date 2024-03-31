package entity

type FetchResult struct {
	HTMLText            string
	DocumentFile        []byte
	DocumentContentType string // MIME type
	DocumentType        string
}

type CrawlingReq struct {
	Task    string `json:"task"`
	SeedURL string `json:"seedurl"`
}

type CrawlingMultipleReq struct {
	Task        string   `json:"task"`
	SeedURLList []string `json:"seedurllist"`
}

type CrawlingResponse struct {
}

type CrawlpageListReq struct {
	Page   int    `json:"page"`
	Count  int    `json:"count"`
	Search string `json:"search"`
}

type CrawlpageListParam struct {
	Page   int    `json:"page"`
	Count  int    `json:"count"`
	Search string `json:"search"`
}

type CrawlhrefListData struct {
	Link string `json:"link"`
	Href string `json:"href"`
}

type CrawlpageListData struct {
	Pagesource string              `json:"pagesource"`
	Link       string              `json:"link"`
	HrefList   []CrawlhrefListData `json:"hreflist"`
}

type CrawlpageListDataWrap struct {
	Total int                 `json:"total"`
	Data  []CrawlpageListData `json:"data"`
}

type CrawlpageListDataRes struct {
	Sc   int                   `json:"sc"`
	St   bool                  `json:"st"`
	Msg  string                `json:"msg"`
	Data CrawlpageListDataWrap `json:"data"`
}

type CrawlpageListParsedData struct {
	Pagesource  string `json:"pagesource"`
	Link        string `json:"link"`
	Task        string `json:"task"`
	Metatitle   string `json:"metatitle"`
	Metacontent string `json:"metacontent"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Category    string `json:"category"`
	TotalView   string `json:"totalview"`
	Hashtag     string `json:"hashtag"`
	Content     string `json:"content"`
	RelatedNews string `json:"relatednews"`
}

type CrawlpageListParsedDataWrap struct {
	Total int                       `json:"total"`
	Data  []CrawlpageListParsedData `json:"data"`
}

type CrawlpageListParsedDataRes struct {
	Sc   int                         `json:"sc"`
	St   bool                        `json:"st"`
	Msg  string                      `json:"msg"`
	Data CrawlpageListParsedDataWrap `json:"data"`
}
