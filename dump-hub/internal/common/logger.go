package common

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
	"io"
	"log"
	"os"
	"path/filepath"
)

/*
InitLogger - Set log file redirection to file
*/
func InitLogger() {
	filePath := "/var/log/dump-hub/"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.Mkdir(filePath, 0644)
		if err != nil {
			log.Println(err)
		}
	}

	filename := filepath.Join(filePath, "dump-hub.log")
	logFile, err := os.OpenFile(
		filename,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		logFile.Close()
		log.Printf("[ERROR] error creating logfile %s", err.Error())
		log.Println("Logging to stdout only")
		log.SetOutput(os.Stdout)
		return
	}

	logWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(logWriter)
}
