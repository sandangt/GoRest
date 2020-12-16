# INFO
## Configuration:
- url: localhost:9062
- exercise url: "localhost:9062/users/:identityID/line_items"
## Init setup:
- Enable variable GO111MODULE through bash script:
```bash
export GO111MODULE=on
```
- Project's location: $GOPATH/src/**GoRest**
## Build and run:
```bash
cd $GOPATH/src/GoRest/
go build
./GoRest
```
## Packages that have unit tests:
- Repositories
- Services
## Test:
- Repositories:
```bash
cd $GOPATH/src/GoRest/Repositories
go test
```
- Services
```bash
cd $GOPATH/src/GoRest/Services
go test
```
