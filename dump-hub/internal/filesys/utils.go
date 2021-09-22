package filesys

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
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/r7wx/dump-hub/internal/common"
)

/*
CleanTMP - Clean tmp folder
*/
func CleanTMP() {
	log.Println("Cleaning tmp folder...")

	dir, err := ioutil.ReadDir("/tmp")
	if err != nil {
		log.Println(err)
	}
	for _, d := range dir {
		os.Remove(path.Join("/tmp", d.Name()))
	}
}

/*
ComputeChecksum - Compute file checksum
*/
func ComputeChecksum(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

/*
SanitizeFilename - Sanitize filename string
*/
func SanitizeFilename(fileName string) string {
	badChars := []string{
		"~", "./", "/", "../", "\"", "&", "$", "#",
		"%20", "%22", "%3c", "%253c", "%3e", "%28", "%29",
		"%2528", "%26", "%24", "%3f", "%3b", "%3d",
	}

	fileName = filepath.Base(fileName)
	fileName = strings.TrimSpace(fileName)
	for _, badChar := range badChars {
		fileName = strings.Replace(fileName, badChar, "", -1)
	}
	fileName = strings.Replace(fileName, "\\", "", -1)

	return fileName
}

/*
MoveTemp - Move uploaded file to tmp folder before
analyze
*/
func MoveTemp(fileName string) (string, error) {
	originPath := filepath.Join(UploadFolder, fileName)
	hiddenPath := filepath.Join(UploadFolder, "."+fileName)
	filePath := filepath.Join("/tmp/", fileName)

	err := os.Rename(
		originPath,
		hiddenPath,
	)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	originFile, err := os.Open(hiddenPath)
	if err != nil {
		return "", err
	}
	defer originFile.Close()

	_, err = io.Copy(file, originFile)
	if err != nil {
		return "", err
	}
	originFile.Close()

	err = os.Remove(hiddenPath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

/*
ReadUploadFolder - Returns the content of "uploads" folder
*/
func ReadUploadFolder() ([]common.File, error) {
	var files = []common.File{}

	fileInfo, err := ioutil.ReadDir(UploadFolder)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		if file.Name()[0] == '.' {
			continue
		}

		fileSize := file.Size() / 1000000
		uFile := common.File{
			FileName: file.Name(),
			Size:     fileSize,
		}
		files = append(files, uFile)
	}

	return files, nil
}
