import React, { Component } from 'react';
import { connect } from 'react-redux';
import store from '../../store'

import './UserData.css';
import LoginButton from './LoginButton';
import LogoutButton from './LogoutButton';

class UserData extends Component {
  render() {
    if (this.props.user.isLoggedIn) {
      return (
        <LogoutButton user={this.props.user} className={this.props.className} logout={this.logout.bind(this)} />
      );
    } else {
      return (
        <LoginButton className={this.props.className} />
      );
    }
  }

  logout = () => {
    store.dispatch({
      type: 'USER_LOGOUT',
    });
  };
}

const mapStateToProps = function(store) {
  return {
    user: store.user,
  };
}

export default connect(mapStateToProps)(UserData);
