echo "Deploying changes..."
# Pull changes from the live branch
git pull

# Build the image with the new changes
docker-compose build

docker push registry.heroku.com/tic2601-t11/web

# Start the new containers
heroku container:release web -a tic2601-t11
echo "Deployed!"