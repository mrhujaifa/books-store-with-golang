book-shop/
├── cmd/
│ └── api/
│ └── main.go
│
├── internal/
│ ├── app/
│ │ ├── app.go
│ │ └── routes.go
│ │
│ ├── config/
│ │ └── config.go
│ │
│ ├── platform/
│ │ ├── db/
│ │ │ └── postgres.go
│ │ ├── logger/
│ │ │ └── logger.go
│ │ └── server/
│ │ └── http.go
│ │
│ ├── middleware/
│ │ ├── auth.go
│ │ ├── cors.go
│ │ ├── logging.go
│ │ └── recovery.go
│ │
│ ├── shared/
│ │ ├── response/
│ │ │ └── json.go
│ │ ├── errors/
│ │ │ └── errors.go
│ │ └── validator/
│ │ └── validator.go
│ │
│ ├── health/
│ │ ├── module.go
│ │ └── handler.go
│ │
│ ├── book/
│ │ ├── module.go
│ │ ├── handler.go
│ │ ├── service.go
│ │ ├── repository.go
│ │ ├── model.go
│ │ └── dto.go
│ │
│ ├── order/
│ │ ├── module.go
│ │ ├── handler.go
│ │ ├── service.go
│ │ ├── repository.go
│ │ ├── model.go
│ │ └── dto.go
│ │
│ ├── user/
│ │ ├── module.go
│ │ ├── handler.go
│ │ ├── service.go
│ │ ├── repository.go
│ │ ├── model.go
│ │ └── dto.go
│ │
│ └── auth/
│ ├── module.go
│ ├── handler.go
│ ├── service.go
│ ├── repository.go
│ ├── model.go
│ └── dto.go
│
├── migrations/
├── docs/
├── .env
├── go.mod
└── go.sum
