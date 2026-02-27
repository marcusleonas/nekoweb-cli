> [!WARNING]  
> The nekoweb-cli is heavily under construction! Use at your own risk!

# nekoweb cli

a simple cli interface for the nekoweb api. currently doesn't support much,
but i'm planning to add everything from the nekoweb api.

## todo

- deploy command
- create/modify/delete files & folders
- import command

## installation

download the latest version from [releases](https://github.com/marcusleonas/nekoweb-cli/releases) and add it
somewhere in your `$PATH`

## implemented commands

`nekoweb auth <api_key>`<br>
authenticate with the nekoweb api

```sh
nekoweb auth <api_key>
```

**description:**

- saves your api key to `~/.nekoweb`
- used for authenticating all future requests

`nekoweb info`<br>
retrive site information

```sh
nekoweb info [flags]
```

if no flags are provided, retrives your **main site's** information.

**flags:**

`-d, --domain <domain>`<br>
get public information of a specific site

```sh
nekoweb info --domain n3bula.lol
```

`-a, --all`<br>
retrives information for all sites associated with your account

```sh
nekoweb info --all
```

`nekoweb ls <path>`<br>
lists all files in a directory on nekoweb

```sh
nekoweb ls <path>
```

**description:**

- if no `<path>` is provided, will list `/`
