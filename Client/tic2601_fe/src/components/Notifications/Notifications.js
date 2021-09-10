import React, {Component } from 'react';
import Notifications from "./Notifications_disp"

class NotificationHome extends Component {

  state = {
    notifications: []
  }

  componentDidMount() {
    fetch('http://localhost:8080/notifications/696969')
    .then(res => res.json())
    .then((data) => {
      this.setState({ notifications: data.data })
    })
    .then((data) => console.log('This is your data', data))
    .catch(console.log)
  }

  render() {
    return (
      <Notifications notifications={this.state.notifications} />
    )
    
  }
}

export default NotificationHome;