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

import StatusIcon from './status-icon';
import Delete from './delete.js';
import React from 'react';
import './data-view.css';

class DataView extends React.Component {
    render() {
        const data = this.props.data;
        const items = [];

        if (data.length > 0) {
            for (let i = 0; i < data.length; i++) {
                items.push(
                    <tr>
                        <td className="table-content">
                            <b>{data[i].filename}</b> &nbsp;
                            <code>({data[i].checksum})</code>
                        </td>
                        <td className="table-content">
                            {data[i].date}
                        </td>
                        <StatusIcon status={data[i].status} />
                        <td className="delete-button">
                            <button onClick={() => this.props.onDeleteClick(data[i])}
                                className={`${(data[i].status !== 4 && data[i].status !== 3) ? "disabled" : ""}`}>
                                <clr-icon shape="trash" size="20"></clr-icon>
                            </button>
                        </td>
                    </tr>
                );
            }
        } else {
            items.push(
                <tr>
                    <td colspan="5" class="no-items">
                        <br />
                        <clr-icon shape="block" size="70"></clr-icon>
                        <br /><br />
                        <code>NO DATA</code>
                        <br /><br />
                    </td>
                </tr>
            );
        }

        return (
            <React.Fragment>
                <Delete toDelete={this.props.toDelete}
                    onDeleteClick={this.props.onDeleteClick}
                    onDelete={this.props.onDelete}
                    onCancel={this.props.onDeleteCancel} />
                <table className="table">
                    <tbody>
                        <tr>
                            <td className="folder" colspan="4">
                                <clr-icon shape="block" size="25"></clr-icon>
                                &nbsp;<b>Data</b>
                            </td>
                        </tr>
                        {items}
                    </tbody>
                </table>
            </React.Fragment>
        );
    }
}

export default DataView;