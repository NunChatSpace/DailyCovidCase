import { 
    DAILYCASES_LOADING, 
    DAILYCASES_LOADING_FAILED, 
    DAILYCASES_LOADING_SUCCESS 
} from "../../contants";

const initialState = {
    isFetching: false,
    isError: false,
    result: null
}

export default (state = initialState, { type, payload }) => {
    switch (type) {
        case DAILYCASES_LOADING:
            return { ...state, isFetching: true, isError: false, result: null };
        case DAILYCASES_LOADING_FAILED:
            return { ...state, isFetching: false, isError: true, result: null };
        case DAILYCASES_LOADING_SUCCESS:
            return { ...state, isFetching: false, isError: false, result: payload };
        default:
            return state;
    }
}