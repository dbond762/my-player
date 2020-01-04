import React, { Component } from 'react';

import VideoList from './VideoList';

class VideoListContainer extends Component {
  state = {
    videos: [],
  }

  componentDidMount() {
    fetch(`/api/search/${this.props.query}`).then((response) => {
      return response.json();
    }).then((videos) => {
      this.setState({
        videos: videos,
      });
    });
  }

  render() {
    return (
      <VideoList videos={this.state.videos} className={this.props.className} />
    );
  }
}

export default VideoListContainer;
