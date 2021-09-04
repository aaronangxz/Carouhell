import React from 'react';
import './App.css';
// import Login from './components/Login/Login'
// import useToken from './useToken';
import { BrowserRouter, Switch, Route, NavLink} from "react-router-dom";
import Navigation from './navigation';
import Home from './components/Home/Home';
import Popular from './components/Popular/Popular';
import Sell from './components/Sell/Sell';
import Notifications from './components/Notifications/Notifications';
import Me from './components/Me/Me';


// function setToken(userToken) {
//   sessionStorage.setItem('token', JSON.stringify(userToken));
// }

// function getToken() {
//   const tokenString = sessionStorage.getItem('token');
//   const userToken = JSON.parse(tokenString);
//   return userToken?.token
// }

function App() {
  // const { token, setToken } = useToken();

  // if(!token) {
  //   return <Login setToken={setToken} />
  // }

  return (
    <div>
      <BrowserRouter>
      <Navigation />
        <Switch>
          <NavLink to path="/" exact component={Home} exact={true} activeClassName="active"/>
          <Route path="/" exact component={Home} exact={true}/>
          <Route path="/popular" exact component={Popular} exact={true} />
          <Route path="/sell" exact component={Sell} exact={true}/>
          <Route path="/notifications" exact component={Notifications}exact={true}/>
          <Route path="/me" exact component={Me}exact={true} />
        </Switch>
      </BrowserRouter>
    </div>
  );
} 
export default App;