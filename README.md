> [!WARNING]
> This Go project will output sensitive information to stdout, including secret details. This is intended for testing purposes only. Do not use this project with real (sensitive) secrets.

# Go Project

This project provides a `run` script to test the Bitwarden Go SDK.

## Prerequisites

- **Go** or **Nix** - Required to build the project
- **GitHub CLI (`gh`)** - Required for downloading test provider artifacts
- **Bitwarden Access Token** - Your API token for authentication

## Quick start

```sh
./run
```

Running the script with no arguments launches interactive mode, which guides you through the setup process.

## Environment Variables

- `API_URL` - Bitwarden API URL (required)
- `IDENTITY_URL` - Bitwarden Identity URL (required)
- `ORGANIZATION_ID` - Organization ID for the access token (required)
- `ACCESS_TOKEN` - Your Bitwarden access token (required)
- `EDITOR` or `VISUAL` - Used for editing the `.env` file in interactive mode (defaults to `vi`)

## Non-Interactive Mode

If you prefer to skip interactive mode and run commands directly:
1. Set the required values in the `.env` file: `API_URL`, `IDENTITY_URL`, `ORGANIZATION_ID`, and `ACCESS_TOKEN`
2. Source the `.env` file in your shell with: `. "$PWD/.env"`

### Test the latest production release of the Bitwarden Go SDK

```bash
./run prod
```

Test will be ran with the latest released Go SDK.

### Test a feature branch version of the Bitwarden Go SDK

```bash
./run test <branch-name>
```

Test will be ran with the Go SDK from the specified GitHub branch. The script will:

- Clone the sdk-sm repository to grab the Go source files on the specified branch
- Find the latest successful workflow run for the specified branch
- Download the required artifacts and move them into place in the local sdk-sm repo
- Overwrite the go.mod file to use the Go SDK from the feature branch
- Run the Go application (in `main.go`) with the Go SDK from the feature branch
