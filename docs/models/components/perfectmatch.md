# PerfectMatch

Whether the subtitle is an exact match.


## Supported Types

### 

```go
perfectMatch := components.CreatePerfectMatchBoolean(bool{/* values here */})
```

### PerfectMatch2

```go
perfectMatch := components.CreatePerfectMatchPerfectMatch2(components.PerfectMatch2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch perfectMatch.Type {
	case components.PerfectMatchTypeBoolean:
		// perfectMatch.Boolean is populated
	case components.PerfectMatchTypePerfectMatch2:
		// perfectMatch.PerfectMatch2 is populated
}
```
