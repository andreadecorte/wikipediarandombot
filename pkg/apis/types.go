package apis

import (
	"fmt"
	"time"
)

type WikiPage struct {
	Pageid               float64   `json:"pageid"`
	Ns                   float64   `json:"ns"`
	Title                string    `json:"title"`
	Contentmodel         string    `json:"contentmodel"`
	Pagelanguage         string    `json:"pagelanguage"`
	Pagelanguagehtmlcode string    `json:"pagelanguagehtmlcode"`
	Pagelanguagedir      string    `json:"pagelanguagedir"`
	Touched              time.Time `json:"touched"`
	Lastrevid            int       `json:"lastrevid"`
	Length               float64   `json:"length"`
	Fullurl              string    `json:"fullurl"`
}

func (p WikiPage) String() string {
	return fmt.Sprintf("%9.f:%s (%g)", p.Pageid, p.Title, p.Length)
}

type Pages struct {
	Items []WikiPage
}
