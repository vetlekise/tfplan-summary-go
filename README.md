# tfsummary
A Go command-line utility that parses and prints a summary of the Terraform plan output.

## Installation

**Using Go:**

```sh
go install github.com/vetlekise/tfsummary@latest
```

**Pre-built binaries:**

Download the latest release for your platform from the [GitHub Releases](https://github.com/vetlekise/tfsummary/releases) page.

## Usage

Export your Terraform plan as JSON, then pass it to the tool:

```sh
terraform plan -out=tfplan.binary
terraform show -json tfplan.binary > tfplan.json
tfsummary -path tfplan.json
```

The `-path` flag defaults to `tfplan.json` in the current directory, so if your file is named `tfplan.json` you can omit it:

```sh
tfsummary
```
