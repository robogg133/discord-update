# Discord Linux Installer

A minimalist Discord installer for Linux systems, written in Go, that automatically downloads and installs the latest version of Discord. _(I made it cuz i was tired for downloading mannualy)_

## Overview

This command-line tool automates the Discord installation process on Linux systems. It downloads the official Discord package, extracts the files to the specified location, and executes the post-installation script.

## Features

- ✅ Automatic download of the latest Discord version
- ✅ Official tar.gz package extraction
- ✅ Preserves original file permissions
- ✅ Customizable installation directory
- ✅ Automatic post-installation script execution
- ✅ Detailed progress feedback
- ✅ Total installation time display

## Prerequisites

- **Go 1.18+** (for compilation)
- **Root access** (for system directory installation like `/usr/share`)
- **Internet connection** (to download Discord)
- **Linux** (specifically designed for Linux systems)

## Installation

### Via `go install`

To install directly with Go:

```bash
go install github.com/robogg133/discord-update@latest
```

---

## Manual compilation

```bash
# Clone the repository
git clone https://github.com/robogg133/discord-update.git
cd discord-update

# Compile the binary
go build .
```


## Basic syntax
```bash
discord-update [options]
```

## Options

| Option |  Description                | Default
|--------|----------------------------|--------
| `-i`   |  Discord installation Path | `/usr/share/discord`