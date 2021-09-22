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
import './queue.css';

class Queue extends React.Component {
    render() {
        const fileItems = [];
        const fileQueue = this.props.fileQueue;
        for (let i = 0; i < fileQueue.length; i++) {
            const fileItem = [];
            fileItem.push(
                <b>{fileQueue[i].file.name}</b>
            );

            if (fileQueue[i].complete) {
                fileItem.push(
                    <div className="progress success labeled">
                        <progress value={fileQueue[i].progress} max="100" ></progress>
                        <span>{fileQueue[i].progress}%</span>
                    </div>
                );
            }
            if (!fileQueue[i].error && !fileQueue[i].complete) {
                if (fileQueue[i].progress > 0) {
                    fileItem.push(
                        <div className="progress labeled">
                            <progress value={fileQueue[i].progress} max="100" ></progress>
                            <span>{fileQueue[i].progress}%</span>
                        </div>
                    );
                } else {
                    fileItem.push(
                        <div className="progress loop">
                            <progress></progress>
                        </div >
                    );
                }
            }
            if (fileQueue[i].error) {
                fileItem.push(
                    <React.Fragment>
                        <div className="progress labeled">
                            <progress value="0" max="100"></progress>
                            <span className="error-text">NaN</span>
                        </div>
                        <b className="error-text">
                            <clr-icon shape="error-standard"></clr-icon>
                            &nbsp;{fileQueue[i].error}
                        </b>
                    </React.Fragment>
                );
            }
            if (fileQueue[i].complete) {
                fileItem.push(
                    <b className="success-text">
                        <clr-icon shape="success-standard"></clr-icon>
                        &nbsp;Upload complete
                    </b>
                );
            }

            fileItem.push(<br />);
            fileItems.push(
                <React.Fragment>
                    {fileItem}
                </React.Fragment>
            );
        }

        return (
            <div className="selected-files">
                {fileItems}
            </div>
        );
    }
}

export default Queue;
