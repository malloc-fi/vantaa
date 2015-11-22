import React from 'react';
import RequireAuthentication from 'decorators/RequireAuthentication';

@RequireAuthentication
class SuperSecure extends React.Component {

  constructor(props) {
    super(props);
  }

  render() {
    return(
      <h1>Super Secure</h1>
    );
  }
}

export default SuperSecure;
