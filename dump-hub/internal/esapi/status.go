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
	"encoding/json"
	"io"
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/r7wx/dump-hub/internal/common"
)

/*
GetStatus - Get status documents (paginated)
*/
func (eClient *Client) GetStatus(from int, size int) (*common.StatusResult, error) {
	query := elastic.NewMatchAllQuery()
	sortQ := elastic.NewFieldSort("status")

	results, err := eClient.client.Search().
		Index("dump-hub-status").
		SortBy(sortQ).
		Query(query).
		From(from).
		Size(size).
		Do(eClient.ctx)
	if err != nil {
		return nil, err
	}

	statusData := common.StatusResult{
		Results: []common.Status{},
		Tot:     0,
	}
	for _, hit := range results.Hits.Hits {
		status := common.Status{}
		err := json.Unmarshal(hit.Source, &status)
		if err != nil {
			return nil, err
		}

		statusData.Results = append(
			statusData.Results,
			status,
		)
	}
	statusData.Tot = int(results.Hits.TotalHits.Value)

	return &statusData, nil
}

/*
GetDocumentStatus - Get the status of a single document (checksum)
*/
func (eClient *Client) GetDocumentStatus(checkSum string) (*common.Status, error) {
	result, err := eClient.client.Get().
		Index("dump-hub-status").
		Id(checkSum).
		Do(eClient.ctx)
	if err != nil {
		return nil, err
	}

	status := common.Status{}
	err = json.Unmarshal(result.Source, &status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

/*
NewStatusDocument - New status document on dump-hub-status index
*/
func (eClient *Client) NewStatusDocument(h *common.Status, checkSum string) error {
	data, err := json.Marshal(h)
	if err != nil {
		return err
	}

	_, err = eClient.client.Index().
		Index("dump-hub-status").
		BodyString(string(data)).
		Id(checkSum).
		Refresh("true").
		Do(eClient.ctx)
	if err != nil {
		return err
	}

	return nil
}

/*
UpdateUploadStatus - Update status field of an upload status document
*/
func (eClient *Client) UpdateUploadStatus(checkSum string, newStatus int) error {
	_, err := eClient.client.Update().
		Index("dump-hub-status").
		Id(checkSum).
		Doc(map[string]interface{}{"status": newStatus}).
		Refresh("true").
		Do(eClient.ctx)
	if err != nil {
		return err
	}

	return nil
}

/*
cleanStatus - Clean unprocessed files and update status (on restart)
*/
func (eClient *Client) cleanStatus() {
	log.Println("Cleaning status of unprocessed files...")

	matchQ := elastic.NewMatchQuery(
		"status",
		common.Processing,
	)
	query := elastic.
		NewBoolQuery().
		Must(matchQ)
	eClient.cleanStatusType(query)

	matchQ = elastic.NewMatchQuery(
		"status",
		common.Deleting,
	)
	query = elastic.
		NewBoolQuery().
		Must(matchQ)
	eClient.cleanStatusType(query)

	matchQ = elastic.NewMatchQuery(
		"status",
		common.Enqueued,
	)
	query = elastic.
		NewBoolQuery().
		Must(matchQ)
	eClient.cleanStatusType(query)
}

func (eClient *Client) cleanStatusType(query *elastic.BoolQuery) error {
	scroll := eClient.client.Scroll().
		Index("dump-hub-status").
		Query(query).
		Size(1)

	for {
		result, err := scroll.Do(eClient.ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		for _, hit := range result.Hits.Hits {
			err = eClient.UpdateUploadStatus(hit.Id, common.Error)
			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}
