# go_validate

Валидатор на базе `go-playground/validator` с русской локализацией ошибок и кастомными правилами.

## Установка

```
go get github.com/shinpi-tech/go_validate
```

## Использование

```go
valid := validate.NewValid()

if err := valid.Validate(entity); err != nil {
    return c.Status(400).JSON(fiber.Map{"error": err.Error()})
}
```

Кастомные правила: `only_num` (только цифры), `slug` (`^[a-z0-9-_.]+$`).
