# MediaContainerWithDecisionPeak

Track peak amplitude. Plex may return a numeric value as a number or numeric string, or the sentinel string "-inf".


## Supported Types

### 

```go
mediaContainerWithDecisionPeak := components.CreateMediaContainerWithDecisionPeakNumber(float64{/* values here */})
```

### 

```go
mediaContainerWithDecisionPeak := components.CreateMediaContainerWithDecisionPeakStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mediaContainerWithDecisionPeak.Type {
	case components.MediaContainerWithDecisionPeakTypeNumber:
		// mediaContainerWithDecisionPeak.Number is populated
	case components.MediaContainerWithDecisionPeakTypeStr:
		// mediaContainerWithDecisionPeak.Str is populated
}
```
