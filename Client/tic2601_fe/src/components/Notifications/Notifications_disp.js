import {React} from 'react'
import '../../App.css';

const windowWidth = window.innerWidth

const Notifications = ({ notifications }) => {
  return (
    <div>
      <h1>Notifications</h1>
      {notifications.map((notification) => (
        <div class="card" style= {{width: windowWidth}}>
          <div class="card-body">
          <h5 class="card-title">{notification.notification_id}</h5>
            <p class="card-text">{notification.notification_text}</p>
          </div>
        </div>
      ))}
    </div>
  )
};

export default Notifications