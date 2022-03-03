# i3keyboardlayout

Custom block for [i3status-rs](https://github.com/greshake/i3status-rust) which allows switching between predefined keyboard layouts
![Preview](/preview.png?raw=true)

## Requirements

- [i3status-rs](https://github.com/greshake/i3status-rust)
- [setxkbmap](https://man.archlinux.org/man/setxkbmap.1.en)
- [Go](https://go.dev/)

## Installation

- Install i3keyboardlayout

    ```console
    go install github.com/julez-dev/i3keyboardlayout@latest
    ```

- Configure a new i3status-rs block

    Example with us and de keyboard

    ```toml
    [[block]]
    block = "custom"
    cycle = ["path/to/i3keyboardlayout us", "path/to/i3keyboardlayout de"]
    json = true
    ```
