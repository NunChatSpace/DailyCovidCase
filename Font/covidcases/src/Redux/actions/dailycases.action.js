import axios from "axios";
import ReactDOM from 'react-dom';
import { DAILYCASES_API, DAILYCASES_LOADING, DAILYCASES_LOADING_FAILED, DAILYCASES_LOADING_SUCCESS } from "../../contants";

export const setDownloadingStatetoFetching = () => ({
    type: DAILYCASES_LOADING
});

export const setDownloadingStatetoFailed = () => ({
    type: DAILYCASES_LOADING_FAILED
});

export const setDownloadingStatetoSuccess = (payload) => ({
    type: DAILYCASES_LOADING_SUCCESS,
    payload
});

export const loadDailyCases = () => {
    return async dispatch => {
        try {
            dispatch(setDownloadingStatetoFetching);
            let result = await axios({
                baseURL: DAILYCASES_API,
                method: 'GET',
                headers: {
                    'Access-Control-Allow-Origin': '*',
                    'Access-Control-Allow-Methods': 'GET,POST,HEAD,PUT,DELETE,PATCH',
                    'Access-Control-Allow-Headers': 'access-control-allow-origin, Origin, Content-Type, Accept, Content-Length, Authorization',
                }
            });
            // console.log(result);
            if (result.data.Status === 200) {
                dispatch(setDownloadingStatetoSuccess(result));
            } else {
                dispatch(setDownloadingStatetoFailed());
            }
        } catch (error) {
            dispatch(setDownloadingStatetoFailed());
        }
    }
}
