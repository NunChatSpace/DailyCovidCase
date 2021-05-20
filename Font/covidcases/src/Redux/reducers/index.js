import { combineReducers } from "redux";
import dailycasesReducer from './dailycases.reducer';
import covidcasesstatReducer from './covidcasesstat.reducer';
import mainReducer from './main.reducer';

export default combineReducers({
    mainReducer,
    dailycasesReducer,
    covidcasesstatReducer
})