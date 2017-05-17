# branch
Get you some git branch without git!

Linux/Unix/MacOS only.

Branch will only print the shortname of the current branch you are on.
# Usage comparison
```
$ git status

On branch master
Your branch is up-to-date with 'origin/master'.
nothing to commit, working tree clean

git branch
* master

$ branch
master 
```
# Branch use case
```
$ git push origin $( branch )
  Counting objects: 3, done.
  Delta compression using up to 8 threads.
  Compressing objects: 100% (3/3), done.
  Writing objects: 100% (3/3), 403 bytes | 0 bytes/s, done.
  Total 3 (delta 1), reused 0 (delta 0)
  remote: Resolving deltas: 100% (1/1), completed with 1 local object.
  To github.com:alistanis/branch.git
     9f383f3..606c899  master -> master
```


