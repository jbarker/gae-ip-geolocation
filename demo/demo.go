package main

import "flag"
import "fmt"
import "net/http"
import "strconv"

const (
	defaultPort string = "8777"
	maxPort int64 = 49152
	minPort int64 = 1024
)

func main() {
	var port string = getPortAsString()

    http.Handle("/", http.FileServer(http.Dir("static")))
    
    fmt.Println("Running at: http://localhost:" + port + "/")
    fmt.Println("To exit: CTRL+C")
    
    http.ListenAndServe(":" + port, nil)
}

func getPortAsString() string {
	var portPtr *string
	var portAsInt int64
	var portAsString string
	var err error
	
	portPtr = flag.String("p", defaultPort, "a port number")
	flag.Parse()
	
	portAsInt, err = strconv.ParseInt(*portPtr, 10, 0)
	if (nil != err) {
		fmt.Println("Error:", err)
		portAsString = defaultPort
	} else if (maxPort < portAsInt || portAsInt < minPort) {
		fmt.Println("Error: port must be between",
			strconv.FormatInt(minPort, 10),
			"and",
			strconv.FormatInt(maxPort, 10))
		fmt.Println("Using default port:",
			defaultPort)
		portAsString = defaultPort
	} else {
		portAsString = strconv.FormatInt(portAsInt, 10)
	}
	
	return portAsString
}
