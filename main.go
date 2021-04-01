package main

import (
	"fmt"

	"github.com/moskv08/go-timezone-rocket/models"
)

func main() {

	fmt.Println("UTC:", models.GetSpecificTimeZone("UTC"))
	fmt.Println("MST:", models.GetSpecificTimeZone("MST"))
	fmt.Println("EST:", models.GetSpecificTimeZone("EST"))

	zone, time := models.GetMyTimeZone()
	fmt.Println(zone, ":", time)
}
