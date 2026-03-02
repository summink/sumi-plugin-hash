# Sumi Plugin Hash

## Description

The `hash` plugin calculates the hash value of a given string or file.

## Install

```bash
sumi plugin add hash@latest
```

## Usage

```bash
sumi hash [value] --file --mode [mode] --copy
```

### Arguments

- `value` - The value to hash.

- `--file` - If set, `value` is treated as a file path.

  - Otherwise, `value` is treated as a string.

- `--mode` - The hash method.

  - Default: `sha256`
  - Valid values: `md5`, `sha1`, `sha256`, `sha512`

- `--copy` - If set, the hash content is copied to the clipboard.
