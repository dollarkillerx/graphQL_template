# graphQL_template
graphQL_template golang

Chi + 99designs/gqlgen

### tree

``` 
.
├── cmd
│   └── main.go
├── configs
│   ├── config.yaml
│   └── gqlgen.yml
├── gen.go
├── go.mod
├── go.sum
├── internal
│   ├── conf
│   │   └── config.go
│   ├── generated
│   │   ├── generated.go
│   │   ├── models_gen.go
│   │   ├── README.md
│   │   └── resolver.go
│   ├── graphql
│   │   ├── common
│   │   │   └── common.graphql
│   │   └── user
│   │       └── auth.graphql
│   ├── middlewares
│   │   ├── auth.go
│   │   ├── context.go
│   │   └── cors.go
│   ├── pkg
│   │   ├── enum
│   │   │   ├── auth.go
│   │   │   └── middleware.go
│   │   └── models
│   │       └── middleware.go
│   ├── resolvers
│   │   ├── resolver.go
│   │   └── user.go
│   ├── storage
│   │   ├── interface.go
│   │   └── simple
│   │       └── simple.go
│   └── utils
│       ├── base.go
│       ├── jwt.go
│       └── logger.go
├── LICENSE
├── README.md
└── tools.go
```

#### example payload

``` 

query captcha {
  captcha{
    base64Captcha
    base64Captcha
  }
}


mutation login {
  loginByPassword(input: {account:"sad",password:"asd",captchaID: "sad",captchaCode:"sdgerg"}) {
    userID
    accessTokenString
  }
}
```