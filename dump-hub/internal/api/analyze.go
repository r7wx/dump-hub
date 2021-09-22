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
	"os"
	"path/filepath"
	"time"

	"github.com/r7wx/dump-hub/internal/common"
	"github.com/r7wx/dump-hub/internal/esapi"
	"github.com/r7wx/dump-hub/internal/filesys"
	"github.com/r7wx/dump-hub/internal/parser"
)

/*
analyzeRequest - API Request Struct
*/
type analyzeRequest struct {
	Filename string `json:"filename"`
	Pattern  string `json:"pattern"`
	Columns  []int  `json:"columns"`
}

/*
analyze - Analyze API Handler
*/
func analyze(eClient *esapi.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var analyzeReq analyzeRequest

		err := json.NewDecoder(r.Body).Decode(&analyzeReq)
		if err != nil {
			log.Println(err)
			http.Error(w, "malformed request", http.StatusBadRequest)
			return
		}

		if len(analyzeReq.Columns) < 1 {
			http.Error(w, "invalid columns value", http.StatusBadRequest)
			return
		}

		fileName := filesys.SanitizeFilename(analyzeReq.Filename)
		originPath := filepath.Join(filesys.UploadFolder, fileName)
		if _, err := os.Stat(originPath); os.IsNotExist(err) {
			log.Println(err)
			http.Error(w, "file does not exist", http.StatusNotFound)
			return
		}

		checkSum, err := filesys.ComputeChecksum(originPath)
		if err != nil {
			log.Println(err)
			http.Error(w, "unable to compute file checksum", http.StatusInternalServerError)
			return
		}

		alreadyUploaded, err := eClient.IsAlreadyUploaded(checkSum)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		if alreadyUploaded {
			log.Printf("%s already uploaded!", checkSum)
			http.Error(w, "file already uploaded", http.StatusBadRequest)
			return
		}

		go analyzeFile(
			eClient,
			analyzeReq,
			fileName,
			checkSum,
		)

		w.WriteHeader(http.StatusOK)
	}
}

/*
analyzeFile - Parse and index the selected file
*/
func analyzeFile(eClient *esapi.Client, req analyzeRequest, fn string, cs string) {
	filePath, err := filesys.MoveTemp(fn)
	if err != nil {
		log.Println(err)
		return
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	status := common.Status{
		Date:     date,
		Filename: fn,
		Checksum: cs,
		Status:   common.Enqueued,
	}
	err = eClient.NewStatusDocument(&status, cs)
	if err != nil {
		log.Println(err)
		return
	}

	parser, err := parser.New(
		req.Pattern,
		req.Columns,
		req.Filename,
		filePath,
	)
	if err != nil {
		log.Println("Unable to create parser object")
		return
	}

	eClient.IndexFile(parser)
}
