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
	TimeToRead           int
}

func (p WikiPage) String() string {
	return fmt.Sprintf("%9.f:%s (%d - %d minutes read)", p.Pageid, p.Title, int(p.Length), int(p.TimeToRead))
}

type Pages struct {
	Items []WikiPage
}
