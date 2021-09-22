package parser

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
	"fmt"
	"strings"

	"github.com/r7wx/dump-hub/internal/common"
	"github.com/r7wx/dump-hub/internal/filesys"
)

/*
Parser - Parser struct
*/
type Parser struct {
	Start     int
	Separator string
	Columns   []int
	Filename  string
	Filepath  string
	Checksum  string
}

/*
New - Create new parser object
*/
func New(pattern string, columns []int, filename string, filepath string) (*Parser, error) {
	p := &Parser{}

	_, err := fmt.Sscanf(
		pattern,
		"{%d}{%1s}",
		&p.Start,
		&p.Separator,
	)
	if err != nil {
		return nil, err
	}

	p.Columns = columns
	p.Filename = filename
	p.Filepath = filepath

	p.Checksum, err = filesys.ComputeChecksum(p.Filepath)
	if err != nil {
		return nil, err
	}

	return p, nil
}

/*
ParseEntry - Parse dump entry from file
*/
func (p *Parser) ParseEntry(entry string) *common.Entry {
	obj := &common.Entry{}
	data := []string{}

	if len(entry) < 1 {
		return nil
	}

	line := strings.Replace(entry, " ", "", -1)
	matches := strings.Split(line, p.Separator)
	if len(matches) < 1 {
		return nil
	}

	for i, match := range matches {
		if len(match) < 1 {
			continue
		}

		for _, column := range p.Columns {
			if i == column {
				data = append(data, match)
			}
		}
	}

	if obj == nil {
		return nil
	}

	obj.Origin = p.Filename
	obj.OriginID = p.Checksum
	obj.Data = data

	return obj
}
