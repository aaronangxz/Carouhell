import React, { Component } from 'react';
import PopularListings from "../Listing/PopularListing"

class Popular extends Component {
  
  state = {
    listings: []
  }
  componentDidMount() {
    fetch('http://localhost:8080/listings')
    .then(res => res.json())
    .then((data) => {
      this.setState({ listings: data.data })
    })
    .then((data) => console.log('This is your data',data))
    .catch(console.log)
  }

  render() {
    const length = this.state.listings.length;
    console.log('This is your data',length)
    return (
      <PopularListings listings={this.state.listings} />
    )
  }
}

export default Popular;