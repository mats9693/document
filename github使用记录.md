# Github使用记录  
记录自己学习和使用github的过程。  

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
git branch -v   
git branch [branch name] 新建分支  
git branch -d [branch name]  
git checkout [branch name] 切换分支  
git checkout -b [branch name] 新建分支并切换到目标分支上  

git merge [branch name] 合并目标分支到当前分支上  

---
### github Pull requests (PR)  
github提供PR功能用来进行项目组外成员贡献自己的代码，也可以用来进行项目组内的code review。  
上面两种情况主要的区别在于PR发起人是否具有repository的提交权限：如果有权限，那么本地新建分支push、PR时选择[new branch] -> [master]即可；如果没有，则需要fork该代码仓库到自己的github，然后修改自己github中的仓库、在目标仓库发起PR、选择[self repository.branch] -> [target branch]。  
PR支持compare and merge/review and comments/delete extra branch等操作。  
PR也可以通过命令行进行操作。  

---
### github issue
github可以通过issue管理开发进度和bug修复情况。
一个issue主要包括标题、内容、指派、标签和节点，其中标签可以自定义，节点需要提前创建。
git commit的时候，可以通过在消息中添加：
 - fix #xx(fixes/fixed #xx)：关联到目标issue，打开issue可以看到本次提交
 - close #xx(closes/closed #xx)：关闭目标issue
