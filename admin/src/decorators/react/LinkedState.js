import React from 'react';
import ReactLink from 'react/lib/ReactLink';

export var LinkedState = (component) => {
  return class extends component {

    linkState = key => {
      return {
        value: this.state[key],
        requestChange: this._generateChangeHandler(key)
      };
    }

    _generateChangeHandler = key => {
      return (value) => {
        var stateObject = {};
        stateObject[key] = value;
        this.setState(stateObject);
      };
    }
  };
};
