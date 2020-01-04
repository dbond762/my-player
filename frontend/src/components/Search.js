import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';

import './Search.css';
import store from '../store';

class Search extends Component {
  state = {
    query: '',
  };

  handleChange = (event) => {
    this.setState({
      query: event.target.value,
    });
  };

  search = (event) => {
    this.props.history.push(`/search/${this.state.query}`);
    store.dispatch({
      type: 'PUSH_SEARCH_HISTORY',
      query: this.state.query,
    });
    event.preventDefault();
  };

  render() {
    return (
      <div className={`search ${this.props.className}`}>
        <input type="search" className="search__field" value={this.state.query} onChange={this.handleChange.bind(this)}/>
        <input type="button" className="search__btn" value="Go" onClick={this.search.bind(this)}/>
      </div>
    );
  }
}

const mapStateToProps = function(store) {
  return {
    searchHistory: store.searchHistory,
  };
}

export default connect(mapStateToProps)(withRouter(Search));
