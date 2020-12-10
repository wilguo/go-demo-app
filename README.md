# Getting Started Instructions

1. Download the following snapshots from GitFarm into your `GOPATH` as separate projects. 
    * https://code.amazon.com/packages/Go3p-Github-OpenTelemetry-Go/trees/mainline
    * https://code.amazon.com/packages/Go3p-Github-OpenTelemetry-Go-Contrib/trees/mainline

    Example file path on your machine to the 2 snapshots
    ```
    ~/go/github.com/open-telemetry/opentelemetry-go
    ~/go/github.com/open-telemetry/opentelemetry-go-contrib
    ```

2. Clone this repo into your `GOPATH`

    Example file path on your machine to this repo
    ```
    ~/go/github.com/wilguo/go-demo-app
    ```
   
3. Open the `go.mod` file in this project and replace lines 6 and 7 to the local path of the projects on your local machine

4. Run `go build`

5. Run `./go-demo-app`

6. Visit `localhost:8080/hello-world`