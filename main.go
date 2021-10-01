//This implementation is done with help from the golang guide: https://golang.org/doc/tutorial/web-service-gin
package DISYS_ME1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Couse representing data about the course
type course struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Teacher string `json:"teacher"`
}

var courses = []course{
	{ID: "BSANDSA1KU", Title: "Analysis, Design and Software Architecture", Teacher: "Paolo Tell"},
	{ID: "BSDISYS1KU", Title: "Distributed Systems", Teacher: "Marco Carbone"},
	{ID: "BSINDBS1KU", Title: "Introduction to Database Systems", Teacher: "Björn Thór Jónsson"},
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCouses)

	router.Run("localhost:8080")
}

// getCourses responds with the list of all courses as JSON.
func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

// postCouses adds an course from JSON received in the request body.
func postCouses(c *gin.Context) {
	var newCourse course

	// Call BindJSON to bind the received JSON to newCourse.
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// Add the new course to the slice.
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

// getCourseByID locates the course whose ID value matches the id parameter sent by the client, then returns that course as a response.
func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of Courses, looking for an course whose ID value matches the parameter.
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}
