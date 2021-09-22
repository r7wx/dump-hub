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

import "log"

/*
CreateIndex - Create elasticsearch index if not exists
*/
func (eClient *Client) CreateIndex(index string, mapping string) error {
	exists, err := eClient.client.IndexExists(index).Do(eClient.ctx)
	if err != nil {
		return err
	}

	if !exists {
		_, err := eClient.client.
			CreateIndex(index).
			Body(mapping).
			Do(eClient.ctx)
		if err != nil {
			return err
		}
		log.Printf("Created elasticsearch index: %s", index)
	}

	return nil
}

/*
Refresh - Refresh an elasticsearch index
*/
func (eClient *Client) Refresh(index string) error {
	_, err := eClient.client.Refresh().
		Index(index).
		Do(eClient.ctx)
	if err != nil {
		return err
	}
	return nil
}
