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

import {
    CdsModal,
    CdsModalHeader,
    CdsModalContent,
    CdsModalActions
} from '@clr/react/modal';
import React from 'react';

class Delete extends React.Component {
    render() {
        if (this.props.file) {
            return (
                <CdsModal onCloseChange={() => this.props.onCancel()}>
                    <CdsModalHeader>
                        <h3 className="modal-title">
                            <cds-icon shape="exclamation-triangle" size="30" solid></cds-icon>&nbsp;Delete file
                        </h3>
                    </CdsModalHeader>
                    <CdsModalContent>
                        <div className="modal-body">
                            <p cds-text="body">Do you want to delete {this.props.file.filename}?</p>
                        </div>
                    </CdsModalContent>
                    <CdsModalActions>
                        <div className="modal-footer">
                            <button type="button"
                                className="btn btn-primary"
                                onClick={() => this.props.onCancel()}>
                                Cancel
                            </button>
                            <button type="button"
                                className="btn btn-danger"
                                onClick={() => this.props.deleteFile(this.props.file)}>
                                Delete
                            </button>
                        </div>
                    </CdsModalActions>
                </CdsModal >
            );
        }

        return null;
    }
}

export default Delete;
