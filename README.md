# cbn

Clean Branch Name - simple program that remove unwanted chars from given text that can be savly used with git

## Install

```sh
go install github.com/egel/cbn/cmd/cbn@latest
```

## Upgrade

To upgrade to latest version use:

```sh
GONOPROXY=github.com/egel go install github.com/egel/cbn/cmd/cbn@latest
```

## Usage

```sh
# single strings
cbn "Am I new best feature?" # => am-i-new-best-feature

# hey! I also works as arg list
cbn This is fancy feature!  # => this-is-fancy-feature
```
