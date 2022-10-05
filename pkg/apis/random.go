package apis

import (
	"errors"
	"fmt"
	"time"

	"github.com/Jeffail/gabs/v2"
)

func getField(v *gabs.Container, key string) (interface{}, error) {
	val := v.ChildrenMap()[key].Data()
	if val == nil {
		return val, errors.New("Missing field " + key)
	}
	return val, nil
}

func parseJson(jsonInput []byte, out *Pages) error {
	jsonParsed, jsonErr := gabs.ParseJSON(jsonInput)
	if jsonErr != nil {
		return jsonErr
	}
	for _, v := range jsonParsed.Path("query.pages").ChildrenMap() {
		page := WikiPage{}
		val, err := getField(v, "title")
		if err != nil {
			return err
		}
		page.Title = val.(string)

		val2, err2 := getField(v, "touched")
		if err2 != nil {
			return err2
		}
		parsedTime, err2b := time.Parse(time.RFC3339, val2.(string))
		if err2b != nil {
			return err2b
		} else {
			page.Touched = parsedTime
		}

		val3, err3 := getField(v, "length")
		if err3 != nil {
			return err3
		}
		page.Length = val3.(float64)

		val4, err4 := getField(v, "fullurl")
		if err4 != nil {
			return err4
		}
		page.Fullurl = val4.(string)

		val5, err5 := getField(v, "pagelanguage")
		if err5 != nil {
			return err5
		}
		page.Pagelanguage = val5.(string)

		val6, err6 := getField(v, "ns")
		if err6 != nil {
			return err6
		}
		page.Ns = val6.(float64)

		val7, err7 := getField(v, "pageid")
		if err7 != nil {
			return err7
		}
		page.Pageid = val7.(float64)
		//for k2,v2 := range v.ChildrenMap() {
		//	log.Warn(k2, "->", v2.Data())
		//}

		out.Items = append(out.Items, page)
	}
	return nil
}

func GetRandomWikiPages(out *Pages, lang string, limit int) error {
	uri := fmt.Sprintf("https://%s.wikipedia.org/w/api.php?action=query&format=json&prop=info&inprop=url&generator=random&redirects=1&grnnamespace=0&grnlimit=%d", lang, limit)
	body, _, _ := getConnection(uri)
	err := parseJson(body, out)
	if err != nil {
		return err
	}

	return nil
}
