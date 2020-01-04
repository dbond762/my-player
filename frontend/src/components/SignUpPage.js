import React, { Component } from 'react';
import store from '../store';
import { withRouter } from 'react-router-dom';
import axios from 'axios';

import './LoginPage.css';

class SignUpPage extends Component {
  state = {
    login: '',
    password: '',
    passwordAgain: '',
  };

  handleChange = () => {
    let login = document.querySelector('#login').value;
    let password = document.querySelector('#password').value;
    let passwordAgain = document.querySelector('#password-again').value;
    let submitButton = document.querySelector('.login-page__submit');

    if (password !== passwordAgain || password === '' || login === '') {
      submitButton.setAttribute('disabled', '');
    } else {
      submitButton.removeAttribute('disabled');
    }

    this.setState({
      login: login,
      password: password,
      passwordAgain: passwordAgain,
    });
  };

  componentDidMount() {
    document.querySelector('.login-page__submit').setAttribute('disabled', '');
  }

  reset = (event) => {
    this.setState({
      login: '',
      password: '',
      passwordAgain: '',
    });

    document.querySelector('.login-page__submit').setAttribute('disabled', '');

    event.preventDefault();
  };

  signup = (event) => {
    axios.post('/api/signup', {
      user: {
        login: this.state.login,
        password: this.state.password,
      },
    }).then((resp) => {
      store.dispatch({
        type: 'USER_LOGIN',
        user: {
          isLoggedIn: true,
          login: this.state.login,
          token: resp.data.user.session_id,
        },
      });

      this.props.history.goBack();
    }).catch((error) => {
      console.error(error);
    });

    event.preventDefault();
  };

  render() {
    return (
      <div className={`login-page ${this.props.className}`}>
        <form className="login-page__form">
          <h2 className="login-page__title">Sign up</h2>
          <div className="login-page__fieldset">
            <div className="login-page__field">
              <label className="login-page__label" htmlFor="login">Login</label>
              <input className="login-page__input" type="text" id="login" value={this.state.login} onChange={this.handleChange.bind(this)} />
            </div>
            <div className="login-page__field">
              <label className="login-page__label" htmlFor="password">Password</label>
              <input className="login-page__input" type="password" id="password" value={this.state.password} onChange={this.handleChange.bind(this)} />
            </div>
            <div className="login-page__field">
              <label className="login-page__label" htmlFor="password-again">Password again</label>
              <input className="login-page__input" type="password" id="password-again" value={this.state.passwordAgain} onChange={this.handleChange.bind(this)} />
            </div>
            <div className="login-page__buttons">
              <input className="login-page__btn login-page__reset" type="reset" value="Reset" onClick={(e) => this.reset(e)} />
              <input className="login-page__btn login-page__submit" type="button" value="Sign up" onClick={(e) => this.signup(e)} />
            </div>
          </div>
        </form>
      </div>
    );
  }
}

export default withRouter(SignUpPage);
