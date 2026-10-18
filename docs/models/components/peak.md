# Peak

Track peak amplitude. Plex may return a numeric value as a number or numeric string, or the sentinel string "-inf".


## Supported Types

### 

```go
peak := components.CreatePeakNumber(float64{/* values here */})
```

### 

```go
peak := components.CreatePeakStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch peak.Type {
	case components.PeakTypeNumber:
		// peak.Number is populated
	case components.PeakTypeStr:
		// peak.Str is populated
}
```
