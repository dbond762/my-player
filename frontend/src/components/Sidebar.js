import React, { Component } from 'react';
import { Link } from 'react-router-dom';

import './Sidebar.css';

class Sidebar extends Component {
  render() {
    return (
      <div className={`sidebar ${this.props.className}`}>
        <Link to="/" className="sidebar__item">Home</Link>
        <Link to="/likes" className="sidebar__item">Likes</Link>
      </div>
    );
  }
}

export default Sidebar;
