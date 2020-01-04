import { createStore, combineReducers } from 'redux';

let searchHistoryReducer = function(state = [], action) {
    if (action.type === 'PUSH_SEARCH_HISTORY') {
        let newState = state.concat([action.query]);
        return newState;
    }
    return state;
}

let userReducer = function(state = {}, action) {
    if (action.type === 'USER_LOGIN') {
        return {
            isLoggedIn: true,
            login: action.user.login,
            token: action.user.session_id,
        };
    } else if (action.type === 'USER_LOGOUT') {
        return {
            isLoggedIn: false,
            login: '',
            token: '',
        }
    }
    return state;
}

const redusers = combineReducers({
    searchHistory: searchHistoryReducer,
    user: userReducer,
});

const store = createStore(redusers);

export default store;
