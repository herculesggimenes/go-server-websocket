package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr","localhost:8999", "Http Service Address")

var upgrader = websocket.Upgrader()
