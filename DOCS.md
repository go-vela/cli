# NAME

vela - CLI for interacting with Vela and managing resources

# SYNOPSIS

vela

```
[--api.addr|-a]=[value]
[--api.token.access|--at]=[value]
[--api.token.refresh|--rt]=[value]
[--api.token|-t]=[value]
[--api.version|--av]=[value]
[--config|-c]=[value]
[--log.level|-l]=[value]
```

**Usage**:

```
vela [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
```

# GLOBAL OPTIONS

**--api.addr, -a**="": Vela server address as a fully qualified url (<scheme>://<host>)

**--api.token, -t**="": token used for communication with the Vela server

**--api.token.access, --at**="": access token used for communication with the Vela server

**--api.token.refresh, --rt**="": refresh access token used for communication with the Vela server

**--api.version, --av**="": API version for communication with the Vela server (default: v1)

**--config, -c**="": path to Vela configuration file (default: /Users/Z001NR1/.vela/config.yml)

**--log.level, -l**="": set the level of logging - options: (trace|debug|info|warn|error|fatal|panic) (default: info)


# COMMANDS

## login

Authenticate and login to Vela

**--api.addr, -a**="": Vela server address as a fully qualified url (<scheme>://<host>)

**--api.token, -t**="": token used for communication with the Vela server

**--api.token.access, --at**="": access token used for communication with the Vela server

**--api.token.refresh, --rt**="": refresh token used for communication with the Vela server

**--yes-all, -y**: auto-confirm all prompts (default: false)

## version

Output version information

**--output, --op**="": format the output in json, spew, wide or yaml

## add, a

Add resources to Vela via subcommands

### deployment

Add a new deployment from the provided configuration

**--description, -d**="": provide the description for the deployment (default: Deployment request from Vela)

**--org, -o**="": provide the organization for the deployment

**--output, --op**="": format the output in json, spew or yaml

**--parameter, -p**="": provide the parameter(s) within `key=value` format for the deployment

**--ref**="": provide the reference to deploy - this can be a branch, commit (SHA) or tag (default: refs/heads/master)

**--repo, -r**="": provide the repository for the deployment

**--target, -t**="": provide the name for the target deployment environment (default: production)

**--task, --tk**="": Provide the task for the deployment (default: deploy:vela)

### repo

Add a new repository from the provided configuration

**--active, -a**="": current status of the repository (default: true)

**--branch, -b**="": default branch for the repository (default: master)

**--clone, -c**="": full clone URL to repository in source control

**--event, -e**="": webhook event(s) repository responds to (default: [push pull_request])

**--link, -l**="": full URL to repository in source control

**--org, -o**="": provide the organization for the repository

**--output, --op**="": format the output in json, spew or yaml

**--private, -p**="": disable public access to the repository (default: false)

**--repo, -r**="": provide the name for the repository

**--timeout, -t**="": max time allowed per build in repository (default: 30)

**--trusted, --tr**="": elevated permissions for builds executed for repo (default: false)

**--visibility, -v**="": access level required to view the repository (default: public)

### secret

Add details of the provided secret

**--commands, -c**="": enable a secret to be used for a step with commands (default: true)

**--event, --ev**="": provide the event(s) that can access this secret (default: [deployment push tag])

**--file, -f**="": provide a file to add the secret(s)

**--image, -i**="": Provide the image(s) that can access this secret

**--name, -n**="": provide the name of the secret

**--org, -o**="": provide the organization for the secret

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the repository for the secret

**--secret.engine, -e**="": provide the engine that stores the secret (default: native)

**--secret.type, --ty**="": provide the type of secret being stored (default: repo)

**--team, -t**="": provide the team for the secret

**--value, -v**="": provide the value for the secret

## cancel, cx

Cancel a resource for Vela via subcommands

### build

Cancel the provided build

**--build, -b**="": provide the number for the build (default: 0)

**--org, -o**="": provide the organization for the build

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the repository for the build

## chown, c

Change ownership of resources for Vela via subcommands

### repo

Change ownership of the provided repository

**--org, -o**="": provide the organization for the repository

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the name for the repository

## compile

Compile a resource for Vela via subcommands

### pipeline

Compile the provided pipeline

**--org, -o**="": provide the organization for the pipeline

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--ref**="": provide the repository reference for the pipeline (default: master)

**--repo, -r**="": provide the repository for the pipeline

## exec

Execute a resource for Vela via subcommands

### pipeline

Execute the provided pipeline

**--file, -f**="": provide the file name for the pipeline (default: .vela.yml)

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--path, -p**="": provide the path to the file for the pipeline

## expand

Expand a resource for Vela via subcommands

### pipeline

Expand the provided pipeline

**--org, -o**="": provide the organization for the pipeline

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--ref**="": provide the repository reference for the pipeline (default: master)

**--repo, -r**="": provide the repository for the pipeline

## generate, gn

Generate resources for Vela via subcommands

### completion

Generate a shell auto completion script

**--bash, -b**="": generate a bash auto completion script (default: false)

**--zsh, -z**="": generate a zsh auto completion script (default: false)

### config

Generate the config file used in the CLI

**--api.addr, -a**="": Vela server address as a fully qualified url (<scheme>://<host>)

**--api.token, -t**="": token used for communication with the Vela server

**--api.token.access, --at**="": access token used for communication with the Vela server

**--api.token.refresh, --rt**="": refresh token used for communication with the Vela server

**--api.version, --av**="": API version for communication with the Vela server

**--log.level, -l**="": set the level of logging - options: (trace|debug|info|warn|error|fatal|panic)

**--org, -o**="": provide the organization for the CLI

**--output, --op**="": format the output in json, spew, or yaml format

**--repo, -r**="": provide the repository for the CLI

**--secret.engine, -e**="": provide the secret engine for the CLI

**--secret.type, --ty**="": provide the secret type for the CLI

### pipeline

Generate a valid Vela pipeline

**--file, -f**="": provide the file name for the pipeline (default: .vela.yml)

**--path, -p**="": provide the path to the file for the pipeline

**--stages, -s**="": enable generating the pipeline with stages (default: false)

**--type, -t**="": provide the type of pipeline being generated

## get, g

Get a list of resources for Vela via subcommands

### build, builds

Display a list of builds

**--org, -o**="": provide the organization for the build

**--output, --op**="": format the output in json, spew, wide or yaml

**--page, -p**="": print a specific page of builds (default: 1)

**--per.page, --pp**="": number of builds to print per page (default: 10)

**--repo, -r**="": provide the repository for the build

### deployment, deployments

Display a list of deployments

**--org, -o**="": provide the organization for the deployment

**--output, --op**="": format the output in json, spew, wide or yaml

**--page, -p**="": print a specific page of deployments (default: 1)

**--per.page, --pp**="": number of deployments to print per page (default: 10)

**--repo, -r**="": provide the repository for the deployment

### hook, hooks

Display a list of hooks

**--org, -o**="": provide the organization for the hook

**--output, --op**="": format the output in json, spew, wide or yaml

**--page, -p**="": print a specific page of hooks (default: 1)

**--per.page, --pp**="": number of hooks to print per page (default: 10)

**--repo, -r**="": provide the repository for the hook

### log, logs

Display a list of build logs

**--build, -b**="": provide the build for the log (default: 0)

**--org, -o**="": provide the organization for the log

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the repository for the log

### repo, repos

Display a list of repositories

**--output, --op**="": format the output in json, spew, wide or yaml

**--page, -p**="": print a specific page of repositories (default: 1)

**--per.page, --pp**="": number of repositories to print per page (default: 10)

### secret, secrets

Display a list of secrets

**--org, -o**="": provide the organization for the secret

**--output, --op**="": format the output in json, spew, wide or yaml

**--page, -p**="": print a specific page of secrets (default: 1)

**--per.page, --pp**="": number of secrets to print per page (default: 10)

**--repo, -r**="": provide the repository for the secret

**--secret.engine, -e**="": provide the engine that stores the secret (default: native)

**--secret.type, --ty**="": provide the type of secret being stored (default: repo)

**--team, -t**="": provide the team for the secret

### service, services

Display a list of services

**--build, -b**="": provide the build for the service (default: 0)

**--org, -o**="": provide the organization for the build

**--output, --op**="": format the output in json, spew, wide or yaml

**--page, -p**="": print a specific page of services (default: 1)

**--per.page, --pp**="": number of services to print per page (default: 10)

**--repo, -r**="": provide the repository for the build

### step, steps

Display a list of steps

**--build, -b**="": provide the build for the step (default: 0)

**--org, -o**="": provide the organization for the step

**--output, --op**="": format the output in json, spew, wide or yaml

**--page, -p**="": print a specific page of steps (default: 1)

**--per.page, --pp**="": number of steps to print per page (default: 10)

**--repo, -r**="": provide the repository for the step

## remove, r

Remove a resource for Vela via subcommands

### config

Remove the config file used in the CLI

**--api.addr, -a**="": removes the API addr from the config file (default: false)

**--api.token, -t**="": removes the API token from the config file (default: false)

**--api.token.access, --at**="": access token used for communication with the Vela server

**--api.token.refresh, --rt**="": refresh token used for communication with the Vela server

**--api.version, --av**="": removes the API version from the config file (default: false)

**--log.level, -l**="": removes the log level from the config file (default: false)

**--org, -o**="": removes the org from the config file (default: false)

**--output, --op**="": removes the output from the config file (default: false)

**--repo, -r**="": removes the repo from the config file (default: false)

**--secret.engine, -e**="": removes the secret engine from the config file (default: false)

**--secret.type, --ty**="": removes the secret type from the config file (default: false)

### repo

Remove the provided repository

**--org, -o**="": provide the organization for the repository

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the name for the repository

### secret

Remove details of the provided secret

**--name, -n**="": provide the name of the secret

**--org, -o**="": provide the organization for the secret

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the repository for the secret

**--secret.engine, -e**="": provide the engine that stores the secret (default: native)

**--secret.type, --ty**="": provide the type of secret being stored (default: repo)

**--team, -t**="": provide the team for the secret

## repair, rp

Repair a resource for Vela via subcommands

### repo

Repair settings of the provided repository

**--org, -o**="": provide the organization for the repository

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the name for the repository

## restart, rs

Restart a resource for Vela via subcommands

### build

Restart the provided build

**--build, -b**="": provide the number for the build (default: 0)

**--org, -o**="": provide the organization for the build

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the repository for the build

## update, u

Update a resource for Vela via subcommands

### config

Update the config file used in the CLI

**--api.addr, -a**="": update the API addr in the config file

**--api.token, -t**="": update the API token in the config file

**--api.token.access, --at**="": access token used for communication with the Vela server

**--api.token.refresh, --rt**="": refresh token used for communication with the Vela server

**--api.version, --av**="": update the API version in the config file

**--log.level, -l**="": update the log level in the config file

**--org, -o**="": update the org in the config file

**--output, --op**="": update the output in the config file

**--repo, -r**="": update the repo in the config file

**--secret.engine, -e**="": update the secret engine in the config file

**--secret.type, --ty**="": update the secret type in the config file

### repo

Update a new repository from the provided configuration

**--active, -a**="": current status of the repository (default: true)

**--branch, -b**="": default branch for the repository (default: master)

**--clone, -c**="": full clone URL to repository in source control

**--event, -e**="": webhook event(s) repository responds to (default: [push pull_request])

**--link, -l**="": full URL to repository in source control

**--org, -o**="": provide the organization for the repository

**--output, --op**="": format the output in json, spew or yaml

**--private, -p**="": disable public access to the repository (default: false)

**--repo, -r**="": provide the name for the repository

**--timeout, -t**="": max time allowed per build in repository (default: 30)

**--trusted, --tr**="": elevated permissions for builds executed for repo (default: false)

**--visibility, -v**="": access level required to view the repository (default: public)

### secret

Update details of the provided secret

**--commands, -c**="": enable a secret to be used for a step with commands (default: true)

**--event, --ev**="": provide the event(s) that can access this secret (default: [deployment push tag])

**--file, -f**="": provide a file to update the secret(s)

**--image, -i**="": provide the image(s) that can access this secret

**--name, -n**="": provide the name of the secret

**--org, -o**="": provide the organization for the secret

**--output, --op**="": Print the output in default, yaml or json format

**--repo, -r**="": provide the repository for the secret

**--secret.engine, -e**="": provide the engine that stores the secret (default: native)

**--secret.type, --ty**="": provide the type of secret being stored (default: repo)

**--team, -t**="": provide the team for the secret

**--value, -v**="": provide the value for the secret

## validate, vd

Validate a resource for Vela via subcommands

### pipeline

Validate a Vela pipeline

**--file, -f**="": provide the file name for the pipeline (default: .vela.yml)

**--org, -o**="": provide the organization for the pipeline

**--path, -p**="": provide the path to the file for the pipeline

**--ref**="": provide the repository reference for the pipeline (default: master)

**--repo, -r**="": provide the repository for the pipeline

**--template**: enables validating a pipeline with templates

## view, v

View details for a resource for Vela via subcommands

### build

View details of the provided build

**--build, -b**="": provide the number for the build (default: 0)

**--org, -o**="": provide the organization for the build

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--repo, -r**="": provide the repository for the build

### config

View the config file used in the CLI

### deployment

View details of the provided deployment

**--deployment, -d**="": provide the number for the deployment (default: 0)

**--org, -o**="": provide the organization for the deployment

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--repo, -r**="": provide the repository for the deployment

### hook

View details of the provided hook

**--hook**="": provide the number for the hook (default: 0)

**--org, -o**="": provide the organization for the hook

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--repo, -r**="": provide the repository for the hook

### log

View details of the provided log

**--build, -b**="": provide the build for the log (default: 0)

**--org, -o**="": provide the organization for the log

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the repository for the log

**--service**="": provide the service for the log (default: 0)

**--step**="": provide the step for the log (default: 0)

### pipeline

View details of the provided pipeline

**--org, -o**="": provide the organization for the pipeline

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--ref**="": provide the repository reference for the pipeline (default: master)

**--repo, -r**="": provide the repository for the pipeline

### repo

View details of the provided repo

**--org, -o**="": provide the organization for the repository

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--repo, -r**="": provide the name for the repository

### secret

View details of the provided secret

**--name, -n**="": provide the name of the secret

**--org, -o**="": provide the organization for the secret

**--output, --op**="": format the output in json, spew or yaml

**--repo, -r**="": provide the repository for the secret

**--secret.engine, -e**="": provide the engine that stores the secret (default: native)

**--secret.type, --ty**="": provide the type of secret being stored (default: repo)

**--team, -t**="": provide the team for the secret

### service

View details of the provided service

**--build, -b**="": provide the build for the service (default: 0)

**--org, -o**="": provide the organization for the service

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--repo, -r**="": provide the repository for the service

**--service, -s**="": provide the number for the service (default: 0)

### step

View details of the provided step

**--build, -b**="": provide the build for the step (default: 0)

**--org, -o**="": provide the organization for the step

**--output, --op**="": format the output in json, spew or yaml (default: yaml)

**--repo, -r**="": provide the repository for the step

**--step, -s**="": provide the number for the step (default: 0)