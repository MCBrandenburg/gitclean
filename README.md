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


## Delete Branches
```
❰~/g/s/g/M/test(OK test3)❱(0)≻ git branch -l
  development
  master
  test1
  test2
* test3
❰~/g/s/g/M/test(OK test3)❱(0)≻ gitclean
GITCLEAN_IGNORE set to 'master,development'
Remove branch test1[y,n,q]? y
Remove branch test2[y,n,q]? n
The following branches will be deleted:
	test1
Continue [y,n]? y
	removing test1
Exiting
❰~/g/s/g/M/test(OK test3)❱(0)≻ git branch -l
  development
  master
  test2
* test3
❰~/g/s/g/M/test(OK test3)❱(0)≻
```

## Quit During Branch Selection
```
❰~/g/s/g/M/testquit(OK test3)❱(0)≻ git branch -l
  development
  master
  test1
  test2
* test3
❰~/g/s/g/M/testquit(OK test3)❱(0)≻ gitclean
GITCLEAN_IGNORE set to 'master,development'
Remove branch test1[y,n,q]? y
Remove branch test2[y,n,q]? q
Exiting
❰~/g/s/g/M/testquit(OK test3)❱(0)≻ git branch -l
  development
  master
  test1
  test2
* test3
❰~/g/s/g/M/testquit(OK test3)❱(0)≻
```

## Don't Confirm
```
❰~/g/s/g/M/testno(OK test3)❱(0)≻ git branch -l
  development
  master
  test1
  test2
* test3
❰~/g/s/g/M/testno(OK test3)❱(0)≻ gitclean
GITCLEAN_IGNORE set to 'master,development'
Remove branch test1[y,n,q]? y
Remove branch test2[y,n,q]? y
The following branches will be deleted:
	test1
	test2
Continue [y,n]? n
Exiting
❰~/g/s/g/M/testno(OK test3)❱(0)≻ git branch -l
  development
  master
  test1
  test2
* test3
❰~/g/s/g/M/testno(OK test3)❱(0)≻
```