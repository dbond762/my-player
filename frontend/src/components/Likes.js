import React, { Component } from 'react';

class Likes extends Component {
  render() {
    return (
      <div className={`likes ${this.props.className}`}>
        <span>Likes</span>
      </div>
    );
  }
}

export default Likes;
