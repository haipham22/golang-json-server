# Golang Sample

## Directory structure

```
jsonserver
├── cmd
│   ├── api.go
│   ├── root.go
│   ├── sample.go
│   └── ...
├── internal // for apps
│   ├── api
│   │   ├── errors
│   │   ├── schemas
│   │   ├── storages
│   │   ├── transport
│   │   └── ...
│   └── another_packages
│       └── ....
├── pkg
│   ├── config
│   ├── databases
│   ├── errors
│   ├── healthcheck
│   ├── models
│   └── utils
│       ├── string
│       │   ├── hash.go
│       │   └── ...
│       └── ...
├── scripts
│   ├── golangci-lint.sh
│   └── ....
├── vendor
├── main.go
├── README.MD
└── ...
```

## Getting started

### Install Dependencies

From the project root, run:

```shell
go build ./...
go test ./...
go mod tidy
```


### Run dev

```shell
go run main.go api
```

Format swag comments on Go code
```shell
go fmt
```

Import order fix
```shell
go-oif ./...
```

Generate wire injector
```shell
wire ./...
```

### Tools
- [google/wire](https://github.com/google/wire) - Dependency Injection
- [go-oif](https://github.com/heyvito/go-oif) - Import Order Fixer
- [golangci-lint](https://github.com/golangci/golangci-lint) - Linter aggregator
