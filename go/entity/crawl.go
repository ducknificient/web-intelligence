package entity

type FetchResult struct {
	Header      string
	PdfFile     []byte
	PdfFilename string
	HTMLText    string
}

type CrawlingReq struct {
	Task    string `json:"task"`
	SeedURL string `json:"seedurl"`
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
