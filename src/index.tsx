import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import axios from 'axios';
import 'uswds';
import App from './views/App';
import store from './store';
import './index.scss';
import * as serviceWorker from './serviceWorker';

if (process.env.NODE_ENV === 'development') {
  import('react-axe').then(axe => {
    axe.default(React, ReactDOM, 1000);
  });
}

axios.interceptors.request.use(
  config => {
    const newConfig = config;
    if (window.localStorage['okta-token-storage']) {
      const json = JSON.parse(window.localStorage['okta-token-storage']);
      if (json.accessToken) {
        newConfig.headers.Authorization = `Bearer ${json.accessToken.accessToken}`;
      }
    }
    return newConfig;
  },
  error => {
    Promise.reject(error);
  }
);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
