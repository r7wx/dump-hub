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

	"github.com/r7wx/dump-hub/internal/common"
	"github.com/r7wx/dump-hub/internal/esapi"
)

/*
deleteRequest - API Request Struct
*/
type deleteRequest struct {
	Checksum string `json:"checksum"`
}

/*
delete - Delete API Handler
*/
func delete(eClient *esapi.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deleteReq deleteRequest

		err := json.NewDecoder(r.Body).Decode(&deleteReq)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		checkSum := deleteReq.Checksum

		currentStatus, err := eClient.GetDocumentStatus(checkSum)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		isComplete := currentStatus.Status == common.Complete
		isError := currentStatus.Status == common.Error
		if !isComplete && !isError {
			http.Error(w, "", http.StatusTeapot)
			log.Println(err)
			return
		}

		eClient.UpdateUploadStatus(checkSum, common.Enqueued)
		go eClient.DeleteEntries(checkSum)

		w.WriteHeader(http.StatusOK)
	}
}
