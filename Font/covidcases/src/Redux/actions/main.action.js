import { MAIN_LOADING,  MAIN_LOADING_FAILED } from "../../contants";


export const setDownloadingStatetoFetching = () => ({
    type: MAIN_LOADING,
});

export const setDownloadingStatetoFailed = () => ({
    type: MAIN_LOADING_FAILED,
});

export const setDownloadingStatetoSuccess = (type, payload) => ({
    type: type,
    payload
});

// export const InitData = () => {
//     return dispatch => {
//         try {
//             dispatch(setDownloadingStatetoFetching());
//             let result = null;
//             loadCasesStat().then((response) => {
//                 if(result) {
//                     result = {
//                         DailyCases: result.DailyCases,
//                         CasesStat: response
//                     }
//                 }
//                 else{
//                     result = {
//                         CasesStat: response
//                     }
//                 }
//                 dispatch(setDownloadingStatetoSuccess(MAIN_LOADING_CASESSTAT_SUCCESS, result))
//             }).catch((err) => {
//                 dispatch(setDownloadingStatetoFailed())
//             });

//             loadDailyCases().then((response) => {
//                 if(result) {
//                     result = {
//                         CasesStat: result.CasesStat,
//                         DailyCases: response
//                     }
//                 }
//                 else{
//                     result = {
//                         DailyCases: response
//                     }
//                 }
//                 dispatch(setDownloadingStatetoSuccess(MAIN_LOADING_DAILYCASES_SUCCESS, result))
//             }).catch((err) => {
//                 dispatch(setDownloadingStatetoFailed())
//             });
//         } catch (error) {
//             dispatch(setDownloadingStatetoFailed());
//         }
//     }
// }

// async function loadCasesStat(){

//     let result = await axios({
//         baseURL: CASESSTAT_API,
//         method: 'GET',
//         headers: {
//             'Access-Control-Allow-Origin': '*',
//             'Access-Control-Allow-Methods': 'GET,POST,HEAD,PUT,DELETE,PATCH',
//             'Access-Control-Allow-Headers': 'access-control-allow-origin, Origin, Content-Type, Accept, Content-Length, Authorization',
//         }
//     });

//     return result;
// }

// async function loadDailyCases() {

//     let result = await axios({
//         baseURL: DAILYCASES_API,
//         method: 'GET',
//         headers: {
//             'Access-Control-Allow-Origin': '*',
//             'Access-Control-Allow-Methods': 'GET,POST,HEAD,PUT,DELETE,PATCH',
//             'Access-Control-Allow-Headers': 'access-control-allow-origin, Origin, Content-Type, Accept, Content-Length, Authorization',
//         }
//     });

//     return result;
// }
