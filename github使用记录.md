# Github使用记录  
记录自己学习和使用github的过程。  
---
### 基础的代码上传、下载操作  
git add .  
git commit -m [message]  
git push [repository name] [branch name]  

git status  
git stash list  
git stash  
git stash pop  
git pull [repository name] [branch name]  
---
### 初始化  
git init  
git clone [url]  
git remote -v   
git remote add [repository name] [url]   
---
### 分支操作  
git branch -v   
git branch [branch name] 新建分支  
git checkout [branch name] 切换分支，会将在当前分支上进行的修改一并转移  
git checkout -b [branch name] 新建分支并切换到目标分支上  

### github Pull requests (PR)  
github提供PR功能用来进行项目组外成员贡献自己的代码，也可以用来进行项目组内的code review。  
上面两种情况主要的区别在于PR发起人是否具有repository的提交权限：如果有权限，那么本地新建分支push、PR时选择[new branch] -> [master]即可；如果没有，则需要fork该代码仓库到自己的github，然后修改自己github中的仓库、在目标仓库发起PR、选择[self repository.branch] -> [target branch]。  
PR支持compare and merge/review and comments/delete extra branch等操作。  
PR也可以通过命令行进行操作。  
