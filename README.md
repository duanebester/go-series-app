# Go Streaming Series

This repository contains the code for the Go Streaming Series on YouTube/Twitch.

## Getting Started
To build & run the API, you can run the following command:
```bash
go build && ./go-series-app run-api
```

Alternatively, you can run the following command to run the report:
```bash
go build && ./go-series-app run-report
```

### Videos
<details>
  <summary>Building a Go app with Cobra and Fiber</summary>

  [Video](https://youtu.be/g1fl41OewQA)

  * Added basic fiber api
  * Added cobra cli
  * Split up services to be used by the api or cli
  * Services is mockable via an interface
</details>

#### Notes
    * Need to look into ENV variables
        * https://github.com/sethvargo/go-envconfig
        * https://github.com/spf13/viper
