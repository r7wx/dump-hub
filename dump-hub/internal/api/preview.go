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
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/r7wx/dump-hub/internal/filesys"
)

/*
previewReq - API Request Struct
*/
type previewReq struct {
	FileName string `json:"filename"`
	Start    int    `json:"start"`
}

/*
previewResult - API Response Struct
*/
type previewResult struct {
	Preview []string `json:"preview"`
}

/*
preview - Preview API Handler
*/
func preview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var previewReq previewReq

		err := json.NewDecoder(r.Body).Decode(&previewReq)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		fileName := filesys.SanitizeFilename(previewReq.FileName)
		filePath := filepath.Join(
			filesys.UploadFolder,
			fileName,
		)
		previewData, err := readPreview(
			filePath,
			previewReq.Start,
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		preview := previewResult{
			Preview: *previewData,
		}
		response, err := json.Marshal(preview)
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

/*
readPreview - Get preview data from a file
(from start point to next "previewSize" lines).
*/
func readPreview(filePath string, start int) (*[]string, error) {
	previewData := []string{}

	if start < 0 {
		start = 0
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	currentLine := 0
	for {
		if currentLine >= (start + filesys.PreviewSize) {
			break
		}

		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Println(err)
				return nil, err
			}
		}
		if isPrefix {
			return nil, errors.New("file line too long")
		}

		if currentLine >= start {
			previewData = append(previewData, string(line))
		}
		currentLine++
	}

	return &previewData, nil
}
