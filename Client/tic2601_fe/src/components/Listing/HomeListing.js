import {React} from 'react'
import '../../App.css';

const windowWidth = window.innerWidth

const Listings = ({ listings }) => {
  return (
    <div>
      <center><h1>Popular</h1></center>
      {listings.map((listing) => (
        <div class="card" style= {{width: windowWidth/2}}>
          <img class="card-img-top" src = {listing.item_img} alt=""></img>
          <div class="card-body">
          <h5 class="card-title">{listing.item_name}</h5>
            <p class="card-text">${listing.item_price}.00</p>
            <a href= {listing.item_img} class="card-link">Details</a>
          </div>
        </div>
      ))}
    </div>
  )
};

export default Listings