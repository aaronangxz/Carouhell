<h1 align = "center"> TIC 2601 </h1>

[![Deployment](https://github.com/aaronangxz/TIC2601/actions/workflows/main.yml/badge.svg)](https://github.com/aaronangxz/TIC2601/actions/workflows/main.yml)

<!-- <p align="center">
<!-- <img src = https://github.com/aaronangxz/TIC2601/actions/workflows/main.yml/badge.svg> -->
<!-- <img alt="GitHub go.mod Go version (subdirectory of monorepo)" src="https://img.shields.io/github/go-mod/go-version/aaronangxz/TIC2601?filename=GoServer%2Fgo.mod&style=plastic">
<img alt="npm" src="https://img.shields.io/npm/v/npm">
<img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/aaronangxz/TIC2601">
<br>
<img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/aaronangxz/TIC2601">
<img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/aaronangxz/TIC2601">
<img alt="GitHub issues" src="https://img.shields.io/github/issues/aaronangxz/TIC2601"> -->

<h2>Master Branch (Live)</h2>

- This branch is deployed on https://tic2601-t11.herokuapp.com, all features are somewhat stable
- Do not merge WIP features here

<h2>Docker Container Update History</h2>

| Version     | ImageID        | Date | Status | PR |
| ----------- | -------------- | ---- | ---| ---|
| 1.0         | 14173bd8f6a2   | 24/09/2021     | Good | |
| 1.1         | 9396ef44918a   | 25/09/2021     | Good | |
| 1.2         | 614d41e318e1   | 26/09/2021     | Good | |
| 1.3         | 3c664a7938cf   | 27/09/2021     | Good | [#33](https://github.com/aaronangxz/TIC2601/pull/33) |
| 1.4         | 8f52518e400f   | 30/09/2021     | Good | |
| 1.5         | 700032e186ac   | 03/10/2021     | FAIL | [#44](https://github.com/aaronangxz/TIC2601/pull/44) |
| 1.6         | d85840c3a753   | 03/10/2021     | Good | [#50](https://github.com/aaronangxz/TIC2601/pull/50) |
| 1.7         | 725c67b4ab0f   | 13/10/2021     | Good | [#52](https://github.com/aaronangxz/TIC2601/pull/52) |
| 1.8         | 7407470db9ce   | 16/10/2021     | Good | [#71](https://github.com/aaronangxz/TIC2601/pull/71) |
| 1.9         | 538070d6030f   | 22/10/2021     | Good | [#84](https://github.com/aaronangxz/TIC2601/pull/84) |
| 2.0         | 46a8b45c3312   | 23/10/2021     | Good | [#100](https://github.com/aaronangxz/TIC2601/pull/100) |
| 2.1         | 538070d6030f   | 24/10/2021     | Good | [#127](https://github.com/aaronangxz/TIC2601/pull/127) |

<h2>Architecture</h2>

<p align="center">
<img src="tic2601-architecture.png" width="1000">
</p>

<h2>Docker Deployment</h2>

1. Write `Dockerfile`
2. To build: `docker build --tag tic2601 .` (single dockerfile) / `docker build -f Dockerfile.server .` (multi dockerfiles)
3. Tag docker image `docker tag <imageid> tic2601:<version>`
4. `docker run tic2601` will run container isolated from network.
5. Use `docker run --publish 8080:8080 tic2601` to expose container to network and port. ([host_port]:[container_port])
6. Deploy to Heroku : https://devcenter.heroku.com/articles/container-registry-and-runtime
7. Login via `heroku container:login`
9. Push image `docker tag <imageid> registry.heroku.com/<app>/<process-type>` , `docker push registry.heroku.com/<app>/<process-type>` app is the name of heroku app, process type is `web` 
10. Release image `heroku container:release web -a tic2601-t11`

<h2>MySQL</h2>

1. Access within docker container:
    - `docker exec -it tic2601-db bash`
    - `mysql -u root -p`
2. Check user permission:
    - `show grants for <username>`
3. Give permission:
    - `GRANT ALL PRIVILEGES ON <Table Name> .* TO 'username'@'%'`
4. Check create query:
    - `SHOW CREATE TABLE <db name>.<table name>`
5. Access remotely:
    - CLEARDB_DATABASE_URL: `mysql://b0bc6fadb8432d:f25c7f6b@us-cdbr-east-04.cleardb.com/heroku_bdc39d4687a85d4?reconnect=true`
    - Set up in MySQL Workbench
    
    <img src="tic2601-db.png">

    - Credentials:
        - Hostname:`us-cdbr-east-04.cleardb.com`
        - Port:`3306`
        - Username:`b0bc6fadb8432d`
        - Password:`f25c7f6b`
        - Default Schema:`heroku_bdc39d4687a85d4`

<h2>Test API endpoints</h2>

Postman Workspace: https://www.postman.com/science-specialist-94927587/workspace/tic2601

<h2>Test Branch</h2>

Refer to https://github.com/aaronangxz/TIC2601/tree/test