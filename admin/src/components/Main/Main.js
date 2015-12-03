require('./Main.css');

import React from 'react';
import Header from 'components/Header/Header';
import { RouteHandler } from 'react-router';

class Main extends React.Component {

  constructor() {
    super()
  }

  render() {
    return(
      <div id="app-container">
        <Header />
        <RouteHandler/>
      </div>
    );
  }
}

export default Main;
