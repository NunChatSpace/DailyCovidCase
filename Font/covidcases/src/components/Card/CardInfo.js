import { Grid, Paper} from '@material-ui/core'
import React, { Component } from 'react'

export default class CardInfo extends Component {

    render() {
        return (
            <div style={{padding:4}}>
                <Paper elevation={3}>
                    <Grid container spacing={3}>
                        <Grid item xs={12} sm={12}>
                            <label style={{ display: 'flex', justifyContent: 'center' }}>
                                {this.props.header}
                            </label>
                        </Grid>
                        <Grid item xs={12} sm={12}>
                            <label style={{ display: 'flex', justifyContent: 'center', fontSize: 40 }}>
                                {this.props.text}
                            </label>
                        </Grid>
                        <Grid item xs={12} sm={12}>
                            <label style={{ display: 'flex', justifyContent: 'center' }}>
                                {this.props.footerText}
                            </label>
                        </Grid>
                    </Grid>
                </Paper>
            </div>
        )
    }
}
