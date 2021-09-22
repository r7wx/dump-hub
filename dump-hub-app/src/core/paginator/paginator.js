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
import './paginator.css';

class Paginator extends React.Component {
    render() {
        const total = this.props.totalItems;
        const current = this.props.currentPage;
        const size = this.props.pageSize;

        if (total > 0) {
            return (
                <div className="pagination-container">
                    <button className={`page-control ${current === 1 ? "disabled" : ""}`}
                        onClick={() => this.props.onPageChange(-1)}> &lt;&lt; Previous
                    </button>
                    <b className="page-current">{current}</b>
                    <button className={`page-control ${!(total >= size * current) ? "disabled" : ""}`}
                        onClick={() => this.props.onPageChange(1)}> Next &gt;&gt;
                    </button>
                </div >
            );
        }

        return null;
    }
}

export default Paginator;