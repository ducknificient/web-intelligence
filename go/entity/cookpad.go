package entity

type CookpadCrawlerReq struct {
	Menu string `json:"menu"`
}

type CookpadRecipeList struct {
	Url string
}

type CookpadValidation struct {
	Url     string `json:"url"`
	IsExist bool   `json:"isexist"`
}

type CookpadImageList struct {
	Filename string
	Image    []byte
}

type CookpadSaveImageParam struct {
	List []CookpadImageList
}
