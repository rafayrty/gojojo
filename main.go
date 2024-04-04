package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {

    listener, err := net.Listen("tcp", "localhost:3000")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    fmt.Println("HTTP Server started at localhost:3000")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())

       // Create a buffer to read data into
       buffer := make([]byte, 1024)

        n, err := conn.Read(buffer)
        if err != nil {
         fmt.Println("Error:", err)
         return
        }

        // Process and use the data (here, we'll just print it)
        file_name := strings.Replace(strings.Split(string(buffer[:n])," ")[1],"/","",-1)

        if file_name != "" {
            file,err := os.ReadFile(file_name)
            content := string(file)
            content_length := len(content)
            //Incase The File does not exist
            if err != nil {
              response := "HTTP/1.1 404 NOT FOUND\r\nContent-Length: 19\r\nServer: Golang(jojo) \r\nContent-Type: text/html; charset=UTF-8\r\n\r\n<h1>Not Found</h1>"

              fmt.Println("Page not found",err)

              _, err = conn.Write([]byte(response))
              if err != nil {
                log.Println("Error writing response:", err)
              }

            }else{
              // Send the HTTP response
             response := "HTTP/1.1 200 OK\r\nContent-Length: "+strconv.Itoa(content_length)+"\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n"+content
              _, err = conn.Write([]byte(response))
              if err != nil {
                  log.Println("Error writing response:", err)
              }
            }
        } else {
            file,err := os.ReadFile("index.html")
            content := string(file)
            content_length := len(content)
            // Send the HTTP response
           response := "HTTP/1.1 200 OK\r\nServer: Golang(jojo)\r\nContent-Length: "+strconv.Itoa(content_length)+"\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n"+content
            _, err = conn.Write([]byte(response))
            if err != nil {
                log.Println("Error writing response:", err)
            }
        }
       conn.Close()
    }
}
