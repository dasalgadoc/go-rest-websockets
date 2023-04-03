# 📑🐹 Handle with JWT in Golang 


Several Endpoints in this repository use JWT authentication. This example probably does not have a clean architecture as I tried to achieve in my previous repositories.
Yet, in order to contribute to new projects with a clean architecture, this brief post outlines the precise methods to handle JWT in Golang.

## 🛠️ Dependencies

[__Golang-jwt__](https://github.com/golang-jwt/jwt)
```bash
go get github.com/golang-jwt/jwt
```

## ✈️ Process 

### 1. Create claims struct
(Example: _claims.go_)
```go
type YourClaims struct {
    // Some custom business elements
    jwt.StandardClaims
}
```
### 2. Generate tokens
(Example: _http_user_login.go_) 

- Create variable with type _YourClaims__ with business information and some Claims implementations 
```go
claims := YourClaims{
    // Some custom business elements
    // Generate expirations
    StandardClaims: jwt.StandardClaims{
        ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
    },
}
```
- Generate the token using a signature and the claims variable created (Example: _)
```go
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
```
- Finally, you can return the token as String using a secret key to cipher its content
```go
stringToken, err := token.SignedString([]byte(<SOME SECRET STRING>))
```

### 3. Ask for the token

Even if certain endpoints of the API don't require it, a validation using JWT is often desired to be conducted at every endpoint.

In several frameworks that allow the creation of middleware, but also it's possible to handle with a dependency (remember DIP), a global function (🫠) or a Chain of Responsibility (remember Command + Mediator + Request/Response + Chain of Responsibility).

Whichever method is used to verify the token, the fundamental processes are as follows.

(Example: _auth.go_) 

1. Get the raw token and clean it
```go
tokenString := strings.TrimSpace(r.Header.Get("x-auth-token"))
```
2. Parse the token with claims
```go
_, err := jwt.ParseWithClaims(tokenString, &<YOUR CLAIMS STRUCT>, func(token *jwt.Token) (interface{}, error) {
   return []byte(<SOME SECRET STRING>), nil
})
```

(Example: _get_claims.go_)

```go
tokenString := strings.TrimSpace(r.Header.Get("x-auth-token"))

token, err := jwt.ParseWithClaims(tokenString, &<YOUR CLAIMS STRUCT>, func(token *jwt.Token) (interface{}, error) {
    return []byte(<SOME SECRET STRING>), nil
})

if err != nil {
    http.Error(w, err.Error(), http.StatusUnauthorized)
    return nil, nil
}

claims, ok := token.Claims.(*<YOUR CLAIMS STRUCT>)

if !ok && !token.Valid {
    return nil, fmt.Errorf("invalid token")
}

return claims, nil
```

