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
	"sync"
)

/*
BulkWorker - Elasticsearch BulkAPI Worker
*/
type BulkWorker struct {
	mutex sync.Mutex
	busy  bool
	cSize int
}

/*
newWorker - Create BulkWorker
*/
func (eClient *Client) newWorker(chunkSize int) *BulkWorker {
	bulkW := BulkWorker{
		cSize: chunkSize,
	}
	return &bulkW
}

/*
isBusy - Get BulkWorker busy value
*/
func (eClient *Client) isBusy() bool {
	eClient.bulkw.mutex.Lock()
	defer eClient.bulkw.mutex.Unlock()
	return eClient.bulkw.busy
}

/*
setBusy - Set BulkWorker busy value
*/
func (eClient *Client) setBusy(value bool) {
	eClient.bulkw.mutex.Lock()
	defer eClient.bulkw.mutex.Unlock()
	eClient.bulkw.busy = value
}
