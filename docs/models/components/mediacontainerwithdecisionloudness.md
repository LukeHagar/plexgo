# MediaContainerWithDecisionLoudness

Integrated loudness in LUFS. Plex may return a numeric value as a number or numeric string, or the sentinel string "-inf".


## Supported Types

### 

```go
mediaContainerWithDecisionLoudness := components.CreateMediaContainerWithDecisionLoudnessNumber(float64{/* values here */})
```

### 

```go
mediaContainerWithDecisionLoudness := components.CreateMediaContainerWithDecisionLoudnessStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mediaContainerWithDecisionLoudness.Type {
	case components.MediaContainerWithDecisionLoudnessTypeNumber:
		// mediaContainerWithDecisionLoudness.Number is populated
	case components.MediaContainerWithDecisionLoudnessTypeStr:
		// mediaContainerWithDecisionLoudness.Str is populated
}
```
