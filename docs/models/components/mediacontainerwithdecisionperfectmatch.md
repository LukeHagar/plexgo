# MediaContainerWithDecisionPerfectMatch

Whether the subtitle is an exact match.


## Supported Types

### 

```go
mediaContainerWithDecisionPerfectMatch := components.CreateMediaContainerWithDecisionPerfectMatchBoolean(bool{/* values here */})
```

### MediaContainerWithDecisionPerfectMatch2

```go
mediaContainerWithDecisionPerfectMatch := components.CreateMediaContainerWithDecisionPerfectMatchMediaContainerWithDecisionPerfectMatch2(components.MediaContainerWithDecisionPerfectMatch2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mediaContainerWithDecisionPerfectMatch.Type {
	case components.MediaContainerWithDecisionPerfectMatchTypeBoolean:
		// mediaContainerWithDecisionPerfectMatch.Boolean is populated
	case components.MediaContainerWithDecisionPerfectMatchTypeMediaContainerWithDecisionPerfectMatch2:
		// mediaContainerWithDecisionPerfectMatch.MediaContainerWithDecisionPerfectMatch2 is populated
}
```
