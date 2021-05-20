import React, { Component } from 'react'
import GoogleMapReact from 'google-map-react';
import './MapDistribution.css'
export default class MapDistribution extends Component {
    static defaultProps = {
        center: {
            lat: 59.95,
            lng: 30.33
        },
        zoom: 11
    };

    constructor(props) {
        super(props);
        this.state = {
            center: {
                lat: -1.2884,
                lng: 36.8233
            },
            zoom: 14
        }
    }

    render() {
        return (
                <div style={{ height: '100%', width: '100%' }}>
                    <GoogleMapReact
                        className='googleMap'
                        bootstrapURLKeys={{ key: 'AIzaSyCZLMXTxjhX_ZEv0O9moAAyelJP4PGbbUI'}}
                        defaultCenter={this.props.center}
                        defaultZoom={this.props.zoom}
                    >
                    </GoogleMapReact>
                </div>

        )
    }
}

// export default GoogleApiWrapper({
//     apiKey: 'AIzaSyCZLMXTxjhX_ZEv0O9moAAyelJP4PGbbUI'
// })(MapDistribution);
