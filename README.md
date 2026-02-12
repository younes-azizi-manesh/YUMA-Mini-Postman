# Go Postman 🛠️

![Go Postman Screenshot](screenshot.png)

**Go Postman** is a lightweight Postman-like tool built in Go for sending and testing HTTP requests efficiently.

---

## 🚀 Getting Started

To run the project, first clone the repository:

```bash
git clone https://github.com/younes-azizi-manesh/YUMA-Mini-Postman.git
cd YUMA-Mini-Postman
```
Then start the server using:

```bash
go run ./cmd/server/main.go
```

After the server is running, open your browser and go to:
```bash
http://localhost:8080
```

🛠️ How to Use

This tool allows you to test various HTTP requests. All HTTP methods are supported. Each request has four main fields:

Method: HTTP method (GET, POST, PUT, DELETE, etc.)

URL: Endpoint URL

Headers: Request headers

Body: Request body

Simply fill in the fields and send the request to see the response.

💻 Technologies

This project is built with Go for handling HTTP requests and uses HTML, CSS, and JavaScript for the user interface.

📝 Project Story

I had been spending some time exploring Go out of curiosity, and it was really interesting for me when I didn’t have internet access. I needed to test an API and had to use curl because I didn’t have Postman. That’s when it suddenly hit me—why not build something myself! I decided to use Go for this, and the result was this small and cute little tool.

⚡ Features

Supports all HTTP methods

Simple and lightweight user interface

Quick testing of endpoints

Works offline (no internet required)