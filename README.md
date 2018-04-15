# oscalkit

[![CircleCI](https://circleci.com/gh/opencontrol/oscalkit.svg?style=svg)](https://circleci.com/gh/opencontrol/oscalkit) [![GoDoc](https://godoc.org/github.com/opencontrol/oscalkit?status.svg)](https://godoc.org/github.com/opencontrol/oscalkit)

> In development

Barebones Go SDK and CLI tool for parsing OSCAL, translating between OSCAL-formatted XML and JSON and for converting from OpenControl projects in to OSCAL.

## Installing

You can download the appropriate oscalkit command-line utility for your system from the [GitHub Releases](https://github.com/opencontrol/oscalkit/releases) page and run it from your local machine directly. For easier execution, you can include it in your `$PATH` environment variable. If you prefer, you can download and install via the included RPM/Deb packages on Linux or Homebrew recipe on macOS. A [Docker image](https://hub.docker.com/r/opencontrolorg/oscalkit/) is also made available on Docker Hub.

### Homebrew

    $ brew tap opencontrol/homebrew-oscalkit
    $ brew install oscalkit

### Docker

> Running the oscalkit Docker container requires either bind-mounting the directory containing your source files or passing file contents in to the command via stdin.

    $ docker pull opencontrolorg/oscalkit:<version>
    $ docker run -it --rm -v $PWD:/data -w /data opencontrolorg/oscalkit convert oscal-core.xml

or via stdin:

    $ docker run -it --rm opencontrolorg/oscalkit convert < oscal-core.xml

## Usage

```
NAME:
   oscalkit - OSCAL toolkit

USAGE:
   oscalkit [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     convert   convert between one or more OSCAL file formats
     validate  validate files against OSCAL XML and JSON schemas
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d    enable debug command output
   --help, -h     show help
   --version, -v  print the version
```

### Convert between XML and JSON

`oscalkit` can be used to convert one or more source files between OSCAL-formatted XML and JSON.

```
   oscalkit convert oscal - convert between one or more OSCAL file formats

USAGE:
   oscalkit convert oscal [command options] [source-files...]

DESCRIPTION:
   Convert between OSCAL-formatted XML and JSON files. The command accepts
   one or more source file paths and can also be used with source file contents
   piped/redirected from STDIN.

OPTIONS:
   --output-path value, -o value  Output path for converted file(s). Defaults to current working directory
   --output-file value, -f value  file name for converted output from STDIN. Defaults to "stdin.<json|xml|yaml>"
   --include-yaml                 If source file format is XML or JSON, also generate equivalent YAML output
```

#### Examples

Convert OSCAL-formatted NIST 800-53 declarations from XML to JSON:

    $ oscalkit convert oscal SP800-53-declarations.xml

Convert OSCAL-formatted NIST 800-53 declarations from XML to JSON via STDIN (note the use of "-"):

    $ cat SP800-53-declarations.xml | oscalkit convert oscal -

### Convert from OpenControl project to OSCAL [Experimental]

> Depends on usnistgov/OSCAL [#92](https://github.com/usnistgov/OSCAL/issues/92)

`oscalkit` also supports converting OpenControl projects to OSCAL-formatted JSON. You will need both the path to the `opencontrol.yaml` file and the `opencontrols/` directory which is created when you run a `compliance-masonry get` command.

```
NAME:
   oscalkit convert opencontrol - convert from OpenControl format to OSCAL "implementation" format

USAGE:
   oscalkit convert opencontrol [command options] [opencontrol.yaml-filepath] [opencontrols-dir-path]

DESCRIPTION:
   Convert OpenControl-formatted "component" and "opencontrol" YAML into
   OSCAL-formatted "implementation" layer JSON

OPTIONS:
   --yaml, -y  Generate YAML in addition to JSON
   --xml, -x   Generate XML in addition to JSON
```

### Examples

Convert OpenControl project to OSCAL-formatted JSON:

    $ oscalkit convert opencontrol ./opencontrol.yaml ./opencontrols/

### Validate against XML and JSON schemas

The tool supports validation of OSCAL-formatted XML and JSON files against the corresponding OSCAL XML schemas (.xsd) and JSON schemas. XML schema validation requires the `xmllint` tool on the local machine (included with macOS and Linux. Windows installation instructions [here](https://stackoverflow.com/a/21227833))

```
NAME:
   oscalkit validate - validate files against OSCAL XML and JSON schemas

USAGE:
   oscalkit validate [command options] [files...]

DESCRIPTION:
   Validate OSCAL-formatted XML files against a specific XML schema (.xsd)
   or OSCAL-formatted JSON files against a specific JSON schema

OPTIONS:
   --schema value, -s value  schema file to validate against
```

#### Examples

Validate FedRAMP framework in OSCAL-formatted JSON against the corresponding JSON schema

    $ oscalkit validate -s oscal-core.json fedramp-annotated-wrt-SP800-53catalog.json

## Developing

`oscalkit` is developed with [Go](https://golang.org/) (1.9+). If you have Docker installed, the included `Makefile` can be used to run unit tests and compile the application for Linux, macOS and Windows. Otherwise, the native Go toolchain can be used.

### Dependency management

The [`dep`](https://github.com/golang/dep) dependency management tool is used to vendor the application's dependencies. The `vendor/` folder containing the dependencies is checked in with the source. With [`dep`](https://github.com/golang/dep), you can verify the dependencies as follows:

    $ dep ensure

### Compile

You can use the included `Makefile` to generate binaries for your OS as follows (requires [Docker](https://docs.docker.com/engine/installation/)):

Compile for Linux:

    $ GOOS=linux GOARCH=amd64 make

Compile for macOS:

    $ GOOS=darwin GOARCH=amd64 make

Compile for Windows:

    $ GOOS=windows GOARCH=amd64 make

### Releasing

The [GoReleaser](https://goreleaser.com/) tool is used to publish `oscalkit` to GitHub Releases. The following release artifacts are currently supported:

- OSX binary
- Linux binary
- Windows binary
- Docker Image
- RPM package
- Deb package
- Homebrew recipe