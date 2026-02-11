# SSEEventStream

Server-Sent Event stream with discriminated event types


## Supported Types

### SSEConnectedEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEConnectedEvent(components.SSEConnectedEvent{/* values here */})
```

### SSEHeartbeatEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEHeartbeatEvent(components.SSEHeartbeatEvent{/* values here */})
```

### SSEErrorEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEErrorEvent(components.SSEErrorEvent{/* values here */})
```

### SSECompleteEvent

```go
sseEventStream := components.CreateSSEEventStreamSSECompleteEvent(components.SSECompleteEvent{/* values here */})
```

### SSEThinkingEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEThinkingEvent(components.SSEThinkingEvent{/* values here */})
```

### SSEContentEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEContentEvent(components.SSEContentEvent{/* values here */})
```

### SSEToolUseStartEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEToolUseStartEvent(components.SSEToolUseStartEvent{/* values here */})
```

### SSEToolUseParameterStreamingCompleteEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEToolUseParameterStreamingCompleteEvent(components.SSEToolUseParameterStreamingCompleteEvent{/* values here */})
```

### SSEToolUseParameterDeltaEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEToolUseParameterDeltaEvent(components.SSEToolUseParameterDeltaEvent{/* values here */})
```

### SSEToolExecutionStartEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEToolExecutionStartEvent(components.SSEToolExecutionStartEvent{/* values here */})
```

### SSEToolExecutionCompleteEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEToolExecutionCompleteEvent(components.SSEToolExecutionCompleteEvent{/* values here */})
```

### SSEPermissionEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEPermissionEvent(components.SSEPermissionEvent{/* values here */})
```

### SSENotificationEvent

```go
sseEventStream := components.CreateSSEEventStreamSSENotificationEvent(components.SSENotificationEvent{/* values here */})
```

### SSEUserMessageCreatedEvent

```go
sseEventStream := components.CreateSSEEventStreamSSEUserMessageCreatedEvent(components.SSEUserMessageCreatedEvent{/* values here */})
```

### SSESessionCreatedEvent

```go
sseEventStream := components.CreateSSEEventStreamSSESessionCreatedEvent(components.SSESessionCreatedEvent{/* values here */})
```

### SSESessionDeletedEvent

```go
sseEventStream := components.CreateSSEEventStreamSSESessionDeletedEvent(components.SSESessionDeletedEvent{/* values here */})
```

