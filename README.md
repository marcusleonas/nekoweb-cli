> [!WARNING]  
> The nekoweb-cli is heavily under construction! Use at your own risk!

# nekoweb cli

a simple cli interface for the nekoweb api. currently doesn't support much,
but i'm planning to add everything from the nekoweb api.

## todo

- deploy command
- create/modify/delete files & folders
- import command

## implemented

```sh
nekoweb auth api_key
```

saves your api key to a ~/.nekoweb config file

```sh
nekoweb info [--domain, -d] [--all, -a]
```

if no flags provided, then will retrive your "main" site's info.
