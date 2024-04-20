package entity

type AlodokterCrawlerReq struct {
}

// {
// 	"id": "588aaec83971207245025149",
// 	"post_id": 521129,
// 	"post_title": "Ablasi Retina",
// 	"permalink": "ablasi-retina",
// 	"display": "block"
// }

type AlodokterPenyakit struct {
	Id        string `json:"id"`
	PostID    int    `json:"post_id"`
	PostTitle string `json:"post_title"`
	Permalink string `json:"permalink"`
	Display   string `json:"display"`
}

// {
// 	"id": "65f783a78d54d6002655e007",
// 	"post_id": 1894279,
// 	"post_title": "Zoralin",
// 	"permalink": "zoralin",
// 	"display": "block"
// }

type AlodokterObat struct {
	Id        string `json:"id"`
	PostID    int    `json:"postid"`
	PostTitle string `json:"post_title"`
	Permalink string `json:"permalink"`
	Display   string `json:"display"`
}

type AlodokterValidation struct {
	Url     string `json:"url"`
	IsExist bool   `json:"isexist"`
}
