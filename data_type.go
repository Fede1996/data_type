package data_type

type SearchList struct {
	UserID                   string   `json:"user_id"`
	MaxSearchList            int64    `json:"max_search_list"`
	MinSimilarityCoefficient float64  `json:"min_similarity_coefficient"`
	SearchSite               []string `json:"search_site"`
	SearchList               []Search `json:"search_list"`
}

type SearchResults struct {
	SearchList []SearchResult `json:"search_list"`
}

type SearchResult struct {
	Id      int      `json:"id"`
	Results []Result `json:"results"`
}

type ReportTable struct {
	Header []string
	Data   [][]string
}

type Search struct {
	Id     int     `json:"id"`
	Params []Param `json:"params"`
	Date   string  `json:"date"`
}

type Param struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

type Result struct {
	Source       string  `json:"source"`
	Url          string  `json:"url"`
	Date         string  `json:"date"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Duration     string  `json:"duration"`
	Weight       float64 `json:"weight"`
	Channel      Channel `json:"channel"`
	ThumbnailUrl string  `json:"thumbnail_url"`
}

type Channel struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type SearchQuery struct {
	Name   string
	Weight float64
}

type SplittedParams struct {
	Required []string
	Optional []Param
}

type sqlParam struct {
	username string
	password string
	ip       string
	port     string
	schema   string
}
