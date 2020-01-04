import React, { Component } from 'react';

import './VideoList.css';

class VideoList extends Component {
  render() {
    return (
      <div className={`video-list ${this.props.className}`}>
        {this.props.videos.map(this.createListItem)}
      </div>
    );
  }

  createListItem(video) {
    return (
      <div className="video-list__item video" key={video.id}>
        <div>
          <img className="video__thumbnail" src={video.thumbnail} alt={video.title}/>
        </div>
        <div className="video__data">
          <h3 className="video__name">{video.title}</h3>
          <span className="video__date">
            {new Date(video.pub_date).toLocaleString('ru-UA', { day: 'numeric', month: 'short', year: 'numeric' })}
          </span>
          <p className="video__description">{video.description}</p>
          <span className="video__like"></span>
        </div>
      </div>
    )
  }
}

export default VideoList;
