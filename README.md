# Microsoft Teams Cache Cleaner

## Disclaimer

**This tool is in development**

Use this tool with caution, especially on production systems. It modifies system processes and files, and incorrect usage may lead to unexpected system behavior. Always test the tool in a controlled environment before using it on critical systems.

## Overview
Microsoft Teams Cache Cleaner is a command-line tool written in Go that helps in managing the Microsoft Teams cache on a Windows system. It checks if Microsoft Teams is running, closes it if it is, and then clears specific cache folders used by Teams.

## Features
- Checks if Microsoft Teams is running.
- Closes Microsoft Teams if it's currently active.
- Purges specific cache directories related to Microsoft Teams.

## Requirements
- Windows operating system.
- Administrative privileges for executing the program.

## Usage
Build the Go program and run the executable from the command line. No arguments are required.

```bash
go build
./tpurge
```

## License

This project is open source and available under the [MIT License](LICENSE).
