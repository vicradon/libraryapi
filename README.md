# libraryapi
A simple library API built with Gin

## Prrequisites
You need Go installed locally to run this project. 

## Setting Up
You can set up this API by first cloning the project, building it, then running the built file. Here are steps.

### Step 1: Clone thr project and change directory into it

```sh
git clone https://github.com/vicradon/libraryapi.git
cd libraryapi
```

### Step 2: Create the .env file for the project

```sh
cp .env.example .env
```

### Step 3: Build and run the project

```sh
go build -o libraryapi
./libraryapi
```

## Running Tests

You can run the included tests using the command below:

```sh
go test -v ./tests/*
```

## API Docs

You can view thr docs for this API [here](https://library.osinachi.me/swagger/index.html)
