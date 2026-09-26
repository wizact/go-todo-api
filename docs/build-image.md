# Build Builder Image
The make commands use a docker container as a host to build the binary files. This docker image is hosted on `ghcr`. In order to update the image with new dependencies follow the instructions below:

* Make sure you have generated GH Personal Access Token and export them as environment variables. Your username should be exported as `CR_URN` and your PAT should be exported as `CR_PAT`.
* Run `make login` to login docker to ghcr. 
* Update `GO_VERSION` in the `Makefile`. The builder Dockerfile receives that version as a build argument.
* Build the image using `make build-builder-image` command.
* The image is tagged from `GO_VERSION` as `ghcr.io/wizact/todo-api-builder:go<version>`.
* Confirm the image is available locally with `docker image ls ghcr.io/wizact/todo-api-builder`.
* Push the image using `docker push ghcr.io/wizact/todo-api-builder:go<version>`.
