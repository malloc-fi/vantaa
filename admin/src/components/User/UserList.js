import React from 'react';
import RequireAuthentication from 'decorators/RequireAuthentication';

@RequireAuthentication
class UserList extends React.Component {

  constructor(props) {
    super(props)
    this.state = this._getUsers();
  }

  _getUsers() {
    return [
      {
        id: 1,
        name: 'unique',
        email: 'unique@example.com'
      },
      {
        id: 2,
        name: 'unique2',
        email: 'unique2@example.com'
      }
    ];
  }
}
