# **FoodChain**

A simple Golang backend application that demonstrates the use of JWT tokens for users.


## **Used Package**

| Name  | Type |
| ------------- | ------------- |
| [gofiber/fiber](https://github.com/gofiber/fiber) | Core  |
| [gofiber/jwt](https://github.com/gofiber/jwt)  | Middleware |
| [arsmn/fiber-swagger](https://github.com/arsmn/fiber-swagger)| Middleware |
| [stretchr/testify](https://github.com/stretchr/testify) | Tests |
| [golang-jwt/jwt](https://github.com/golang-jwt/jwt)| Auth |
| [mongodb/mongo-go-dirver](https://github.com/mongodb/mongo-go-driver) | Database |
| [patrickmn/go-cache](https://github.com/patrickmn/go-cache) | Cache |
| [go-playground/validator](https://github.com/go-playground/validator) | Validation |
| [uber-go/zap](https://github.com/uber-go/zap) | Logger |
| [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Hashing |

## **Template Structure**

### `/cmd`
Main applications for this project.

### `/docs`
Folder with API Documentation. This directory contains config files for auto-generated API Docs by Swagger.
### `/pkg`
Folder with project-specific functionality. This directory contains all the project-specific code tailored only for your business use case.

## **API Endpoint**
| URL  | Port |HTTP Method | Operation
| --------|------------|--------------- |--------|
| v1/signup | 3000  | Post| User sign up to receive tokens |
| v1/login | 3000  | Post| Authenticates a user and generates a JWT |
| v1/refresh | 3000  | Post| Generates a refresh token for a user |
| v1/user | 3000  | Get| Returns a list of the users |
| v1/user/:id | 3000  | Get| Returns the detailed information of an user |
| v1/user | 3000  | Post| Creates a new user |
| v1/user/:id | 3000  | Put| Updates an existing user |
| v1/user/:id | 3000  | Delete| Deletes an user |