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
	"log"

	"github.com/r7wx/dump-hub/internal/common"
)

/*
DeleteEntries - Delete entries associated to a file (checkSum)
(Delete the index with the name dump-hub-doc-<checksum>)
*/
func (eClient *Client) DeleteEntries(checkSum string) {
	for {
		if !eClient.isBusy() {
			break
		}
	}
	eClient.setBusy(true)

	eClient.UpdateUploadStatus(
		checkSum,
		common.Deleting,
	)

	indexName := "dump-hub-doc-" + checkSum
	_, err := eClient.client.DeleteIndex(
		indexName,
	).Do(eClient.ctx)
	if err != nil {
		eClient.UpdateUploadStatus(
			checkSum,
			common.Error,
		)
		log.Println(err)
		return
	}

	eClient.Refresh(indexName)
	_, err = eClient.client.Delete().
		Index("dump-hub-status").
		Id(checkSum).
		Do(eClient.ctx)
	if err != nil {
		log.Println(err)
	}

	eClient.setBusy(false)
}
