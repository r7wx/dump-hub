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
	"bufio"
	"log"
	"os"

	"github.com/olivere/elastic/v7"

	"github.com/r7wx/dump-hub/internal/common"
	"github.com/r7wx/dump-hub/internal/parser"
)

/*
IndexFile - Process file for indexing
*/
func (eClient *Client) IndexFile(parser *parser.Parser) {
	for {
		if !eClient.isBusy() {
			break
		}
	}
	eClient.setBusy(true)

	log.Printf("Processing %s...", parser.Checksum)
	eClient.UpdateUploadStatus(
		parser.Checksum,
		common.Processing,
	)

	file, err := os.Open(parser.Filepath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	entries := []*common.Entry{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(entries) >= eClient.bulkw.cSize {
			err := eClient.BulkInsert(
				entries,
				parser.Checksum,
			)
			if err != nil {
				eClient.UpdateUploadStatus(
					parser.Checksum,
					common.Error,
				)
				log.Println(err)
				return
			}
			entries = []*common.Entry{}
		}

		entry := parser.ParseEntry(scanner.Text())
		if entry == nil {
			continue
		}

		entries = append(entries, entry)
	}

	if len(entries) > 0 {
		err := eClient.BulkInsert(
			entries,
			parser.Checksum,
		)
		if err != nil {
			eClient.UpdateUploadStatus(
				parser.Checksum,
				common.Error,
			)
			log.Println(err)
			return
		}
	}

	indexName := "dump-hub-doc-" + parser.Checksum
	eClient.Refresh(indexName)

	log.Printf(
		"Processing complete: %s",
		parser.Checksum,
	)
	eClient.UpdateUploadStatus(
		parser.Checksum,
		common.Complete,
	)

	eClient.setBusy(false)
}

/*
BulkInsert - Index entries with BulkAPI on correct doc index
*/
func (eClient *Client) BulkInsert(e []*common.Entry, checksum string) error {
	indexName := "dump-hub-doc-" + checksum

	err := eClient.CreateIndex(
		indexName,
		entryMapping,
	)
	if err != nil {
		return err
	}

	bulkRequest := eClient.client.Bulk()
	for _, entry := range e {
		req := elastic.NewBulkIndexRequest().
			OpType("index").
			Index(indexName).
			Doc(entry)

		bulkRequest = bulkRequest.Add(req)
	}

	_, err = bulkRequest.
		Do(eClient.ctx)
	if err != nil {
		return err
	}

	return nil
}
