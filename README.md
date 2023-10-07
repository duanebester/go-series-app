# Go Streaming Series

This repository contains the code for the Go Streaming Series on YouTube/Twitch.

## Getting Started
To build the project, run the following command:
```bash
go build && ./go-series-app run-api
```

Then you can run the api with:
```bash
./go-series-app run-api
```

### [Building a Go app with Cobra and Fiber](https://youtu.be/g1fl41OewQA)
    * Added basic fiber api
    * Added cobra cli
    * Split up services to be used by the api or cli
    * Services is mockable via an interface

#### Notes
    * Need to look into ENV variables
        * https://github.com/sethvargo/go-envconfig
        * https://github.com/spf13/viper
