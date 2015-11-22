// Bootstrapping module
import React from 'react';
import Router from 'react-router';
import routes from 'routes';
import RouterContainer from 'services/RouterContainer';
import AuthActions from 'actions/AuthActions';

var router = Router.create({ routes });
RouterContainer.set(router);

let jwt = localStorage.getItem('jwt');
if (jwt) {
  AuthActions.loginUser(jwt);
}

// Router.run(routes, Router.HistoryLocation, (Root, state) => {
//   React.render(<Root {...state}/>, document.getElementById('content'));
// });

router.run(function(Handler) {
  React.render(<Handler />, document.getElementById('content'));
});
