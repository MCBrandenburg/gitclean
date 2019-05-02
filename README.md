# gitclean
Guided Branch Cleanup Utility


# Install

`go get -u github.com/MCBrandenburg/gitclean`

# Configuration

You can set standard branches to be ignored using the environment variable `GITCLEAN_IGNORE`

Example:

```
GITCLEAN_IGNORE=master,development
```

# Usage
In the root of the project you want to clean, enter this in your prompt: `gitclean`

```
❰~/g/s/g/M/test(OK master)❱(0)≻ git branch -l
  development
* master
  test1
  test2
❰~/g/s/g/M/test(OK master)❱(0)≻ gitclean
GITCLEAN_IGNORE set to 'master,development'
Remove branch test1[y,n,q]? n
Remove branch test2[y,n,q]? y
removing test2
❰~/g/s/g/M/test(OK master)❱(0)≻ git branch -l
  development
* master
  test1
❰~/g/s/g/M/test(OK master)❱(0)≻
```