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
import Results from './results/results.js';
import Bar from './bar/bar.js';
import React from 'react';
import axios from 'axios';
class Search extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            query: '',
            results: [],
            loading: true,
            page: 1,
            resultsTot: 0
        };
        this.onQueryChange = this.onQueryChange.bind(this);
        this.onPageChange = this.onPageChange.bind(this);
    }

    render() {
        return (
            <React.Fragment>
                <Bar query={this.state.query}
                    onQueryChange={this.onQueryChange} />
                <Loading isLoading={this.state.loading} />
                <table className="table">
                    <tbody>
                        <Results results={this.state.results}
                            onQueryChange={this.onQueryChange} />
                    </tbody>
                </table>
                <Paginator currentPage={this.state.page}
                    totalItems={this.state.resultsTot}
                    pageSize={25}
                    onPageChange={this.onPageChange} />
            </React.Fragment>
        );
    }

    componentDidMount() {
        this.search(this.state.query);
    }

    onQueryChange(newQuery) {
        this.setState({
            page: 1,
            query: newQuery
        },
            () => { this.search(); }
        );
    }

    search() {
        const base_api = process.env.REACT_APP_BASE_API;
        const uri = base_api + 'search';

        this.setState({ loading: true });
        const request = {
            "query": this.state.query,
            "page": this.state.page
        };
        axios.post(uri, request)
            .then(response => this.handleSearchResponse(response.data))
            .catch(() => {
                this.setState({
                    results: [],
                    resultsTot: 0,
                    loading: false
                });
            });
    }

    handleSearchResponse(response) {
        this.setState({
            results: response.results,
            resultsTot: response.tot,
            loading: false
        });
    }

    onPageChange(delta) {
        const newPage = this.state.page + delta;
        this.setState({
            page: newPage
        }, () => this.search());
    }
}

export default Search;