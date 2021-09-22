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

import Paginator from '../core/paginator/paginator.js';
import Loading from '../core/loading/loading.js';
import DataView from './data-view/data-view.js';
import React from 'react';
import axios from 'axios';

class Data extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            page: 1,
            loading: true,
            data: [],
            dataTot: 0,
            toDelete: null
        };
        this.onPageChange = this.onPageChange.bind(this);
        this.onDelete = this.onDelete.bind(this);
        this.onDeleteClick = this.onDeleteClick.bind(this);
        this.onDeleteCancel = this.onDeleteCancel.bind(this);
    }

    render() {
        return (
            <React.Fragment>
                <Loading isLoading={this.state.loading} />
                <DataView data={this.state.data}
                    toDelete={this.state.toDelete}
                    onDelete={this.onDelete}
                    onDeleteClick={this.onDeleteClick}
                    onDeleteCancel={this.onDeleteCancel} />
                <Paginator currentPage={this.state.page}
                    totalItems={this.state.dataTot}
                    pageSize={25}
                    onPageChange={this.onPageChange} />
            </React.Fragment>
        );
    }

    componentDidMount() {
        this.getStatus();
        this.interval = setInterval(() => this.getStatus(), 5000);
    }

    componentWillUnmount() {
        clearInterval(this.interval);
    }

    getStatus() {
        const base_api = process.env.REACT_APP_BASE_API;
        const uri = base_api + 'status';

        this.setState({ loading: true });
        const request = {
            "page": this.state.page
        };
        axios.post(uri, request)
            .then(response => this.handleResponse(response['data']))
            .catch(() => { this.setState({ loading: true }); });
    }

    handleResponse(response) {
        this.setState({
            loading: false,
            data: response['results'],
            dataTot: response['tot']
        });
    }

    onPageChange(delta) {
        const newPage = this.state.page + delta;
        this.setState({
            page: newPage
        }, () => this.getStatus());
    }

    onDelete(toDelete) {
        const base_api = process.env.REACT_APP_BASE_API;
        const uri = base_api + 'delete';

        this.setState({ loading: true });
        const request = {
            "checksum": toDelete.checksum
        };
        axios.post(uri, request)
            .then(() => {
                this.setState({
                    loading: false,
                    toDelete: null
                }, () => this.getStatus());
            })
            .catch(() => {
                this.setState({
                    loading: true,
                    toDelete: null
                });
            });
    }

    onDeleteClick(toDelete) {
        this.setState({
            toDelete: toDelete
        });
    }

    onDeleteCancel() {
        this.setState({
            toDelete: null
        });
    }
}

export default Data;