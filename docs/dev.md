# Development Container

To bring up the development environment, make [devcontainer cli](https://github.com/devcontainers/cli) is installed and run the following command:

```
devcontainer up --workspace-folder .
```
Then get the container id `docker ps` and run the following command to enter the container:

```
docker exec -it <container_id> bash
```

You can now clone your dotfiles and start developing.


# Testing Domain Events

To test domain events, make sure nats client is installed and run the following command:

```
nats subscribe "User.NewUserRegistered" -s nats:4222
```

