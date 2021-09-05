import React from 'react'

const Listings = ({ listings }) => {
  return (
    <div>
      <center><h1>Listings</h1></center>
      {listings.data.map((listing) => (
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">{listing.data.item_id}</h5>
            <h6 class="card-subtitle mb-2 text-muted">{listing.data.item_name}</h6>
            <p class="card-text">{listing.data.item_price}</p>
          </div>
        </div>
      ))}
    </div>
  )
};

export default Listings