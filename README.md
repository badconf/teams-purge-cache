# Microsoft Teams Cache Cleaner

## Overview
Microsoft Teams Cache Cleaner is a command-line tool written in Go that helps in managing the Microsoft Teams cache on a Windows system. It checks if Microsoft Teams is running, and then clears specific cache files used by Teams.

## Features
- Checks if Microsoft Teams is running, warn if yes.
- Purges specific cache files related to Microsoft Teams in these directories:
  - %appdata%\Microsoft\teams\tmp
  - %appdata%\Microsoft\teams\Blob_storage
  - %appdata%\Microsoft\teams\Cache
  - %appdata%\Microsoft\teams\IndexedDB
  - %appdata%\Microsoft\teams\GPUCache
  - %appdata%\Microsoft\teams\databases
  - %appdata%\Microsoft\teams\Local Storage

## Requirements
- Windows operating system.
- Administrative privileges for executing the program.

## Usage
Build the Go program and run the executable from the command line. No arguments are required.

```bash
go build tpurge.go
./tpurge.exe
```

Run the program in a console or terminal with the Windows session of the user for whom you want to purge the Teams cache.

## License

This project is open source and available under the [MIT License](LICENSE).
