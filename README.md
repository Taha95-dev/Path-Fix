# Pathfix

**Find missing commands. Fix your PATH.**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## The Problem

You type `nvim` and get:

'nvim' is not recognized as an internal or external command

You know it's installed. You know it's somewhere on your system. But your PATH doesn't know where.

## The Solution

```bash
pathfix nvim
```

## Pathfix:

1. Checks if the command is in your PATH
2. Searches common install locations
3. Tells you exactly where it is
4. Shows you how to add it to your PATH

## Installation
Download Binary
Download the latest release from Releases

## Or Build from Source

git clone https://github.com/YOUR_USERNAME/pathfix
cd pathfix
go build -o pathfix

## Usage

```bash
pathfix nvim
```

## Otput Example

🔍 Searching for 'nvim'...
✅ Found: C:\Users\NAME\scoop\apps\neovim\0.12.2\bin\nvim.exe

📌 To fix, add this to your PATH:
   setx PATH "%PATH%;C:\Users\NAME\scoop\apps\neovim\0.12.2\bin"

Then restart your terminal.


