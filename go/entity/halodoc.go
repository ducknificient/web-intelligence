package entity

type HalodocListPenyakitReq struct {
}

type TempHalodoc struct {
	Object1 struct {
		Title1 struct {
			B struct {
				Result []struct {
					Name    string `json:"name"`
					Content string `json:"content"`
				} `json:"result"`
			} `json:"b"`
		} `json:"14048149"`
	} `json:"object1"`
}

type HalodocResultData struct {
	// Created_at time.Time `json:"created_at"`
	// Updated_at time.Time `json:"updated_at"`
	AuthorData struct {
		Name        string `json:"name"`
		Entity_type string `json:"entity_type"`
		Slug        string `json:"slug"`
	} `json:"author"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
}

type HalodocListPenyakit struct {
	Result []HalodocResultData `json:"result"`
	Total  int                 `json:"total_count"`
}

type Object1 struct {
}
