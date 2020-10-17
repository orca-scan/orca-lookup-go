# orca-lookup-go

This open source project is a an example of [how to scan barcodes using a smartphone](https://orcascan.com/mobile) and [present data from your system](https://orcascan.com/docs/api/lookup-url) using [Go](https://golang.org/) and the HTTP web framework [gin](https://github.com/gin-gonic/gin).

**How it works:**

1. A user [scans a barcode](https://orcascan.com/mobile) using their smartphone
2. Orca Scan sends a HTTP GET request to your endpoint with `?barcode=value`
3. Your system queries a database or internal API for a `barcode` match
4. Your system returns the data in JSON format with keys matching column names
5. The [Orca Scan mobile](https://orcascan.com/mobile) app presents that data to the user

*If the mobile user has [update permission](https://orcascan.com/docs/getting-started/adding-users#selecting-user-permissions) and saves the data, it will saved to your Orca sheet.*

## Install

First ensure you have [Go](https://golang.org/dl/) installed:

```bash
# should return 1.13 or higher
go version
```

Then execute the following:

```bash
# download this example code
git clone https://github.com/orca-scan/orca-lookup-go.git

# go into the new directory
cd orca-lookup-go

# install dependencies
go get -d ./...
```

## Run

```bash
# start the project
go run main.go
```

Visit [http://localhost:5000?barcode=4S3BMHB68B3286050](http://localhost:5000?barcode=4S3BMHB68B3286050) to see the following:

```json
{
    "VIN": "4S3BMHB68B3286050",
    "Make": "SUBARU",
    "Model": "Legacy",
    "Manufacturer Name": "FUJI HEAVY INDUSTRIES U.S.A",
    "Vehicle Type": "PASSENGER CAR",
    "Year": 1992
}
```

## How this example works

This [simple example](main.go) uses the [gin](https://github.com/gin-gonic/gin) framework:

```go
package main

// pull in the GIN framework
import "github.com/gin-gonic/gin"

// main go entry point
func main() {

	// create a GIN router to handle requests
	router := gin.Default()

	// GET / handler
	router.GET("/", func(c *gin.Context) {

		// get the incoming barcode sent from Orca Scan (scanned by a user)
		barcode := c.Query("barcode")

		// TODO: query a database or API to retrieve some data based on barcode value

		// return JSON object (property names must match Orca column names)
		c.JSON(200, gin.H{
			"VIN":               barcode,
			"Make":              "SUBARU",
			"Model":             "Legacy",
			"Manufacturer Name": "FUJI HEAVY INDUSTRIES U.S.A",
			"Vehicle Type":      "PASSENGER CAR",
			"Year":              1992,
		})
	})

	// start the webserver on port 5000
	router.Run(":5000")
}
```

## Troubleshooting

If you run into any issues not listed here, please [open a ticket](https://github.com/orca-scan/orca-lookup-go/issues).

## Contributing

To contribute simply:

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request

## History

For change-log, check [releases](https://github.com/orca-scan/orca-lookup-go/releases).

## License

Licensed under [MIT License](LICENSE) &copy; [Orca Scan](https://orcascan.com)