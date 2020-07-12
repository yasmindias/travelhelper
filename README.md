# Travel Helper #
Given a set of routes with their respective costs this project calculates the route with the smallest cost.

## Run Locally
> ### Requirements
> - [Installed Golang environment](https://golang.org/doc/install)

1. Run `go install github.com/yasmindias/travelhelper` to create the executable file in $GOPATH
2. Run `travelhelper [$FILENAME]` where `$FILENAME` is the complete path to you `.csv` file.
    > In the project there is a csv file in the path `resources/input_routes.csv` that is used for running tests. It can also be used for running the project but bear in mind that the tests may fail after alterations.
3. Insert the route in the format "ORG-DST" and the calculated route will be printed in the terminal.

**Example**

```shell
$ go install github.com/yasmindias/travelhelper             
$ travelhelper resources/input_routes.csv               
  Please enter the route: GRU-CDG
  Best route: GRU - BRC - SCL - ORL - CDG > $40
```

## Run Tests

Before running the tests the file `resources/input_routes.csv` is loaded to memory. The tests have pre-defined result expectations based on the original version of this file, so if changes were some tests may fail.

> Run `go test ./...`
