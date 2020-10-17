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
