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
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/r7wx/dump-hub/internal/filesys"
)

/*
upload - Upload API Handler
*/
func upload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024 * 1024)

		id := r.FormValue("id")
		if len(id) <= 0 {
			log.Printf("(ERROR) (%s) ID value not found", r.URL)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		fileName := r.FormValue("filename")
		if len(fileName) <= 0 {
			log.Printf("(ERROR) (%s) filename value not found", r.URL)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		raw := r.FormValue("offset")
		if len(raw) <= 0 {
			log.Println("(ERROR) offset value not found")
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		offset, err := strconv.Atoi(raw)
		if err != nil {
			log.Printf("(ERROR) %s", err.Error())
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		raw = r.FormValue("file_size")
		if len(raw) <= 0 {
			log.Println("(ERROR) file_size value not found")
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		fileSize, err := strconv.Atoi(raw)
		if err != nil {
			log.Printf("(ERROR) %s", err.Error())
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		if fileSize > filesys.MaxFileSize {
			log.Println("(ERROR) MAX_FILE_SIZE reached")
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		data, metadata, err := r.FormFile("data")
		if err != nil {
			log.Printf("(ERROR) (%s) %s", r.URL, err)
			http.Error(w, "File too large", http.StatusBadRequest)
			return
		}
		defer data.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(data)
		filePath := filepath.Join("/tmp/", id)
		err = writeToFile(
			filePath,
			offset,
			buf.Bytes(),
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		if (offset + int(metadata.Size)) == fileSize {
			err := finishUpload(id, fileName)
			if err != nil {
				log.Printf("(ERROR) (%s) %s", r.URL, err)
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}

		r.MultipartForm.RemoveAll()
		w.WriteHeader(http.StatusOK)
	}
}

/*
writeToFile - Write bytes to file (with offset)
*/
func writeToFile(filePath string, offset int, data []byte) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	whence := io.SeekStart
	_, err = f.Seek(int64(offset), whence)
	if err != nil {
		return err
	}

	f.Write(data)
	f.Sync()
	f.Close()

	return nil
}

/*
finishUpload - Perform post-upload actions
*/
func finishUpload(id string, fileName string) error {
	tmpPath := filepath.Join("/tmp/", id)

	fileName = filesys.SanitizeFilename(fileName)
	filePath := filepath.Join(filesys.UploadFolder, fileName)

	tmpFile, err := os.Open(tmpPath)
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, tmpFile)
	if err != nil {
		return err
	}
	tmpFile.Close()

	err = os.Remove(tmpPath)
	if err != nil {
		return err
	}

	return nil
}
