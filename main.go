// main.go
package main

import (
	"fmt"
	"go-detect_service/config"
	"go-detect_service/internal/database"
	routes "go-detect_service/internal/routes/v1"
	"runtime"
)

func printHello() {
	fmt.Println("#-----------------------------------------------------#")
	fmt.Println("# start go-detect_service with", runtime.Version())
	fmt.Println("#")
	fmt.Println("# *****  *****  *        *    *   *  ***** ")
	fmt.Println("# *      *   *  *       * *   **  *  *     ")
	fmt.Println("# *  **  *   *  *      *****  * * *  *  ** ")
	fmt.Println("# *   *  *   *  *      *   *  *  **  *   * ")
	fmt.Println("# *****  *****  *****  *   *  *   *  ***** ")
	fmt.Println("#")
	fmt.Println("#-----------------------------------------------------#")
	fmt.Println()
}

func registerServer() {
	config.Init()
	printHello()
	database.InitMysql()
	routes.InitRouter()
}

func main() {
	registerServer()
}
