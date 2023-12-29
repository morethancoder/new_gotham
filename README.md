## GOTHAM Stack Starter App

welcome to the most basic app to get you started with the amazing SSR webstack (GOTHAM).
more in depth examples (session, routing, todo app) found [here](https://github.com/morethancoder/hello_gotham).

### Prerequisites

Ensure the following dependencies are installed before setting up and running the app:

* If you are on **MacOS** and have [homebrew](https://brew.sh) installed, simply run:

```bash
# macOS users with Homebrew
brew install go git mariadb node
```


* For linux or windows users kindly follow offical instructions:

- [Golang](https://golang.org/dl/)
- [Git](https://git-scm.com/)
- [MariaDB](https://mariadb.org/download/)
- [Node.js](https://nodejs.org/)


### Tailwindcss and Templ

```bash
# Bun
curl -fsSL https://bun.sh/install | bash

# Tailwind CSS using Bun
bun install -g tailwindcss@latest

# Templ
go install github.com/a-h/templ/cmd/templ@latest
```

### Usage
* clone this repo:
```bash
git clone https://github.com/morethancoder/new_gotham 
```
* make sure your mariadb server is running.
* create a .env file inside the project directory, and fill database credentials:
```bash
cd new_project && mv dotenv .env
```
* tidy project go dependencies:
```bash
go mod tidy
```
* templ:
```bash
make templ
```
* test:
```bash
make test
```
* run:
```bash
make run
```

* enjoy the best SSR web stack


