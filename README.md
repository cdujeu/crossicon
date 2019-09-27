# crossicon
Go tool for converting a PNG to ICO and to write the image as a byte slice inside a go file.

## Usage

```
Usage:
  ./crossicon [flags]

Flags:
  -h, --help             help for ./crossicon
      --bytes            Write to bytes files for both unix and windows
      --ico              Write ICO to file
  -i, --input string     Path to PNG input file
  -o, --output string    Prefix for writing output files
  -p, --package string   Package name
```

Flags input, output and package are mandatory.  
Use at least one of `--bytes` or `--ico` flags

## Credits

PNG to ICO conversion https://github.com/Kodeworks/golang-image-ico  
To Bytes array files inspired by https://github.com/cratonica/2goarray

## License

Apache License V2