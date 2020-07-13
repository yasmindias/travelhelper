# Travel Helper #
Given a set of routes with their respective costs this project calculates the route with the smallest cost.

## Run Locally
### Requirements
- [Installed Golang environment](https://golang.org/doc/install)

1. Run `go install github.com/yasmindias/travelhelper` to create the executable file in $GOPATH
2. Run `travelhelper [$FILENAME]` where `$FILENAME` is the complete path to the `.csv` file.
    > In the project there is a csv file in the path `resources/input_routes.csv` that is used for running tests. It can also be used for running the project but bear in mind that the tests may fail after alterations.


### Use on Command Line
**Find Best Route**

Insert the route in the format "ORG-DST" and the calculated route will be printed in the terminal.

#### Example

```shell
$ go install github.com/yasmindias/travelhelper             
$ travelhelper resources/input_routes.csv               
  Please enter the route: GRU-CDG
  Best route: GRU - BRC - SCL - ORL - CDG > $40
```


### Use on REST API
The API can be accessed on `localhost:3000`. The export for the postman collection can be found [here](https://github.com/yasmindias/travelhelper/tree/master/postman-collection)

**Add New Route**

`POST /api/routes`

#### Request Example
```json
{
    "Origin": "SCL",
    "Destiny": "BRC",
    "Cost": 19
}
```

**Find Best Route**

`GET /api/bestroute?origin=GRU&destiny=CDG`

#### Response Example
```json
{
    "Cost": 40,
    "Path": [
        "GRU",
        "BRC",
        "SCL",
        "ORL",
        "CDG"
    ]
}
```

## Run Tests

Before running the tests the file `resources/input_routes.csv` is loaded to memory. The tests have pre-defined result expectations based on the original version of this file, so if changes were some tests may fail.

> Run `go test ./...`
