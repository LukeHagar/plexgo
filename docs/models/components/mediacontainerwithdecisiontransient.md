# MediaContainerWithDecisionTransient

Whether the subtitle is temporary or downloaded.


## Supported Types

### 

```go
mediaContainerWithDecisionTransient := components.CreateMediaContainerWithDecisionTransientBoolean(bool{/* values here */})
```

### MediaContainerWithDecisionTransient2

```go
mediaContainerWithDecisionTransient := components.CreateMediaContainerWithDecisionTransientMediaContainerWithDecisionTransient2(components.MediaContainerWithDecisionTransient2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mediaContainerWithDecisionTransient.Type {
	case components.MediaContainerWithDecisionTransientTypeBoolean:
		// mediaContainerWithDecisionTransient.Boolean is populated
	case components.MediaContainerWithDecisionTransientTypeMediaContainerWithDecisionTransient2:
		// mediaContainerWithDecisionTransient.MediaContainerWithDecisionTransient2 is populated
}
```
