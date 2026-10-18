# Loudness

Integrated loudness in LUFS. Plex may return a numeric value as a number or numeric string, or the sentinel string "-inf".


## Supported Types

### 

```go
loudness := components.CreateLoudnessNumber(float64{/* values here */})
```

### 

```go
loudness := components.CreateLoudnessStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch loudness.Type {
	case components.LoudnessTypeNumber:
		// loudness.Number is populated
	case components.LoudnessTypeStr:
		// loudness.Str is populated
}
```
