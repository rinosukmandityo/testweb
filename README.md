# Functional Suite

`functional_spec/` contains some logics for each parking lot command and automated testing suite that will validate the correctness of parking lot command for the sample input and output.

We do not support Windows at this point in time. If you don't have access to an OSX or Linux machine, we recommend setting up a Linux machine you can develop against using something like [VirtualBox](https://www.virtualbox.org/) or [Docker](https://docs.docker.com/docker-for-windows/#test-your-installation).

This needs following mandatory apps:
[Golang to be installed](https://golang.org/doc/install), and
[Git to be installed](https://www.atlassian.com/git/tutorials/install-git), followed by some libraries. The steps are listed below.

## Setup

First, put `parking_lot.zip` into your linux environment, after [Golang](https://golang.org/doc/install) installation and [Git](https://www.atlassian.com/git/tutorials/install-git) installation finished. Then run the following commands in your current directory.

```
$ go version # confirm Golang present
go version go1.9.4 linux/amd64
$ git version # confirm Git present
git version 2.17.2
```
Copy all following commands into your terminal to prepare all files inside `bin/` directory, make sure that you are still in the same directory of `parking_lot.zip`:
```
export GOPATH=$(go env GOPATH)  # to set GOPATH as go working directory
sudo mkdir -p $GOPATH/src/github.com/rinosukmandityo/ # create this directory to run our test scenario or parking lot app
sudo unzip parking_lot.zip -d $GOPATH/src/github.com/rinosukmandityo/ # unzip parking_lot.zip into github.com/rinosukmandityo directory
sudo chmod +x $GOPATH/src/github.com/rinosukmandityo/parking_lot/bin/* # convert all files inside bin directory into executable files
cd $GOPATH/src/github.com/rinosukmandityo/parking_lot/ # go to parking_lot directory
```
Run this command to Install Dependencies and Run Test File to Validate
```
parking_lot $ bin/setup
```

## Usage

You can run the full suite from `parking_lot` by doing
```
parking_lot $ bin/run_functional_specs
```

You can run the app by doing
```
parking_lot $ bin/parking_lot
```

Or you can run the app with file by doing
```
parking_lot $ bin/parking_lot functional_spec/fixtures/file_input.txt
```
