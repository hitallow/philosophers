# GoLang Philosophers phrases
Philosophers phrases is a little test appliction with Golang and fiber.

# Setup

You must follow these steps

```
  $ go mod init github.com/hitallow/philosophers 
  $ go get github.com/gofiber/fiber/v2
  $ go run server.go
```
# Routes
This small application has four routes, they are:

| PATH                                                              | Description                                      |
| ----------------------------------------------------------------- | ------------------------------------------------ |
| http://localhost:8000                                             | Test if app is running                           |
| http://localhost:8000/api/philosophers                            | Return a list of saved philosophers              |
| http://localhost:8000/api/philosophers/{philosopher_name}/random  | Get a random phrase of philosopher_name repassed |
| http://localhost:8000/api/philosophers/{philosopher_name}/phrases | Get all phrases of philosopher_name              |




