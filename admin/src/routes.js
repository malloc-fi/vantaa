import React from 'react';
import {Route} from 'react-router';

import Main from 'components/main';
import SuperSecure from 'components/restricted';

const routes = (
  <Route handler={Main}>
    <Route name="secure" handler={SuperSecure} />
  </Route>
);

export default routes;
