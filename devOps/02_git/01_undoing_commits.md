* commit is only local

* git reset --hard HEAD~1
    * HEAD~1 means we want to revert the commit last 1 commit

* git reset --soft HEAD~1 
    * to keep the changes in your working directory

* git commit --amend 
    * changes the last commit

* git reset --hard HEAD~1
    * This will remove commit from local but what if we already pushed it to remote
    * git push --force (after the above command)
    * don't do this in master or develop branch
    * only when working alone in a branch

* what is the alternate for reverting commit in master branch
    * git revert \<commit_hash\>
    * creates a new commit to revert the old commit's changes
    * git reset \<commit_hash\>
        * removes old commit