# MediaContainerWithDecisionTimed

Whether lyrics are timestamped.


## Supported Types

### 

```go
mediaContainerWithDecisionTimed := components.CreateMediaContainerWithDecisionTimedBoolean(bool{/* values here */})
```

### MediaContainerWithDecisionTimed2

```go
mediaContainerWithDecisionTimed := components.CreateMediaContainerWithDecisionTimedMediaContainerWithDecisionTimed2(components.MediaContainerWithDecisionTimed2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mediaContainerWithDecisionTimed.Type {
	case components.MediaContainerWithDecisionTimedTypeBoolean:
		// mediaContainerWithDecisionTimed.Boolean is populated
	case components.MediaContainerWithDecisionTimedTypeMediaContainerWithDecisionTimed2:
		// mediaContainerWithDecisionTimed.MediaContainerWithDecisionTimed2 is populated
}
```
