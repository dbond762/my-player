import React, { Component } from 'react';

class LogoutButton extends Component {
  render() {
    return (
      <div className={`user-data ${this.props.className}`}>
        <span className="user-data__user">{this.props.user.login}</span>
        <a type="button" href="#" className="user-data__link" onClick={this.props.logout}>Logout</a>
      </div>
    );
  }
} 

export default LogoutButton;
