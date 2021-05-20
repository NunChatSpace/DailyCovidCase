import {
    MAIN_REFRESH
} from "../../contants";

const initialState = {
    isRefresh: false
}

export default (state = initialState, { type }) => {
    switch (type) {
        case MAIN_REFRESH:
            return { ...state, isRefresh: true};
        default:
            return { ...state, isRefresh: false};
    }
}