# Termfi

Termfi is a terminal command for uploading and sharing files using S3-compatible storage directly.

## Environment

Before installation, you need to have go > 1.22.x installed. You also need to have S3 compatible storage set up. For development purposes you can use `minio` docker container in `./docker` folder.

You can start `minio` docker containers using a sequence of commands:

```sh
cd ./docker
docker-compose up
```

this will start `minio` containers and nginx orchestrating them. Api can be accessed on port 9000 and the admin console is available on port 9001 with credential `acc: minioadmin` `pass: minioadmin`. There you can set up an access key that you will need when configuring the app.

## Installation

Build termfi to a directory that is in your PATH variable. (~/sbin for example)

```sh
go build <PATH_TO_DIRECTORY>/termfi ./main.go
```

Now you can start using termfi. For example:

```sh
termfi -h
```

prints to the command usage.

## Configuration

Before you can start using termfi upload and download you need to configure a connection to **S3 compatible** storage.

### Command

This can be done using the `config` subcommand.

For example:

```sh
termfi config --storageEndpoint http://localhost:9000 --storageAccessKeyId sNsGdTloGZ370SVzTV01 --storageAccessKey 7ay84IMDTnAZfUoT2cTChIEPHKbqdJKHRx7LBgS5 --storageBucket termfi
```

will configure termfi to your local endpoint on port **9000** with the provided credentials and **termfi** bucket.

_Note that bucket doesn't need to exist, the application will create it automatically and will also set the access policy to private write and public read so bear that in mind_

### Edit

You can also edit the configuration file directly. You can print the path to your configuration file using:

```sh
termfi config pwd
```

## Upload

Upload can be done using the `upload` subcommand.

Example:

```sh
termfi upload ./dir/file.txt
```

This command will upload the local file **./dir/file.txt** to the storage and print instructions on how to download this file.

The file can be downloaded using `termfi download` subcommand. It can be also downloaded using any other tool `curl` for example but this method will not preserve the original fire name because it will be prefixed with `termfi` metadata.

## Download

Download can be done using the `download` subcommand.

Example:

```sh
termfi download http://localhost:9000/ermfoo/tf-b40be25d-file.txt
```

This command will download the file to local `./file.txt`.

You can download any file using `termfi` but note that you will need to provide the flag `--no-check` if you want to download non-termfi uploaded file.

## Future

This is just a proof of concept. There is the possibility that some additional features will be added in the future.

### Ideas

- Client side encryption
- HTTP server instead of direct access to storage that will allow features such as login.
