# gocloak tokencache

This is a token cache based on [gocloak](https://github.com/Nerzal/gocloak).

This small library automates the refetching of keycloak tokens using gocloak.

## Usage

### Installation
```
go get github.com/escb005/gocloak-tokencache
```

### Importing
```
import tokencache "github.com/escb005/gocloak-tokencache"
```

### Creating a token cache
```go
// configuration variables
ctx := context.Background()
realm := "test-realm"       // The realm of the token
expiresSkew := 30           // How long before expiry the token should be refetched (in seconds)
grantType := "password"     // Specifc options for the token
clientId := "token-cache"
username := "TestUser"
password := "password"

// create the cache
cache := tokencache.NewTokenCache("https://my.keycloak",
    tokencache.WithRealm(realm),
    tokencache.WithExpiresSkew(expiresSkew),
    tokencache.WithTokenOptions(gocloeak.TokenOptions{
        GrantType: &grantType,
        ClientID: &clientId,
        Username: &username,
        Password: &password,
    })
)

// get the token from the cache
token, err := cache.GetToken()
if err != nil {
    panic("Could not fetch token!")
}

// Do something with the token
```