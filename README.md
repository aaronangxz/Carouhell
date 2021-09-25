<h1 align = "center"> TIC 2601 </h1>
<!-- <p align="center">
<img alt="GitHub go.mod Go version (subdirectory of monorepo)" src="https://img.shields.io/github/go-mod/go-version/aaronangxz/TIC2601?filename=GoServer%2Fgo.mod&style=plastic">
<img alt="npm" src="https://img.shields.io/npm/v/npm">
<img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/aaronangxz/TIC2601">
<br>
<img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/aaronangxz/TIC2601">
<img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/aaronangxz/TIC2601">
<img alt="GitHub issues" src="https://img.shields.io/github/issues/aaronangxz/TIC2601">
</p> -->

<h2>Master Branch (Live)</h2>

- This branch is deployed on https://tic2601-t11.herokuapp.com/test, all features are somewhat stable
- Do not merge WIP features here

<h2>Docker Container Update History</h2>

| Version     | ImageID        | Date |
| ----------- | -------------- | ---- |
| 1.0         | 14173bd8f6a2   | 24/09/2021     |
| 1.1         | 9396ef44918a   | 25/09/2021     |

<h2>Architecture</h2>

<p align="center">
<img src="tic2601-architecture.png" width="1000">
</p>

<h2>Docker Deployment</h2>

1. Write `Dockerfile`
2. To build: `docker build --tag tic2601 .`
3. `docker run tic2601` will run container isolated from network.
3. Use `docker run --publish 8080:8080 tic2601` to expose container to network and port. ([host_port]:[container_port])
3. Use `docker run --publish 8080:8080 tic2601` to expose container to network and port. ([host_port]:[container_port])
4. Deploy to Heroku : https://devcenter.heroku.com/articles/container-registry-and-runtime
5. Push image `docker tag <image> registry.heroku.com/<app>/<process-type>` , `docker push registry.heroku.com/<app>/<process-type>` app is the name of heroku app, process type is `web` 
6. Release image `heroku container:release web -a tic2601-t11`

<h2>Getting Started</h2>

Refer to https://github.com/aaronangxz/TIC2601/tree/test