package esapi

/*
The MIT License (MIT)
Copyright (c) 2021 r7wx
Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation
files (the "Software"), to deal in the Software without
restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following
conditions:
The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"encoding/json"
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/r7wx/dump-hub/internal/common"
)

/*
Search - Search entries on dump-hub index (Paginated)
*/
func (eClient *Client) Search(queryString string, from int, size int) (*common.SearchResult, error) {
	var results *elastic.SearchResult
	var err error

	if len(queryString) < 1 || queryString == "*" {
		query := elastic.NewMatchAllQuery()
		results, err = eClient.client.Search().
			Index("dump-hub-doc-*").
			Query(query).
			From(from).
			Size(size).
			Do(eClient.ctx)
		if err != nil {
			return nil, err
		}
	} else {
		query := elastic.
			NewMultiMatchQuery(
				queryString,
				"_all",
			).
			Type("match_phrase")
		results, err = eClient.client.Search().
			Index("dump-hub-doc-*").
			Query(query).
			From(from).
			Size(size).
			Do(eClient.ctx)
		if err != nil {
			return nil, err
		}
	}

	searchResult := common.SearchResult{
		Results: []common.Entry{},
		Tot:     0,
	}
	for _, hit := range results.Hits.Hits {
		entry := common.Entry{}
		err := json.Unmarshal(hit.Source, &entry)
		if err != nil {
			log.Println(err)
			break
		}

		searchResult.Results = append(
			searchResult.Results,
			entry,
		)
	}
	searchResult.Tot = int(results.Hits.TotalHits.Value)

	return &searchResult, nil
}
