# cchuffman

`cchuffman` is a command-line file compression tool written in Go, leveraging Huffman coding to efficiently compress and decompress files. It provides a simple yet powerful interface for reducing file sizes, making it easier to store and share data.

This is a Go study project based on the [Coding Challenges](https://codingchallenges.fyi) exercises, particularly [this one](https://codingchallenges.fyi/challenges/challenge-huffman).

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Compression**: Compresses files using Huffman coding, significantly reducing file size.
- **Decompression**: Decompresses files back to their original state.
- **Simple CLI**: User-friendly command-line interface.

## Installation

### Prerequisites

You can build and install `cchuffman` from source using Go. Make sure you have Go installed on your system.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/cchuffman.git
   cd cchuffman
   ```

1. Build the executable:
   ```bash
   make build
   ```

1. (Optional) Move the binary to a directory in your PATH for easy access:
   ```bash
   mv cchuffman /usr/local/bin/
   ```

## Usage

### Compress a File

To compress a file, use the following command:

```bash
./cchuffman --compress input_filepath [output_filepath]
```

* `input_filepath`: The path to the file you want to compress.
* `output_filepath`: The name of the output compressed file.

### Decompress a File

To decompress a file, use the following command:

```bash
./cchuffman --decompress input_filepath [output_filepath]
```

* `input_filepath`: The path to the compressed file.
* `output_filepath`: The name of the decompressed file to be generated.

### Help

For help or additional options, use:

```bash
./cchuffman --help
```

