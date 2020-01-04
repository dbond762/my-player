import React, { Component } from 'react';
import store from '../store';
import { withRouter } from 'react-router-dom';

import './LoginPage.css';

class LoginPage extends Component {
  state = {
    login: '',
    password: '',
  };

  handleChangeLogin = (event) => {
    this.setState({
      login: event.target.value,
    });
  };

  handleChangePassword = (event) => {
    this.setState({
      password: event.target.value,
    });
  };

  reset = (event) => {
    this.setState({
      login: '',
      password: '',
    });
    event.preventDefault();
  };

  login = (event) => {
    fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        user: {
          login: this.state.login,
          password: this.state.password,
        },
      }),
    }).then((resp) => {
      return resp.json();
    }).then((data) => {
      console.log(data);
      store.dispatch({
        type: 'USER_LOGIN',
        user: {
          isLoggedIn: true,
          login: this.state.login,
          token: data.user.session_id,
        },
      });

      this.props.history.goBack();
    }).catch(error => {
      console.error(error);
    });

    event.preventDefault();
  };

  render() {
    return (
      <div className={`login-page ${this.props.className}`}>
        <form className="login-page__form">
          <h2 className="login-page__title">Login</h2>
          <div className="login-page__fieldset">
            <div className="login-page__field">
              <label className="login-page__label" htmlFor="login">Login</label>
              <input className="login-page__input" type="text" id="login" value={this.state.login} onChange={this.handleChangeLogin.bind(this)} />
            </div>
            <div className="login-page__field">
              <label className="login-page__label" htmlFor="password">Password</label>
              <input className="login-page__input" type="password" id="password" value={this.state.password} onChange={this.handleChangePassword.bind(this)} />
            </div>
            <div className="login-page__buttons">
              <input className="login-page__btn login-page__reset" type="reset" value="Reset" onClick={(e) => this.reset(e)} />
              <input className="login-page__btn login-page__submit" type="button" value="Login" onClick={(e) => this.login(e)} />
            </div>
          </div>
        </form>
      </div>
    );
  }
}

export default withRouter(LoginPage);
