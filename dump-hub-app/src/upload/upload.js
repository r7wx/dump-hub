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

import FileInput from './file-input/file-input.js';
import FileView from './file-view/file-view.js';
import Queue from './queue/queue.js';
import * as uuid from 'uuid';
import React from 'react';
import axios from 'axios';

class Upload extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            fileQueue: [],
            files: [],
            loading: true,
        };
        this.getFiles = this.getFiles.bind(this);
        this.deleteFile = this.deleteFile.bind(this);
        this.onFileSelect = this.onFileSelect.bind(this);
    }

    render() {
        return (
            <React.Fragment>
                <FileView files={this.state.files}
                    getFiles={this.getFiles}
                    deleteFile={this.deleteFile}
                    loading={this.state.loading} />
                <FileInput onSelect={this.onFileSelect} />
                <Queue fileQueue={this.state.fileQueue} />
            </React.Fragment>
        );
    }

    componentDidMount() {
        this.getFiles();
    }

    getFiles() {
        const base_api = process.env.REACT_APP_BASE_API;
        const uri = base_api + 'files';
        this.setState({ loading: true });
        axios.get(uri)
            .then(response => this.handleFileResponse(response['data']))
            .catch(() => { this.setState({ loading: true }); });
    }

    handleFileResponse(response) {
        this.setState({
            loading: false,
            files: response['files']
        });
    }

    deleteFile(file) {
        const base_api = process.env.REACT_APP_BASE_API;
        const file_id = btoa(file.filename);
        const uri = base_api + 'files/' + file_id;

        this.setState({ loading: true });
        axios.delete(uri)
            .then(() => {
                this.setState({ loading: false });
                this.getFiles();
            })
            .catch(() => { });
    }

    onFileSelect(event) {
        let selectedFiles = this.state.fileQueue;
        const files = event.target.files;
        for (let i = 0; i < files.length; i++) {
            const selectedFile = {
                file: files[i],
                uuid: uuid.v4(),
                pending: true,
                complete: false,
                error: null,
                chunkSent: 0,
                progress: 0
            };

            let alreadySelected = false;
            for (let j = 0; j < selectedFiles.length; j++) {
                if (selectedFiles[j].file.name === selectedFile.file.name) {
                    alreadySelected = true;
                }
            }

            if (!alreadySelected) {
                selectedFiles.push(selectedFile);
            }
        }

        this.setState({
            fileQueue: selectedFiles
        }, () => this.uploadQueue());
    }

    uploadQueue() {
        const files = this.state.fileQueue;
        for (let i = 0; i < files.length; i++) {
            if (files[i].pending) {
                files[i].pending = false;
                this.setState({
                    fileQueue: files
                }, () => this.uploadFile(files, i));
            }
        }
    }

    uploadFile(files, i) {
        const base_api = process.env.REACT_APP_BASE_API;
        const uri = base_api + 'upload';
        const chunkSize = 30 * 1000000;
        const fileSize = files[i].file.size;
        const chunks = Math.ceil(fileSize / chunkSize);

        let chunk = 0;
        while (chunk < chunks) {
            if (files[i].error != null) {
                break;
            }

            const offset = chunk * chunkSize;
            const slice = files[i].file.slice(offset, offset + chunkSize);

            const formData = new FormData();
            formData.append('id', files[i].uuid);
            formData.append('filename', files[i].file.name);
            formData.append('offset', offset.toString());
            formData.append('file_size', fileSize.toString());
            formData.append('data', slice);

            const config = {
                headers: {
                    'content-type': 'multipart/form-data'
                }
            };
            axios.post(uri, formData, config)
                .then(() => {
                    files[i].chunkSent++;
                    files[i].progress = Math.ceil(
                        (files[i].chunkSent * 100) / chunks
                    );
                    if (files[i].chunkSent === chunks) {
                        files[i].complete = true;
                        this.getFiles();
                    }
                    this.setState({
                        fileQueue: files
                    });
                })
                .catch(() => {
                    files[i].error = "Unable to upload selected file";
                    this.setState({
                        fileQueue: files
                    });
                });

            chunk++;
        }
    }
}

export default Upload;