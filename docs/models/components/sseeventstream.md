# SSEEventStream

Server-Sent Event stream with discriminated event types


## Supported Types

### SSECompleteEvent

```go
sseEventStream := components.CreateSSEEventStreamComplete(components.SSECompleteEvent{/* values here */})
```

### SSEConnectedEvent

```go
sseEventStream := components.CreateSSEEventStreamConnected(components.SSEConnectedEvent{/* values here */})
```

### SSEContentEvent

```go
sseEventStream := components.CreateSSEEventStreamContent(components.SSEContentEvent{/* values here */})
```

### SSEErrorEvent

```go
sseEventStream := components.CreateSSEEventStreamError(components.SSEErrorEvent{/* values here */})
```

### SSEHeartbeatEvent

```go
sseEventStream := components.CreateSSEEventStreamHeartbeat(components.SSEHeartbeatEvent{/* values here */})
```

### SSENotificationEvent

```go
sseEventStream := components.CreateSSEEventStreamNotification(components.SSENotificationEvent{/* values here */})
```

### SSEPermissionEvent

```go
sseEventStream := components.CreateSSEEventStreamPermission(components.SSEPermissionEvent{/* values here */})
```

### SSESessionCreatedEvent

```go
sseEventStream := components.CreateSSEEventStreamSessionCreated(components.SSESessionCreatedEvent{/* values here */})
```

### SSESessionDeletedEvent

```go
sseEventStream := components.CreateSSEEventStreamSessionDeleted(components.SSESessionDeletedEvent{/* values here */})
```

### SSEThinkingEvent

```go
sseEventStream := components.CreateSSEEventStreamThinking(components.SSEThinkingEvent{/* values here */})
```

### SSEToolExecutionCompleteEvent

```go
sseEventStream := components.CreateSSEEventStreamToolExecutionComplete(components.SSEToolExecutionCompleteEvent{/* values here */})
```

### SSEToolExecutionStartEvent

```go
sseEventStream := components.CreateSSEEventStreamToolExecutionStart(components.SSEToolExecutionStartEvent{/* values here */})
```

### SSEToolUseParameterDeltaEvent

```go
sseEventStream := components.CreateSSEEventStreamToolUseParameterDelta(components.SSEToolUseParameterDeltaEvent{/* values here */})
```

### SSEToolUseParameterStreamingCompleteEvent

```go
sseEventStream := components.CreateSSEEventStreamToolUseParameterStreamingComplete(components.SSEToolUseParameterStreamingCompleteEvent{/* values here */})
```

### SSEToolUseStartEvent

```go
sseEventStream := components.CreateSSEEventStreamToolUseStart(components.SSEToolUseStartEvent{/* values here */})
```

### SSEUserMessageCreatedEvent

```go
sseEventStream := components.CreateSSEEventStreamUserMessageCreated(components.SSEUserMessageCreatedEvent{/* values here */})
```

