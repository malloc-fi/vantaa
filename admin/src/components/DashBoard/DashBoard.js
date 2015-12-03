import React from 'react';
import RequireAuthentication from 'decorators/RequireAuthentication';

@RequireAuthentication
class DashBoard extends React.Component {

  constructor(props) {
    super(props)
  }

  render() {
    return(
      <div className="interation-container">
        <div className="interaction-area"></div>
        <div className="interaction-message"></div>
        <div className="interaction-commands"></div>
      </div>
    );
  }
}

export default DashBoard;
