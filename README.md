nvim-ox
=======

nvim-ox is a small Markdown preview plugin for Neovim. Its Go command renders
buffer text to one self-contained HTML page, without JavaScript or a persistent
server. It was inspired by [tatum](https://github.com/elijah-potter/tatum/).

## Requirements

- Go 1.25 or newer
- Neovim 0.10 or newer
- A system opener supported by `vim.ui.open`, such as `xdg-open` on Linux,
  for local desktop previews

## Install the Go command

From this directory, install the renderer and run the tests:

```sh
make install
make test
```

Installation defaults to `$HOME/go/bin/ox`. Ensure `$HOME/go/bin` is included
in the `PATH` inherited by Neovim. To choose a different destination, override
`PREFIX`:

```sh
make install PREFIX=/another/bin
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

The plugin finds `ox` on `PATH`; it does not construct a path relative to the
package. Keep the executable in a directory available to every environment
where Neovim runs.

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

## Headless SSH sessions

When Neovim runs on a remote VM, `vim.ui.open` also runs on that VM. Its
`xdg-open` call therefore looks for a desktop session and browser on the VM,
not on the laptop connected through SSH. In a typical headless session this
call fails, even when `xdg-open` is installed. The normal temporary HTML file
also exists only on the remote filesystem.

X11 forwarding is a separate graphical setup: it can display a browser
installed on the VM through the laptop's X display, but it is not required or
configured by nvim-ox.

To open a snapshot in the laptop's regular browser without running a server,
first create a new preview at a predictable path **on the VM**. This command
pipes the entire current buffer through the renderer:

```vim
:%write !ox > /tmp/ox-preview.html
```

This assumes `ox` is installed on the remote VM and available to its shell
`PATH`; use the absolute installed path if necessary. The `%` selects the
entire buffer, including unsaved changes. Writing to `!command`
pipes Markdown to the renderer; it does not save or modify the Markdown source
file. The shell redirection creates only the generated HTML snapshot. This
avoids having to recover the random temporary path used by the normal `:Ox`
browser flow.

Next, copy the snapshot **from the VM to the laptop** by running this command
on the laptop, replacing `user@vm` with the actual SSH destination:

```sh
scp user@vm:/tmp/ox-preview.html /tmp/ox-preview.html
```

Finally, open the laptop-local copy on Linux:

```sh
xdg-open /tmp/ox-preview.html
```

Or on macOS:

```sh
open /tmp/ox-preview.html
```

## Standalone command

The renderer is also a stdin-to-stdout command:

```sh
ox < README.md > /tmp/ox-preview.html
```

## Current limitations

Each `:Ox` invocation creates a new snapshot. There is no automatic reload.
Relative local images and links are not yet resolved against the Markdown
source directory. The preview intentionally has no JavaScript or external
stylesheets.
