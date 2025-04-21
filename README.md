# Systems-Test2

A simple, progressively enhanced concurrent TCP server written in Go. This project walks through three stages of development — starting from a basic echo server to a feature-rich concurrent server with logging, timeouts, and commands.

Table of Contents
About the Project

Versions

Original: Simple Echo Server

Version 1: Echo Server with Commands

Final Server: Concurrent TCP Server (Final)

Build & Run
go run main.go --port=4000
nc localhost 4000 (In new terminal) 
Test all commands


This is a small project to learn and demonstrate building a TCP server in Go. The server starts simple and evolves step-by-step into a concurrent, command-driven TCP service that handles multiple clients, logs events, and supports interactive commands.

Original: Simple Echo Server
Listens for TCP connections on port 4000

Echoes back any message it receives from a client

Logs client connections and disconnections to the console


 Version 1: Echo Server with Commands
Adds command support:

/time — returns current server time

/quit — disconnects the client

Any other input is echoed back

Logs connections and disconnections to the console

Final: Concurrent TCP Server (Final)
Handles multiple concurrent clients

Tracks number of connected clients

Implements inactivity timeout (30 seconds)

Logs per-client messages into individual log files

Supports commands:

/echo [message]

/time

/date

/joke

/clients

/help

HW#2 Video Link
https://youtu.be/zaUo_s6V-Lsk

/quit or bye

Logs server-wide events to the console
