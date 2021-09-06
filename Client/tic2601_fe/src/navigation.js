import React from 'react';
import { Nav, NavItem} from 'reactstrap';
import { NavLink } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faHome, faFire, faTags, faBell, faUserCircle} from '@fortawesome/free-solid-svg-icons';
import 'bootstrap/dist/css/bootstrap.css';
import './App.css';

const tabs = [{
    route: "/",
    icon: faHome,
    label: "Home",
  },{
    route: "/Popular",
    icon: faFire,
    label: "Popular"
  },{
    route: "/Sell",
    icon: faTags,
    label: "Sell"
  },{
    route: "/Notifications",
    icon: faBell,
    label: "Notifications"
  },{
    route: "/Me",
    icon: faUserCircle,
    label: "Me"
  }]

  const Navigation = (props) => {

    return (
      <div>
        {/*Top Nav Bar*/}
        <nav className="navbar navbar-expand-md navbar-light d-none d-lg-block sticky-top" role="navigation">
          <div className="container-fluid">
              <a className="navbar-brand" href="/home">TIC2601</a>
              <Nav className="ml-auto">
                <NavItem>
                  <NavLink to="/" className="nav-link">
                    Home
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink to="/popular" className="nav-link">
                    Popular
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink to="/sell" className="nav-link">
                    Sell
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink to="/notifications" className="nav-link">
                    Notifications
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink to="/me" className="nav-link">
                    Me
                  </NavLink>
                </NavItem>
              </Nav>
          </div>
        </nav>
        {/*Bottom Nav Bar*/}
        <nav className="navbar fixed-bottom navbar-light bg-light d-block d-lg-none bottom-tab-nav" role="navigation">
        <Nav className="w-100">
          <div className="d-flex flex-row justify-content-around w-100">
            {
              tabs.map((tab, index) =>(
                <NavItem key={`tab-${index}`}>
                  <NavLink to={tab.route} className="nav-link bottom-nav-link " activeClassName="active" exact={true}>
                    <div className="d-flex flex-column justify-content-between align-items-center">
                      <FontAwesomeIcon size= "lg" icon={tab.icon}/>
                      <div className="bottom-tab-label">{tab.label}</div>
                    </div>
                  </NavLink>
                </NavItem>
              ))
            }
          </div>
        </Nav>
      </nav>
      </div>
    )
  };


export default Navigation;