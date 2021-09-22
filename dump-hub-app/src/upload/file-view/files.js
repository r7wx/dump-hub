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

import Delete from './delete.js';
import React from 'react';
import './file-view.css';

class Files extends React.Component {
    constructor(props) {
        super(props);
        this.state = { toDelete: null };
        this.onDelete = this.onDelete.bind(this);
        this.onDeleteCancel = this.onDeleteCancel.bind(this);
    }

    render() {
        const files = this.props.files;
        if (files.length > 0) {
            const entries = [];
            for (let i = 0; i < files.length; i++) {
                entries.push(
                    <tr>
                        <td className="file">
                            <clr-icon shape="file" size="25"></clr-icon>&nbsp;
                            <b>{files[i].filename}</b>&nbsp;
                            <b> {files[i].size} MB</b>
                        </td>
                        <td className="delete-button">
                            <button onClick={() => this.onDeleteClick(files[i])}>
                                <clr-icon shape="trash" size="20"></clr-icon>
                            </button>
                        </td >
                    </tr>
                );
            }

            return (
                <React.Fragment>
                    <Delete file={this.state.toDelete}
                        onCancel={this.onDeleteCancel}
                        deleteFile={this.onDelete} />
                    {entries}
                </React.Fragment>
            );
        }

        return (
            <tr>
                <td colspan="5" className="no-items">
                    <br />
                    <clr-icon shape="folder-open" size="70"></clr-icon>
                    <br /><br />
                    <code>EMPTY FOLDER</code>
                    <br /><br />
                </td>
            </tr>
        );
    }

    onDelete(file) {
        this.props.deleteFile(file);
        this.setState({ toDelete: null });
    }

    onDeleteClick(file) {
        this.setState({ toDelete: file });
    }

    onDeleteCancel() {
        this.setState({ toDelete: null });
    }
}

export default Files;