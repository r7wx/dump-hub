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

class Alerts extends React.Component {
    render() {
        const alerts = this.props.alerts;
        const items = [];
        for (let i = 0; i < alerts.length; i++) {
            if (alerts[i].type === 0) {
                items.push(
                    <div className="alert alert-success" role="alert">
                        <div className="alert-items">
                            <div className="alert-item static">
                                <span className="alert-text"> <b>Success:</b> {alerts[i].message} </span>
                            </div>
                        </div>
                        <button type="button"
                            className="close"
                            aria-label="Close"
                            onClick={() => this.props.removeAlert(alerts[i])}>
                            <clr-icon aria-hidden="true" shape="close"></clr-icon>
                        </button>
                    </div >
                );
            } else {
                items.push(
                    <div className="alert alert-danger" role="alert">
                        <div className="alert-items">
                            <div className="alert-item static">
                                <span className="alert-text"> <b>Error:</b> {alerts[i].message} </span>
                            </div>
                        </div>
                        <button type="button"
                            className="close"
                            aria-label="Close"
                            onClick={() => this.props.removeAlert(alerts[i])}>
                            <clr-icon aria-hidden="true" shape="close"></clr-icon>
                        </button >
                    </div >
                );
            }
        }

        if (items.length > 0) {
            items.push(
                <br />
            );
        }

        return (
            <React.Fragment>
                {items}
            </React.Fragment>
        );
    }
}

export default Alerts;