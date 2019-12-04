# k8-go-demo
super simple go service running on kubernetes

## what it is
This is a microservice written in go that queries an online stock tracking system, does some data aggregation and then outputs a single line of data in response to requests on http

The microservice is built into a docker image and then this image is run in containers on kubernetes

## prequisites
The environment variable XKEY needs to be set to a valid apikey for the https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT online stock tracking system

For example

    export XKEY=demo

For the docker image push/pull to work a docker login with a valid id must be used
On minikube, the credentials must also be set

    kubectl create secret docker-registry dockersecret --docker-server=https://index.docker.io/v1/ --docker-username=<username> --docker-password=<password> --docker-email=<email>


## go testing and building

To run tests on the go program, use the command

    make test

There are unit tests, they should all pass
To see test coverage in a browser

    make browsercover

## building the docker image

To build the docker image use the command

    make image

## testing on kubernetes

If minikube is set up then the command

    make k8test

should show some sample output from the service
