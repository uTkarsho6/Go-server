# Basic Go Web Server
A lightweight, high-performance web server built from scratch using **Go** (Golang). This project demonstrates the core principles of backend development, including static file serving, custom routing, and HTTP form processing without the use of external frameworks.

## Features:

* Static File Serving: Automatically serves HTML content from a dedicated directory using http.FileServer.

* Dynamic Routing: Custom handlers for specific endpoints like /hello.

* Form Data Processing: Parses and extracts data from POST requests (Name and Address fields).

* Robust Error Handling: Built-in validation for 404 errors and HTTP method restrictions.

## Tech Stack:
Language: Go (Golang)

Standard Library: net/http (Server logic), fmt (I/O), log (Error logging)

Frontend: HTML5
