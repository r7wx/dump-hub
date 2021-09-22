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

import './preview-table.css';
import React from 'react';

class PreviewTable extends React.Component {
    render() {
        const prevTable = this.props.previewTable;
        if (prevTable.length > 0) {
            const tableHeaders = [];
            const tableRows = [];

            for (let i = 0; i < this.props.maxCols; i++) {
                if (this.isColumnSelected(i)) {
                    tableHeaders.push(
                        <th>
                            <button class="column-button" onClick={() => this.props.toggleColumn(i)}>
                                <clr-icon shape="minus-circle" size="20" ></clr-icon>
                            </button>
                        </th >
                    );
                } else {
                    tableHeaders.push(
                        <th>
                            <button class="column-button" onClick={() => this.props.toggleColumn(i)}>
                                <clr-icon shape="plus-circle" size="20" ></clr-icon>
                            </button>
                        </th >
                    );
                }
            }
            for (let i = 0; i < prevTable.length; i++) {
                const tableRow = [];
                for (let j = 0; j < prevTable[i].length; j++) {
                    tableRow.push(
                        <td className={`${this.isColumnSelected(j) ? "selected-column" : ""}`}>
                            {prevTable[i][j]}
                        </td >
                    );
                }
                tableRows.push(
                    <tr>
                        {tableRow}
                    </tr>
                );
            }

            return (
                <div className="table-container">
                    <table className="table">
                        <thead>
                            <tr>
                                {tableHeaders}
                            </tr>
                        </thead>
                        <tbody>
                            {tableRows}
                        </tbody>
                    </table>
                </div>
            );
        }

        return null;
    }

    isColumnSelected(index) {
        const selected = this.props.columns;
        return (selected.indexOf(index) >= 0);
    }
}

export default PreviewTable;