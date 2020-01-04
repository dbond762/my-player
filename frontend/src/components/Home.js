import React, { Component } from 'react';

import VideoListContainer from './VideoListContainer';

class Home extends Component {
  render() {
    return (
      <VideoListContainer query="" className={this.props.className} />
    );
  }
}

export default Home;
