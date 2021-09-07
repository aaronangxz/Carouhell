import Listings from "../Listing/HomeListing"
import React, {Component } from 'react';
// const Home = () => {
// const [error, setError] = useState(null);
//     const [isLoaded, setIsLoaded] = useState(false);
//     const [users, setUsers] = useState([]);
//     useEffect(() => {
//         fetch("http://localhost:8080/listings/1")
//             .then(res => res.json())
//             .then(
//                 (data) => {
//                     setIsLoaded(true);
//                     setUsers(data);
//                 },
//                 (error) => {
//                     setIsLoaded(true);
//                     setError(error);
//                 }
//             )
//       }, [])
// if (error) {
//         return <div>Error: {error.message}</div>;
//     } else if (!isLoaded) {
//         return <div>Loading...</div>;
//     } else {
//         return (
//             <ul>
//                 {users.map(user => (
//                 <li key={user.id}>
//                     {user.name} 
//                 </li>
//                 ))}
//             </ul>
//         );
//     }
// }
// export default Home;

class Home extends Component {

  state = {
    listings: []
  }

  componentDidMount() {
    fetch('http://localhost:8080/listings')
    .then(res => res.json())
    .then((data) => {
      this.setState({ listings: data.data })
    })
    .then((data) => console.log('This is your data', data))
    .catch(console.log)
  }

  render() {
    return (
      <Listings listings={this.state.listings} />

    )
    
  }
}

export default Home;