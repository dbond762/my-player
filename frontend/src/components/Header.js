import React, { Component } from 'react';
import { Link } from 'react-router-dom';

import './Header.css';
import logo from '../assets/img/logo.png';
import Search from './Search';
import UserData from './UserData/UserData';

class Header extends Component {
  render() {
    return (
      <header className={`header ${this.props.className}`}>
        <Link to="/">
          <img className="header__logo" src={logo} alt="logo"/>
        </Link>
        <Search className="header__search"/>
        <UserData className="header__user-data" />
      </header>
    );
  }
}

export default Header;
