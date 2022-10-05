# yamllint-to-checkstyle

A commandline program that takes [yamlint](https://github.com/adrienverge/yamllint) parsable output and converts it to [checkstyle](https://checkstyle.sourceforge.io/) format.

## Installation

```shell
go install github.com/hunts/yamllint-to-checkstyle@latest
```

## Usage

```shell
yamllint --format=parsable file.yaml | yamllint-to-checkstyle
```
