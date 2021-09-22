package api

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
	"net/http"

	"github.com/r7wx/dump-hub/internal/esapi"
)

/*
searchReq - API Request Struct
*/
type searchReq struct {
	Query string `json:"query"`
	Page  int    `json:"page"`
}

/*
search - Search API Handler
*/
func search(eClient *esapi.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var searchReq searchReq

		err := json.NewDecoder(r.Body).Decode(&searchReq)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		query := "*"
		if len(searchReq.Query) > 0 {
			query = searchReq.Query
		}

		from := pageSize * (searchReq.Page - 1)
		results, err := eClient.Search(
			string(query),
			from,
			pageSize,
		)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		response, err := json.Marshal(results)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
