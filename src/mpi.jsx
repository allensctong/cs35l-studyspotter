import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './mainpage.jsx';
import './main.css'; // This can be your global CSS file

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
