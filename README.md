# tcluster
tcluster opens tmux panes with connections to remote hosts, similar to
clusterssh. It does not handle sending input to them - tmux supports
that by itself, see the [section tmux](#tmux).

Example:
```sh
$ cat ~/.tcluster.yaml
hosts:
  - full-partial-host.full.domain
  - full-partial-host-02.full.domain
  - another-host.full.domain
  - hostname-only
$ tcluster partial-host
```

Opens a new window with two panes, which are opening an ssh connection
to full-partial-host.full.domain and full-partial-host-02.full.domain.

Each argument is interpreted as a regular expression by golangs
[regexp](https://golang.org/pkg/regexp/) package.

For configuration examples see the `test_data` directory.

# Planned
Planned are two keywords to modify tclusters behaviour:

## then
`then` to support opening multiple windows at once with different
regexes.

Example:
```sh
$ cat ~/.tcluster.yaml
hosts:
  - full-partial-host.full.domain
  - full-partial-host-02.full.domain
  - another-host.full.domain
  - hostname-only
$ tcluster full-part then full.domain
```

To open a window with connections to full-partial-host{,02}.full.domain
and another window with connections to
full-partial-host{,02}.full.domain and another-host.full.domain

## -
`-` to specify that the next expression should not match the
specified hosts.
```sh
$ cat ~/.tcluster.yaml
hosts:
  - full-partial-host.full.domain
  - full-partial-host-02.full.domain
  - another-host.full.domain
  - hostname-only
$ tcluster full.domain - 02
```
Would open a new window with connections to
full-partial-host.full.domain and another-host.full.domain.

# tmux
Tmux supports inputting into multiple panes at once through the
window-option `synchronize-panes`.

Example:
```conf
bind S set -w synchronize-panes
```
Pressing prefix+S now toggles inputting into all panes of the window at
once.

paste-buffer currently does not work with synchronized panes,
a workaround is passing the input to send-keys:

```conf
bind '+' choose-buffer 'run "tmux send-keys $(tmux show-buffer -b %%)"'
```

However this doesn't preserve whitespace.
