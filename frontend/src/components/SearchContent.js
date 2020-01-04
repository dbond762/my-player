import React from 'react';
import { withRouter } from 'react-router-dom';

import VideoListContainer from './VideoListContainer';

function SearchContent(props) {
    return (
        <VideoListContainer query={props.match.params.query} className={props.className} />
    );
}

export default withRouter(SearchContent);
