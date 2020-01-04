import React, { Component } from 'react';
import { Link } from 'react-router-dom';

class LoginButton extends Component {
  render() {
    return (
      <div className={`user-data ${this.props.className}`}>
        <Link className="user-data__link" to="/login">Login</Link>
        <Link className="user-data__link" to="/signup">Sign up</Link>
      </div>
    );
  }
}

export default LoginButton;
