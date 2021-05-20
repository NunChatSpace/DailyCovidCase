import React, { Component } from 'react';
import { Grid, Container } from '@material-ui/core';
import MapDistribution from './Map/MapDistribution';
import Chart from './Chart/Chart';
import DailyCases from './DailyCases';
import CovidCasesStatistics from './CovidCasesStatistics';
import { connect } from 'react-redux';

export class Main extends Component {

    render() {
        console.log("Main render")
        return (
            <Container>
                <div>
                    <Grid container spacing={3}>
                        <Grid item xs={12} sm={12}>
                            <div id='mapContainer' style={{ margin: 2, height: 400 }}>
                                <MapDistribution />
                            </div>
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <div id='graphContainer' style={{ margin: 2 }}>
                                <CovidCasesStatistics/>
                            </div>
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <div style={{ margin: 2 }}>
                                {/* <Chart /> */}
                                <label style={{color: "white"}}>
                                    Cannot access API from provider
                                </label>
                            </div>
                        </Grid>
                        <Grid item xs={12} sm={12}>
                            <DailyCases/>
                        </Grid>
                    </Grid>
                </div>
            </Container>
        )
    }
}

const mapStateToProps = ({ mainReducer }) => ({ mainReducer });

const mapDispatchToProps = {};

export default connect(mapStateToProps, mapDispatchToProps)(Main);