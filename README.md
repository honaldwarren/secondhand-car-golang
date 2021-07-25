# secondhand-car-golang

Welcome! This is an example of a simple REST API implementation using clean architecture (onion layer) written in Go language complete with Dependency Injection, following SOLID principles.

It has the following dependencies:

 - [Chi (Router)](https://github.com/go-chi/chi)
 - [Gorilla/Schema (Form to Struct)](https://github.com/gorilla/schema)
 - [GoJSONQ (ODM)](https://github.com/thedevsaddam/gojsonq)


**Installation**
-------

Install Go Locally

    https://golang.org/doc/install

Clone the Source

    git clone https://github.com/honaldwarren/secondhand-car-golang

Setup Dependencies

    go get -u github.com/go-chi/chi
    go get -u github.com/gorilla/schema
    go get -u github.com/thedevsaddam/gojsonq
    
Run the App

    Linux
    go build
    ./service-pattern-go
    
    Windows
    go build
    .\secondhand-car-golang.exe


Visit

    http://localhost:9000/api/car/filter


	


**Apply Filter**
-------

Parameters Sample

    http://localhost:9000/api/car/filter?Make=Honda,Toyota&Model=Civic,Camry&MinYear=2000&MaxYear=2020&MinPrice=15000&MaxPrice=20000

    Make: Honda,Toyota
    Model: Civic,Camry
    MinYear: 2000
    MaxYear: 2020
    MinPrice: 15000
    MaxPrice: 20000


Using Postman

![Imgur Image](https://imgur.com/KZwUu2H.jpg)