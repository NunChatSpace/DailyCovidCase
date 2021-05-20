import React, { Component } from 'react'
import './componentStyle.css'
import Highcharts from "highcharts/highstock";
import HighchartsReact from "highcharts-react-official";
window.Highcharts = Highcharts;

export default class Chart extends Component {
    render() {
        return (
            <div className='wrapper'>
                {/* <label style={{ color: 'whitesmoke' }}>{this.props.title}</label> */}
                <HighchartsReact
                    highcharts={Highcharts}
                    constructorType={"stockChart"}
                    options={this.props.options}
                />
            </div>
        )
    }
}
