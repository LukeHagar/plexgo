# Transient

Whether the subtitle is temporary or downloaded.


## Supported Types

### 

```go
transient := components.CreateTransientBoolean(bool{/* values here */})
```

### Transient2

```go
transient := components.CreateTransientTransient2(components.Transient2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch transient.Type {
	case components.TransientTypeBoolean:
		// transient.Boolean is populated
	case components.TransientTypeTransient2:
		// transient.Transient2 is populated
}
```
