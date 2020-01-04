import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Header from './Header';
import Sidebar from './Sidebar';
import Home from './Home';
import Likes from './Likes';
import './App.css';
import SearchContent from './SearchContent';
import LoginPage from './LoginPage';
import SignUpPage from './SignUpPage';

class App extends Component {
  componentDidMount() {
    window.addEventListener('scroll', () => {
      let header = document.querySelector('.app__header');
      if (window.pageYOffset !== 0) {
        header.classList.add('app__header--fixed');
      } else {
        header.classList.remove('app__header--fixed');
      }
    });
  }

  render() {
    return (
      <div className="app">
        <Router>
          <Header className="app__header" />
          <div className="app__page">
            <Sidebar className="app__sidebar" />
            <Switch>
              <Route exact path="/" component={() => <Home className="app__content" />} />
              <Route path="/likes" component={() => <Likes className="app__content" />} />
              <Route path="/search/:query" component={() => <SearchContent className="app__content" />} />
              <Route path="/login" component={() => <LoginPage className="app__content" />} />
              <Route path="/signup" component={() => <SignUpPage className="app__content" />} />
            </Switch>
          </div>
        </Router>
      </div>
    );
  }
}

export default App;
