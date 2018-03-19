# Remove Mobijake

> Removes extra mobijake on file names downloaded using `youtube-dl`

## Running the program

Requires a path to the `--download-archive` path. `youtube-dl` allows you to keep a copy of all the file you download by providing this file path. Using this file, the program weeds out the mobijake from the file names in the specified folder.

> $ go run main.go -i /path/to/folder -a /path/to/archive

[![asciicast](https://asciinema.org/a/IW9h1ZQbB9REUzs4uSKxEnNtr.png)](https://asciinema.org/a/IW9h1ZQbB9REUzs4uSKxEnNtr)
