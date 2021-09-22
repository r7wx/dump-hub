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
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link,
    NavLink
} from 'react-router-dom';
import Analyze from './analyze/analyze.js';
import Search from './search/search.js';
import Upload from './upload/upload.js';
import Data from './data/data.js';
import './app.css';

class App extends React.Component {
    render() {
        return (
            < Router >
                <div className="main-container">
                    <header className="header header-6">
                        <div className="branding">
                            <Link to="/" className="nav-link nav-icon" >
                                <clr-icon shape="block"></clr-icon>
                            </Link>
                        </div >
                        <div className="header-nav">
                            <NavLink to="/" exact className="nav-link nav-icon" activeClassName="active">
                                <clr-icon shape="search"></clr-icon>
                            </NavLink>
                            <NavLink to="data" className="nav-link nav-icon" activeClassName="active">
                                <clr-icon shape="block"></clr-icon>
                            </NavLink>
                            <NavLink to="analyze" className="nav-link nav-icon" activeClassName="active">
                                <clr-icon shape="plus-circle"></clr-icon>
                            </NavLink>
                            <NavLink to="upload" className="nav-link nav-icon" activeClassName="active">
                                <clr-icon shape="folder"></clr-icon>
                            </NavLink>
                        </div>
                    </header >
                    <div className="content-container">
                        <div className="content-area">
                            <Switch>
                                <Route path="/" exact>
                                    <Search />
                                </Route>
                                <Route path="/data" exact>
                                    <Data />
                                </Route>
                                <Route path="/analyze" exact>
                                    <Analyze />
                                </Route>
                                <Route path="/upload" exact>
                                    <Upload />
                                </Route>
                            </Switch>
                        </div>
                    </div>
                </div >
            </Router >
        );
    }
}

export default App;