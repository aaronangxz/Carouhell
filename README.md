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

<h2>Getting Started</h2>

<h3>Environment</h3>

1. Install [Node.js](https://nodejs.org/en/) and npm for front end

2. Install [Go](https://golang.org/doc/install) for back end

3. Install [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) for version control

4. Set up [SSH](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh) or [PAT](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) for authentication

4. Clone this repository to your local machine<br>

- SSH: `git clone git@github.com:aaronangxz/TIC2601.git`

- HTTPS: `git clone https://github.com/aaronangxz/TIC2601.git`

<h3>JS Front End</h3>

1. The front end directory is `/Client/tic2601_fe/`

2. run `npm install` to install all dependencies.

3. To run project locally, run `npm start`.

<h3>Golang Back End</h3>

1. The back end directory is `/GoServer/`

2. run `go mod download` and `go mod vendor` to install all dependencies.

3. To run project locally, run `go run main.go`.

<h2>Code Contributions</h2>

1. Always `git pull` first to retrieve latest changes of master branch.

2. Create a new git branch with your name and description: "XuanZe/Do_something"<br>
`git branch XuanZe/Do_something`<br>
`git checkout XuanZe/Do_something`

3. Start writing your code. `git rebase master` to fetch changes from master branch into your branch.

4. Once done, <br>
`git add .`<br>
`git commit -m "<commit message>"`<br>
`git push -u origin XuanZe/Do_something`

5. Create Pull Request, we can merge to master after reviewing.