nvim-ox
=======

nvim-ox is a small Markdown preview plugin for Neovim. Its Go command renders
buffer text to one self-contained HTML page, without JavaScript or a persistent
server. It was inspired by [tatum](https://github.com/elijah-potter/tatum/).

## Requirements

- Go 1.25 or newer
- Neovim 0.10 or newer
- A system opener supported by `vim.ui.open`, such as `xdg-open` on Linux

## Build

From this directory, build the package-local executable and run the tests:

```sh
mkdir -p bin
go build -o bin/ox ./cmd
go test ./...
```

## Install as a native optional package

Place this directory at `~/.config/nvim/pack/plugged/opt/nvim-ox`, or create a
symlink there to this checkout. If `XDG_CONFIG_HOME` is set, use it in place of
`~/.config`.

For example, using a symlink:

```sh
mkdir -p "${XDG_CONFIG_HOME:-$HOME/.config}/nvim/pack/plugged/opt"
ln -s /absolute/path/to/nvim-ox "${XDG_CONFIG_HOME:-$HOME/.config}/nvim/pack/plugged/opt/nvim-ox"
```

Load the package from your Neovim configuration:

```lua
vim.cmd.packadd("nvim-ox")
```

The `bin/ox` executable is found relative to the package, so it does not need to
be installed globally.

## Preview a buffer

Open a Markdown buffer and run:

```vim
:Ox
```

The complete current buffer is sent to the renderer, so unsaved and unnamed
buffers work. The generated HTML is written to a temporary file and opened with
`vim.ui.open`. No server or RPC host keeps running.

An optional mapping is:

```lua
vim.keymap.set("n", "<leader>o", "<cmd>Ox<CR>", { desc = "Preview Markdown" })
```

To use a different executable location, set this before running `:Ox`:

```lua
vim.g.ox_binary = "/absolute/path/to/ox"
```

## Standalone command

The renderer is also a stdin-to-stdout command:

```sh
bin/ox < README.md > /tmp/ox-preview.html
```

## Current limitations

Each `:Ox` invocation creates a new snapshot. There is no automatic reload.
Relative local images and links are not yet resolved against the Markdown
source directory. The preview intentionally has no JavaScript or external
stylesheets.
