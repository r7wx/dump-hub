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

import PreviewTable from './preview-table/preview-table.js';
import PreviewBox from './preview-box/preview-box.js';
import Loading from '../core/loading/loading.js';
import FileView from './file-view/file-view.js';
import Pattern from './pattern/pattern.js';
import Alerts from './alerts/alerts.js';
import Submit from './submit/submit.js';
import axios from 'axios';
import React from 'react';

class Analyze extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            loading: true,
            alerts: [],
            files: [],
            preview: [],
            previewTable: [],
            maxCols: 0,
            selected: undefined,
            start: 0,
            separator: ':',
            columns: [],
            pending: false
        };
        this.selectFile = this.selectFile.bind(this);
        this.onStartChange = this.onStartChange.bind(this);
        this.onSeparatorChange = this.onSeparatorChange.bind(this);
        this.toggleColumn = this.toggleColumn.bind(this);
        this.removeAlert = this.removeAlert.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
    }

    render() {
        return (
            <React.Fragment>
                <Alerts alerts={this.state.alerts}
                    removeAlert={this.removeAlert} />
                <Loading isLoading={this.state.loading} />
                <FileView files={this.state.files}
                    selected={this.state.selected}
                    selectFile={this.selectFile} />
                <br />
                <Pattern start={this.state.start}
                    separator={this.state.separator}
                    onStartChange={this.onStartChange}
                    onSeparatorChange={this.onSeparatorChange} />
                <br />
                <PreviewBox preview={this.state.preview} />
                <PreviewTable maxCols={this.state.maxCols}
                    previewTable={this.state.previewTable}
                    columns={this.state.columns}
                    toggleColumn={this.toggleColumn} />
                <Submit columns={this.state.columns}
                    separator={this.state.separator}
                    pending={this.state.pending}
                    onSubmit={this.onSubmit} />
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
        this.selectFile(
            this.state.files[0]
        );
    }

    getPreview(file) {
        if (!file) { return; }
        const start = this.state.start;
        const base_api = process.env.REACT_APP_BASE_API;
        const uri = base_api + 'preview';
        const request = {
            "filename": file.filename,
            "start": parseInt(start, 10)
        };

        axios.post(uri, request)
            .then(response => this.handlePreviewResponse(response['data']))
            .catch((error) => {
                console.log(error.message);
                this.setState({
                    previewTable: [],
                    maxCols: 0,
                    preview: ["Unable to get file preview..."]
                });
            });
    }

    handlePreviewResponse(response) {
        this.setState({
            preview: response['preview']
        }, () => this.buildPreviewTable());
    }

    buildPreviewTable() {
        this.setState({
            previewTable: [],
            maxCols: 0
        }, () => {
            const separator = this.state.separator;
            if (separator.length <= 0) { return; }
            let maxColumns = this.state.maxCols;
            const preview = this.state.preview;
            for (let i = 0; i < preview.length; i++) {
                const content = preview[i];
                let values = [];
                values = content.replace(' ', '')
                    .split(separator);
                if (values.length > maxColumns) {
                    maxColumns = values.length;
                }
            }

            this.setState({
                maxCols: maxColumns
            }, () => {
                for (let i = 0; i < preview.length; i++) {
                    const content = preview[i];
                    this.buildPreviewRow(content);
                }
            });
        });
    }

    buildPreviewRow(content) {
        const separator = this.state.separator;
        const tableRow = [];
        for (let i = 0; i < this.state.maxCols; i++) {
            tableRow[i] = 'N/A';
        }
        const values = content.replace(' ', '').split(separator);
        for (let i = 0; i < values.length; i++) {
            if (values[i].length > 1) {
                tableRow[i] = values[i];
            }
        }
        const prevTable = this.state.previewTable;
        prevTable.push(tableRow);
        this.setState({
            previewTable: prevTable,
            columns: []
        });
    }

    toggleColumn(index) {
        const selected = this.state.columns;
        if (selected.indexOf(index) >= 0) {
            const toDelete = selected.indexOf(index);
            selected.splice(toDelete, 1);
        } else {
            selected.push(index);
        }
        this.setState({
            columns: selected
        });
    }

    selectFile(file) {
        if (file !== this.state.selected) {
            this.setState({
                selected: file
            }, () => this.getPreview(file));
        }
    }

    removeAlert(alert) {
        const alerts = this.state.alerts;
        const toDelete = alerts.indexOf(alert);
        if (toDelete >= 0) {
            alerts.splice(toDelete, 1);
        }
        this.setState({ alerts: alerts });
    }

    onStartChange(event) {
        this.setState({
            start: event.target.value
        }, () => this.getPreview(this.state.selected));
    }

    onSeparatorChange(event) {
        this.setState({
            separator: event.target.value
        }, () => this.buildPreviewTable());
    }

    onSubmit() {
        this.setState({
            pending: true
        }, () => this.analyzeFile());
    }

    analyzeFile() {
        const validFile = this.state.selected
            && this.state.selected.filename;
        if (validFile) {
            const fn = this.state.selected.filename;
            const columns = this.state.columns;
            let pattern = '{' + this.state.start + '}';
            pattern = pattern + '{' + this.state.separator + '}';
            const base_api = process.env.REACT_APP_BASE_API;
            const uri = base_api + 'analyze';
            const request = {
                "filename": fn,
                "pattern": pattern,
                "columns": columns
            };
            axios.post(uri, request)
                .then(() => {
                    const alerts = this.state.alerts;
                    alerts.push({
                        "message": fn + ' will be analyzed in background.',
                        "type": 0
                    });
                    this.setState({
                        alerts: alerts,
                        loading: true,
                        selected: null,
                        maxCols: 0,
                        preview: [],
                        previewTable: [],
                        columns: [],
                        pending: false
                    }, () => this.getFiles());
                })
                .catch(() => {
                    const alerts = this.state.alerts;
                    alerts.push({
                        "message": 'Unable to analyze ' + fn + '.',
                        "type": -1
                    });
                    this.setState({
                        alerts: alerts,
                        loading: true,
                        selected: null,
                        maxCols: 0,
                        preview: [],
                        previewTable: [],
                        columns: [],
                        pending: false
                    }, () => this.getFiles());
                });
        }
    }
}

export default Analyze;