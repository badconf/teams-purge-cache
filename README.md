# Microsoft Teams Cache Cleaner

## Overview

Over time, the Microsoft Teams client can progressively become sluggish, leading users to experience performance issues. One of the solutions recommended by Microsoft to overcome this problem is [clearing the cache](https://learn.microsoft.com/en-us/microsoftteams/troubleshoot/teams-administration/clear-teams-cache). Although the process is straightforward, users who are not tech-savvy may find it challenging. By simply running the `tpurge.exe` program, they can effortlessly resolve their Microsoft Teams performance issues. This tool is designed to make the process of clearing the Teams cache as seamless as possible, offering a one-click solution for those who may be uncomfortable with more technical procedures.

## Features

- Checks if Microsoft Teams is running, and warn to close it if needed.
- Purges the cache directory related to Microsoft Teams: `%appdata%\Microsoft\Teams`

## Requirements

Windows operating system.

## Usage

Run the program in a console or terminal with the Windows session of the user for whom you want to purge the Teams cache. No arguments are required.

For those who want to build the executable from source, simply run:

```bash
go build tpurge.go
tpurge.exe
```

## License

This project is open source and available under the [MIT License](LICENSE).
