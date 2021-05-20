import React, { Component } from 'react'
import Chart from './Chart/Chart';
import { connect } from 'react-redux'
import { loadCasesStat } from '../Redux/actions/covidcasesstat.action'

export class CovidCasesStatistics extends Component {
    componentDidMount() {
        this.props.loadCasesStat();
    }

    render() {
        var options = {
            rangeSelector: {
                selected: 3
            },
            title: {
                text: "ยอดผู้ป่วยสะสม"
            },
            series: []
        };

        if (this.props.covidcasesstatReducer.result){ 
            var tmpData = this.props.covidcasesstatReducer.result.data.Data;
            var tmpDataLength = this.props.covidcasesstatReducer.result.data.DataLength;
            var confirmed = {
                name: "ผู้ป่วยรายใหม่",
                data: []
            };
            var recovered = {
                name: "รักษาหายแล้ว",
                data: []
            };
            var hospitalized = {
                name: "รักษาอยู่ รพ",
                data: []
            };
            var deaths = {
                name: "เสียชีวิต",
                data: []
            };

            for(var i = 0; i < tmpDataLength; i++)
            {
                confirmed.data.push([new Date(tmpData[i].Date).getTime(), tmpData[i].Confirmed])
                recovered.data.push([new Date(tmpData[i].Date).getTime(), tmpData[i].Recovered])
                hospitalized.data.push([new Date(tmpData[i].Date).getTime(), tmpData[i].Hospitalized])
                deaths.data.push([new Date(tmpData[i].Date).getTime(), tmpData[i].Deaths])

                options = {
                    rangeSelector: {
                        selected: 3
                    },
                    title: {
                        text: "ยอดผู้ป่วยสะสม"
                    },
                    series: [confirmed, recovered, hospitalized, deaths]
                }
                // viewData.push([new Date(tmpData[i].Date), tmpData[i].Confirmed, tmpData[i].Recovered, tmpData[i].Hospitalized, tmpData[i].Deaths]);
            }
        }

        return (
            <Chart options={options}/>
        )
    }
}

const mapStateToProps = ({ covidcasesstatReducer }) => ({ covidcasesstatReducer });

const mapDispatchToProps = {
    loadCasesStat
};

export default connect(mapStateToProps, mapDispatchToProps)(CovidCasesStatistics);