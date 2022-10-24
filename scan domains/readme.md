# TLD Zone Cleaner

This project is an interesting one. I made this project in conjunction with the [Zone Clean](../zone%20clean/) project, to scan the domains with regex. This project's intention is to scan all the domains for possibly malicious domains, usually typosquatting. 

## How to use it

- Get a list of domains, either from the [Zone Clean](../zone%20clean/) project, or by manually gathering them.
- Run `go run main.go`
- Give the script a file name, not including the .txt
- Give the script a regex to scan each domain with
- Done. Domains that match the regex will be appended to a new file.