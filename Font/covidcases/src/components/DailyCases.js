import { Grid } from '@material-ui/core'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { loadDailyCases } from '../Redux/actions/dailycases.action'
import CardInfo from './Card/CardInfo'

export class DailyCases extends Component {
    constructor(props){
        super(props)
    }

    componentDidMount() {
        this.props.loadDailyCases();
    }

    render() {
        
        const viewData = this.props.dailycasesReducer.result;
        var confirmed = 0;
        var newConfirmed = 0;
        var recovered = 0;
        var newRecovered = 0;
        var hospitalized = 0;
        var newHospitalized = 0;
        var deaths = 0;
        var newDeaths = 0;
        
        if (viewData) {
            
            confirmed = viewData.data.Data.Confirmed;
            newConfirmed = viewData.data.Data.NewConfirmed > 0 ? `(+${viewData.data.Data.NewConfirmed})` : `(${viewData.data.Data.NewConfirmed})`;
            recovered = viewData.data.Data.Recovered;
            newRecovered = viewData.data.Data.NewRecovered > 0 ? `(+${viewData.data.Data.NewRecovered})` : `(${viewData.data.Data.NewRecovered})`;
            hospitalized = viewData.data.Data.Hospitalized;
            newHospitalized = viewData.data.Data.NewHospitalized > 0 ? `(+${viewData.data.Data.NewHospitalized})` : `(${viewData.data.Data.NewHospitalized})`;
            deaths = viewData.data.Data.Deaths;
            newDeaths = viewData.data.Data.Deaths > 0 ? `(+${viewData.data.Data.NewDeaths})` : `(${viewData.data.Data.NewDeaths})`;
        }

        return (
            <Grid container spacing={3}>
                <Grid item xs={6} sm={3}>
                    <CardInfo text={confirmed} footerText={newConfirmed} header='ผู้ป่วยสะสม'></CardInfo>
                </Grid>
                <Grid item xs={6} sm={3}>
                    <CardInfo text={recovered} footerText={newRecovered} header='รักษาหายแล้ว'></CardInfo>
                </Grid>
                <Grid item xs={6} sm={3}>
                    <CardInfo text={hospitalized} footerText={newHospitalized} header='รักษาอยู่ใน รพ.'></CardInfo>
                </Grid>
                <Grid item xs={6} sm={3}>
                    <CardInfo text={deaths} footerText={newDeaths} header='เสียชีวิต'></CardInfo>
                </Grid>
            </Grid>
        )
    }
}

const mapStateToProps = ({ dailycasesReducer }) => ({ dailycasesReducer });

const mapDispatchToProps = {
    loadDailyCases
};

export default connect(mapStateToProps, mapDispatchToProps)(DailyCases);