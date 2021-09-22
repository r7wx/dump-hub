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

import React from 'react';
import './file-view.css';

class Files extends React.Component {
    render() {
        const entries = [];
        const files = this.props.files;
        if (files.length > 0) {
            for (let i = 0; i < files.length; i++) {
                entries.push(
                    <tr>
                        <td className={`entry ${this.props.isSelected(files[i]) ? "selected" : ""}`}
                            onClick={() => this.props.selectFile(files[i])} >
                            <clr-icon shape="file" size="25"></clr-icon>&nbsp;
                            <b>{files[i].filename}</b>&nbsp;
                            <b>{files[i].size} MB</b>
                        </td>
                    </tr >
                );
            }
        } else {
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

        return (
            <React.Fragment>
                {entries}
            </React.Fragment>
        );
    }
}

export default Files;