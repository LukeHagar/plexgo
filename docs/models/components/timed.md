# Timed

Whether lyrics are timestamped.


## Supported Types

### 

```go
timed := components.CreateTimedBoolean(bool{/* values here */})
```

### Timed2

```go
timed := components.CreateTimedTimed2(components.Timed2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch timed.Type {
	case components.TimedTypeBoolean:
		// timed.Boolean is populated
	case components.TimedTypeTimed2:
		// timed.Timed2 is populated
}
```
