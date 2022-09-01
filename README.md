# Trivybeat

Welcome to Trivybeat 👋


## Getting Started with Trivybeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7
* [Mage](https://magefile.org/)

### Build from source code

Create directory and clone the repo:
```
mkdir -p ${GOPATH}/src/github.com/DmitryZ-outten/trivybeat
git clone https://github.com/DmitryZ-outten/trivybeat ${GOPATH}/src/github.com/DmitryZ-outten/trivybeat
```

Build the excutable for your OS:
```
mage build
```

### Run

Adjust the `trivybeat.yml` file for your needs. e.g. specify the trivy server and elasticsearch connection

To run Trivybeat with debugging output enabled, run:
```
./trivybeat -c trivybeat.yml -e -d "*"
```

### Test

To test Trivybeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```

### Cleanup

To clean  Trivybeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.

## Docker image

https://hub.docker.com/r/dmyz/trivybeat

