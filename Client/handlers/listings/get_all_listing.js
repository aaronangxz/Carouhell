function getAllListing() {
    fetch('https://' + getCurrentEnv() + '.herokuapp.com/get_all_listings', {
            method: 'GET'
        })
        .then(response => response.json())
        .then(result => { /*result.Data*/
            for (const d of result.Data) {
                //console.log(d);  
                displayListing(d);
            }
        })
        .catch(error => { console.log('NO:', JSON.stringify(error)); });
}