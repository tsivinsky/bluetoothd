# bluetooth daemon

It checks power of connected bluetooth devices and sends notifications if they're low

## Requirements

- bluetoothctl
- upower

## Usage

```bash
bluetoothd -i 5
```

`-i` flag sets interval in minutes for daemon to sleep between checks
