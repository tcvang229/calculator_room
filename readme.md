# Welcome to Calculator Room

Calculator Room is a quick and easy web app that'll solve basic expressions
using Math.JS. It'll spin up a HTTP server
using GoLang for the back-end and HTML, CSS, and JavaScript for the front. The
server also establishes a websocket connection for each client that connects
to the server. Specific frameworks and packages will be listed below.

Calculator Room allows communication between clients and servers utilizing
websockets for real-time updates. Onload, it'll use REST API to update the history of expressions by accessing the returned JSON data. 
The messages sent by the client(s) are stored in a JSON file and can be
accessed by the REST API serviced by the Go server.
## Installaton

Transfer repository to local machine:
* Git URL: `https://github.com/tcvang229/calculator_room.git`

        git clone https://github.com/tcvang229/calculator_room.git

* or download ZIP from up above

## Compile and Run 

Compile .go files in the `/src/` directory using the command line. 

        go build -o server    

        or      

        go run server.go 

Use command, `go help` to learn more about how to use the go command

Then start the executable file 
* Linux/Mac OS: 
        
        ./server
* Windows: 

        server.exe

A local HTTP server will be created.

## How To Use

1. Access the application at this URL: https://sleepy-lowlands-26389.herokuapp.com/.
2. Multiple clients (creating new tabs containing the URl) can be created and tested.
3. Enter a valid expression; **Do NOT use 'x' for multiplication, use '*' instead**.

* You can review Math.JS library to find out what is a valid expression and other features it offers.

![client example](https://raw.githubusercontent.com/tcvang229/calculator_room/master/other/client.png)

## Framework and Packages

Calculator Room uses a few Go packages and JavaScript libraries. 

Front-End: 
* Math.js v7.2.0
* jQuery v3.5.1
* Bootstrap v3.4.1

Back-End:
* github.com/gorilla/websocket
* encoding/json
* log
* net/HTTP

## Support

The only dependancy that may require some manual installation is the 
`github.com/gorilla/websocket` GoLang package. If GoLang compiler is asking
to install the package, run this command to install it:    

        go get github.com/gorilla/websocket

## Relevant Links

* https://github.com/gorilla/websocket
* https://mathjs.org/
* https://code.jquery.com/
* https://getbootstrap.com/docs/4.5/getting-started/introduction/ 
