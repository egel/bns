# cbn

Clean Branch Name - simple program that remove unwanted characters from given text that can be savly used within git branch.

## Features

- Gently clean unwanted characters from given strings
- Option to allow only ASCII characters
- Option to keep original case
- Customization of a string connector
- Limit output characters

## Install

```sh
go install github.com/egel/cbn/cmd/cbn@latest
```

### Git configuration

Add new gitconfig alias e.g.: `cob` (acronym of `checkout -b`) or `brn` (branch new), but you can use anything suite you.

```
# file ~/.gitconfig

[alias]
	cob = !sh -c 'git checkout -b $(cbn $@)' -
```

then you can use it like following and automatically checkout to clean branch

```sh
git cob "fancy name of your new branch"
```

## Usage

### Gently clean text

```sh
cbn "Am I new best feature?" # => am-i-new-best-feature
```

### Keep original case

```sh
cbn -o "Task: Implement Login Feature" # => Task-Implement-Login-Feature
```

### Completly remove all Non-ASCII chars

```sh
# Arabic
cbn -f "تحسين: amélioration du dashboard de l'utilisateur" # => amlioration-du-dashboard-de-l-utilisateur

# Chinese (mandarin)
cbn -f "BUG：搜索功能不工作" # => bug
```

### Allow multiple string arguments

> [!NOTE]
> While using with multiple strings arguments, pay attention as without quotes, some arguments may be understood as flags and trigger option command instead of desired output.

```sh
# hey! I also works as arg list!
cbn This is fancy feature!  # => this-is-fancy-feature
```

## Upgrade

To upgrade to the latest version use:

```sh
GONOPROXY=github.com/egel go install github.com/egel/cbn/cmd/cbn@latest
```

## Uninstall

1.  remove the binary from system

    ```sh
    rm $(command -v cbn)
    ```

1.  remove the alias from your `~/.gitconfig`

## License

Apache-2.0 license
