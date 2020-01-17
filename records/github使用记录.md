# Github使用记录  
> 本文记录了笔者学习和使用github的过程。  

---
### 基础的代码上传、下载操作  
git add .  
git commit -m [message]  
git push [repository name] [branch name]  

git status  
git stash list  
git stash clear  
git stash  
git stash pop [stash@{[n]}]  
git pull [repository name] [branch name]  

---
### 初始化  
git init  
git clone [url]  

git remote -v   
git remote add [repository name] [url]   

---
### 分支操作  
git branch -[a/r]v 列出所有/远端/本地分支  
git branch [branch name] 新建分支  
git branch -d [branch name] 使用"-D"强制删除  
git checkout [branch name] 切换分支  
git checkout -b [branch name] 新建分支并切换到目标分支上  

git merge [branch name] 合并目标分支到当前分支上  
git merge [branch A] [branch B] 把A分支合并到B分支

---
### 标签
git tag [tag name] [commit id] 为指定commit生成tag  
git push remote [tag name] 提交指定tag  
git push remote --tags 提交所有tag

---
### 比较
git diff 比较*工作区*和*暂存区*的所有差异  
git diff [file name] 比较指定文件在*工作区*和*暂存区*的差异  

git diff --cache 比较*暂存区*和*HEAD*的所有差异  
git diff --cache [file name] 比较指定文件在*暂存区*和*HEAD*的差异  

git difftool [commit 1] [commit 2] 比较两个commit之间的差异

---
### 回退
git checkout [file name] 将*工作区*的指定文件恢复成*暂存区*中的样子  
git reset [file name] 将*暂存区*的指定文件恢复成*HEAD*中的样子  
git reset --hard 将*工作区*和*暂存区*的所有文件恢复成*HEAD*中的样子  

---
### 日志
git log --online --graph -n/--all 使用图显示*最近n个/所有*commit，每个commit显示为一行  
git log [file name] 查看涉及指定文件的所有commit  
git blame [file name] 查看指定文件各行最后修改所在的commit及作者

---
### 其他
git ls-files --others 查看没有使用git管理的文件  
git mergetool 解决conflict(?)

---
### github Pull requests (PR)  
github提供PR功能，使项目组外成员可以贡献自己的代码，也可以用来进行项目组内的code review。  
上面两种情况主要的区别在于PR发起人是否具有repository的提交权限：
 - 如果有权限，那么本地新建分支push、PR时选择[new branch] -> [master]即可；
 - 如果没有，则需要fork该代码仓库到自己的github，然后修改自己github中的仓库、在自己的仓库发起PR、选择[self repository.branch] -> [target branch]。    

PR支持compare and merge/review and comments/delete extra branch等操作。  

PR也可以通过命令行进行操作。  

---
### github issue
github可以通过issue管理开发进度和bug修复情况。  
一个issue主要包括title、comment、assign、labels和milestone；其中labels可以自定义，milestone需要提前创建。  
git commit的时候，可以通过在消息中添加指定标志，来操作issue（#xx中，xx为issue编号）：
 - fix #xx(fixes/fixed #xx)：关联到目标issue，打开issue可以看到本次提交
 - close #xx(closes/closed #xx)：关联并关闭目标issue

---
###### Mario
I've been pretending to work hard, but you're really growing up.
