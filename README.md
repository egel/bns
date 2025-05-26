# Branch Name Sanitizer

Branch Name Sanitizer - simple program that savely cleans text inputs focusing on comply with [Git repository naming conventions][weblink-git-scm-check-ref-format].

## Features

- Gently clean all accent characters from input, e.g: `é` => `e`, `ă` => `a`, etc.
- Persist characters meaning conversion, e.g. `ö` => `ou`, `&` => `-and-`, etc.
- Customization of strings connector
- Option to allow only ASCII characters
- Option to keep original case

## Install

```sh
go install github.com/egel/bns/cmd/bns@latest
```

### Git configuration

Add new gitconfig alias name e.g.: `cob` (acronym of `checkout -b`), `brn` (branch new), or anything else that suite your preferences.

```
# file ~/.gitconfig

[alias]
	cob = !sh -c 'git checkout -b $(bns $@)' -
```

then you can automatically checkout to a new, clean branch using command like following:

```sh
git cob "fancy name of your new branch"
```

## Usage

### Default

```sh
bns "Introduce new dashboard UI" # => introduce-new-dashboard-ui
```

### Customize string connector

flag: `-c`, `--connector`

```sh
bns -c "_" "Develop AI-powered bug reporter" # => develop_ai_powered_bug_reporter
```

### Keep original case

flag: `-o`, `--original-case`

```sh
bns -o "Task: Implement Login Feature" # => Task-Implement-Login-Feature
```

### Completly remove all Non-ASCII chars

flag: `-f`, `--force-ascii`.

If you need a strict non-ASCII complience.

```sh
# Arabic
bns -f "تحسين: amélioration du dashboard de l'utilisateur" # => amlioration-du-dashboard-de-l-utilisateur

# Chinese (mandarin)
bns -f "BUG：搜索功能不工作" # => bug
```

### Allow multiple string arguments

> [!NOTE]
> When passing multiple strings as arguments to your tool, make sure to enclose each one in quotes to ensure they are treated as input values rather than flags or options.

```sh
# hey! I also works as arg list!
bns This is fancy new feature!  # => this-is-fancy-new-feature
```

## Upgrade

To upgrade to the latest version use:

```sh
GONOPROXY=github.com/egel go install github.com/egel/bns/cmd/bns@latest
```

## Uninstall

1.  remove the binary from system

    ```sh
    rm $(command -v bns)
    ```

1.  remove the alias from your `~/.gitconfig`

## License

Apache-2.0 license

[weblink-git-scm-check-ref-format]: https://git-scm.com/docs/git-check-ref-format
