# gitclean
Guided Branch Cleanup Utility


# Install

`go get -u github.com/MCBrandenburg/gitclean`

# Usage
In the root of the project you want to clean, enter this in your prompt: `gitclean`

```
❰~/g/s/g/M/gitclean(OK readme)❱(0)≻ git branch -l
  demo
  master
* readme
❰~/g/s/g/M/gitclean(OK readme)❱(0)≻ gitclean
Remove branch demo? ('yes' to remove): yes
removing demo
Remove branch master? ('yes' to remove): no
❰~/g/s/g/M/gitclean(OK readme)❱(0)≻ git branch -l
  master
* readme
❰~/g/s/g/M/gitclean(OK readme)❱(0)≻
```