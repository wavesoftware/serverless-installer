Installer for Openshift Serverless
==================================

[![Build Status](https://travis-ci.com/wavesoftware/serverless-installer.svg?branch=master)](https://travis-ci.com/wavesoftware/serverless-installer)

Installer will create a usable installation of Openshift's Serverless Platform.

## Usage

### Interactive

This command invoked without any parameters will enter TUI. 
User is asked couple questions and then his answers are save on disk to
be later used by deploy command.

```bash
serverless-installer
```

### Deploy

This command will deploy a serverless platform on currently configured Openshift cluster.

```bash
serverless-installer deploy --answers serverless-installer-answers.yaml
```

Answers are expected to be produced by interactive mode.

## Installation

### Go

if you have usable Go `>= 1.13` installation simply invoke:

```bash
go get -u github.com/wavesoftware/serverless-installer
```

This will produce a binary that you can use.

### Precompiled binaries

**To be determined!**

Procompiled binaries can be found in GitHub's releases page:

 * [Linux amd64](https://github.com/wavesoftware/serverless-installer/releases/latest)
 * [MacOSX amd64](https://github.com/wavesoftware/serverless-installer/releases/latest)
 * [Windows amd64](https://github.com/wavesoftware/serverless-installer/releases/latest)

## Development

To build simply invoke (it will check & test the code as well):

```bash
make
```

Go `>= 1.13` is required to build.

Contributions are welcome!

To contribute, follow the standard [git flow](http://danielkummer.github.io/git-flow-cheatsheet/) of:

1. Fork it
1. Create your feature branch (`git checkout -b feature/my-new-feature`)
1. Commit your changes (`git commit -am 'Add some feature'`)
1. Push to the branch (`git push origin feature/my-new-feature`)
1. Create new Pull Request

Even if you can't contribute code, if you have an idea for an improvement 
please open an [issue](https://github.com/wavesoftware/serverless-installer/issues).
