# File chunking

This project is something I made to get multiple files, and split every 25,000 lines into their own seperate file.

To be specific, I created this in conjunction with the [Beemo IDs](../beemo%20ids/) project so that I could mass report raid bots to [Dangerous Discord](https://dangercord.com)

## How to use it

To use this, simply run `go run main.go` inside the directory, and provide the full name of a file within the working directory (including file extension), if you want to add multiple files, respond with "Y" when prompted to do so, and provide another file.

Once you have all the files you want, reply with "N" and the code will start to spit it out in chunks of 25,000.