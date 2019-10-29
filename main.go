package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"path/filepath"
	"strconv"
)

// Slide type
type Slide struct {
	Heading    string
	Subheading string
	Keywords   []string
	Images     []string
	Videos     []string
	Points     []string
}

// Presentation type
type Presentation struct {
	Title  string
	Slides [] Slide
}

// Global variables
var (
	port                 int
	presentationFilePath string
	mediafilesPath       string
)

// Initialise godazo
func init() {
	flags()
}

// Flags to check on the commandline
func flags() {
	// Port
	flag.IntVar(&port, "port", 8080, "Port to serve godazo")

	// Presentation file
	flag.StringVar(&presentationFilePath, "presentation", filepathFromGodazoPackage("/presentation/example.json"),
		"Json file with slides")

	// Media files
	flag.StringVar(&mediafilesPath, "media", filepathFromGodazoPackage("/media"), "A directory with media files "+
		"(photos and videos) for the presentation")
}

func filepathFromGodazoPackage(filename string) string {
	// Path to godazo
	godazoProjectPath, errorGodazoProjectPath := filepath.Abs(os.Getenv("GOPATH") + "/src/github.com/rutaihwa/godazo/")

	if errorGodazoProjectPath != nil {
		fmt.Printf("// Error while looking for godazo project path")
		fmt.Printf("File error: %v\n", errorGodazoProjectPath)
		os.Exit(1)
	}

	return filepath.Join(godazoProjectPath, filename)
}

func main() {
	// Parse user inputs
	flag.Parse()

	// Read presentation
	file, errorReadingFile := ioutil.ReadFile(presentationFilePath)

	if errorReadingFile != nil {
		fmt.Printf("// Error while reading file %s\n", presentationFilePath)
		fmt.Printf("File error: %v\n", errorReadingFile)
		os.Exit(1)
	}

	var workingPresentation Presentation

	errorUnmarshalingPresentation := json.Unmarshal(file, &workingPresentation)

	if errorUnmarshalingPresentation != nil {
		fmt.Println("error:", errorUnmarshalingPresentation)
		os.Exit(1)
	}

	// Make Gin be on release mode
	gin.SetMode(gin.ReleaseMode)

	// Web server router
	router := gin.Default()

	// Testing working server
	// A super simple route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Templates
	router.LoadHTMLGlob(filepathFromGodazoPackage("templates/*"))

	// Static and media files
	router.Static("/public", filepathFromGodazoPackage("/public"))
	router.Static("/media", mediafilesPath)

	// Route for presentation
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":  workingPresentation.Title,
			"slides": workingPresentation.Slides,
		})
	})

	// Where will godazo be running at
	fmt.Println("godazo will start running at http://localhost:" + strconv.Itoa(port))

	// Running the server
	router.Run(":" + strconv.Itoa(port))
}
