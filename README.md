# webpDAV
webpDav is a simple webDAV server for converting images to WebP.
This software creates a mountable WebDAV endpoint on your [home]server. Mounting webpDAV on your computer makes converting images (JPG and PNG) as easy as dragging them into the folder. webpDAV implements a file system watcher that checkes if a new image is placed into the share and automatically encodes the images to the WebP format using libwebp.

## Getting Started
The following instructions will get you a copy of the project up and running on your local machine for development and testing purposes. Using **docker** is the recommended installation, since the dependencie to libwebp is already handled within the Dockerfile. 

For those who want to build this project without docker please ensure to **install libwebp** on your system. 

### Requirements 
[lipwebp](https://developers.google.com/speed/webp/docs/api)

#### Install libwebp
For macOS use:
```bash
brew install webp
```

Linux users can use: 
```bash
sudo apt-get update 
sudo apt-get install libwebp-dev
```

### Docker 

Building the docker image within the source folder using the following command.

```bash
docker build -t "webpDAV:latest" . 
```

Then start the image using: 
```bash
docker run -p 8080:8080 webpDAV
```

Example docker-compose for easier start 
```yaml
version: "3.9"

services: 
  webpDAV: 
    build:
      context: .
      dockerfile: dockerfile
    container_name: webpDAV
    restart: unless-stopped
    ports: 
      - "8080:8080"
    volumes:
      - ./data:/data
```

### MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
