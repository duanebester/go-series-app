# Go Streaming Series

This repository contains the code for the Go Streaming Series on [YouTube](https://www.youtube.com/playlist?list=PLzIkykhdNahzyImgGb7LmEv15l-2seduU/[Twitch](https://www.twitch.tv/duanebester).

## Getting Started
To build & run the API, you can run the following command:
```bash
go build && ./go-series-app run-api
```

Alternatively, you can run the following command to run the report:
```bash
go build && ./go-series-app run-report
```

For running w/ hot reloading via [air](https://github.com/cosmtrek/air):
```bash
air -c .air.toml run-api
```

### Videos

1. [Building a Go app with Cobra and Fiber](https://youtu.be/g1fl41OewQA)
1. [Adding ENV config and hot reloading to a Go app](https://youtu.be/U9x1V1adQzI)
1. [Adding JWT Authentication to a Go Fiber app](https://youtu.be/EFx3rlXgae4)
1. [Refactoring Go API - Breaking out services](https://youtu.be/RZnTd29Fnr4)
1. [Refactoring Go API - Breaking out repositories]()

<details>
<summary>Video Notes</summary>

10/06/2023
* Added basic fiber api
* Added cobra cli
* Split up services to be used by the api or cli
* Services is mockable via an interface

10/07/2023
* Added viper for env config https://github.com/spf13/viper
* Added air for hot reloading

10/08/2023
* Added JWT auth
* Added middleware for auth
* Refactored API to use handlers and middleware

10/09/2023
* refactored services
* refactored handlers

10/10/2023
* refactored services again
* added repository layer
* changed models to use uuids
* started postman collection

Need to look into:
* dockerizing the _whole_ app
* UUIDs for the db
* authz/authn for API hmac for IoT devices
* passkey based login
* Run seedDB multiple times
</details>
