# yaORM
Yet Another Object–Relational Mapping, is a ORM using Go 1.18

## Summary
1. [How To Use](#how-to-use)
  a. [Defining Models](#defining-models)

## How To Use
### Defining Models

You can declare a simple model using the following struct:
```go
import (
  "github.com/fiuskylab/yaorm/model"
)
type User struct {
  model.Model
  Name  string `json:"name"`
  Email string `json:"email"`
}
```

The `model.Model` is:
```go
// Model is the base for yaORM
type Model struct {
	ID        int       `json:"id" yaorm:"pk,autoincrement:true"`
	CreatedAt time.Time `json:"created_at" yaorm:"trigger:create"`
	UpdatedAt time.Time `json:"updated_at" yaorm:"trigger:update"`
	DeletedAt time.Time `json:"deleted_at" yaorm:"trigger:delete"`
}
```
