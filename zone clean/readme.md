# TLD Zone Cleaner

This project is an interesting one. I made this project so that I could pull just the domains out from a TLD zone file from [CZDS](https://czds.icann.org)

## How to use it

Well, to use this project you need some extra stuff.
- First off, you need a [CZDS](https://czds.icann.org) account.
- Then, you need to request a TLD's zone file.
- Next, you need to download that file, and move the .txt into this directory.
- Finally, you run `go run main.go` and type in the name of the file. (Not including the .txt)