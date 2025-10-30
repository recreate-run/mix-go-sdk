# PromptMode

Custom prompt handling mode:
- 'default': Use base system prompt only (customSystemPrompt ignored)
- 'append': Append customSystemPrompt to base system prompt (50KB limit)
- 'replace': Replace base system prompt with customSystemPrompt (100KB limit)


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `PromptModeDefault` | default             |
| `PromptModeAppend`  | append              |
| `PromptModeReplace` | replace             |