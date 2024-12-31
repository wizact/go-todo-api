# Build Builder Image
The make commands use a docker container as a host to build the binary files. This docker image is hosted on `ghcr`. In order to update the image with new dependencies follow the instructions below:

* Make sure you have generated GH Personal Access Token and export them as environment variables. Your username should be exported as `CR_URN` and your PAT should be exported as `CR_PAT`.
* Run `make login` to login docker to ghcr. 
* Update the dockerfile located at `./build/Dockerfile`
* Commit all the changes and ensure there is no untracked or uncomitted files in the folder. The hash of your commit will be used as the tag for the container image.
* Build the image using `make build-builder-image` command.
* You should be able to find the latest revision by running the `docker images | grep builder` command.
* Update the makefile `BUILD_IMAGE_VERSION` variable with the latest tag of the image.
* Push the image to the container registry using `docker push  ghcr.io/wizact/todo-api-builder:<new tag>`
