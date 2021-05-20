import React, { Component } from 'react';
import Chart from "react-google-charts";
import './componentStyle.css'
export default class BarChart1 extends Component {
    constructor(props){
        super(props);
        this.state = {
            data: props.data
        }
    }

    render() {
        return (
            <div className='wrapper'>
                <Chart
                    chartType="ColumnChart"
                    loader={<div>Loading Chart</div>}
                    data={this.state.data}
                    options={{
                        chartArea: { width: '50%' },
                        hAxis: {
                            title: 'Total Population',
                            minValue: 0,
                        },
                        vAxis: {
                            title: 'Cases',
                        },
                    }}
                    legendToggle
                />
            </div>
        )
    }
}
