# Notes-cli

Notes-cli is a command-line interface for quickly taking and organizing notes.

## Getting Started

### Installation

1. Download the appropriate binary from the list of [releases](https://github.com/joshmalbrecht/notes-cli/releases).

2. Rename the binary how ever you would like and install the binary in a directory that exists on your $PATH (Example: `/usr/local/bin/`)

### Configuration

The CLI reads a configuration file named `config.json` located in the `.config` directory in the user's home directory.
If the configuration directory does not exist, you can initialize the directory and file using `notes-cli init`.

The following configurations can be configured in `config.json`:

| Key             | Example Values          | Description                                                         |
| --------------- | ---------------------- | ------------------------------------------------------------------  |
| "NotesLocation" | "/home/user" | Defines the absolute path where the notes will be stored. |
| "FileExtension" | "txt and md" | Defines file extension that will be used to save the notes files. The default value is "md". |
| "TextEditorCommand" | "vi, nano, and vim" | Command that invokes a text editor with a specified filename as an argument, allowing the user to view and modify the contents of that file. The default value is "vi". |

