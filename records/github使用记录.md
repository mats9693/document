# Github使用记录  
> 本文记录了笔者学习使用git和github的过程。  
> 推荐一个git图文讲解+实际练习的[网站](https://learngitbranching.js.org/)，[github地址](https://github.com/pcottle/learnGitBranching)

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
（**pull**是**fetch+merge**的语法糖）  

#### 个人习惯
git stash  
git pull  
git stash pop  
...handle conflict if exist...  
git add .  
git commit -m [message]  （需要比较代码、部分提交时，会使用ide进行提交）  
git push  

---
### 初始化  
git init  
git clone [url]  

git remote -v   
git remote add [repository name] [url]   

---
### 分支  
git branch -[a/r]v 列出所有/远端/本地分支  
git branch [branch name] 新建分支  
git branch -d [branch name] 删除分支（使用"-D"强制删除）  
git branch -f [branch A] [branch B or commit B] 将A分支强制移动到B分支或B提交  
git checkout [branch name] 切换分支  
git checkout -b [branch name] 新建分支并切换到目标分支上  
（checkout实际上操作的是**HEAD**）  

git merge [branch name(s)] 合并目标分支到当前分支上  
git merge [branch A] [branch B] 把A分支合并到B分支
git rebase [branch A] [branch B] 把B分支移动到A分支上（B分支可以为空，为空时默认为当前分支）  

---
### 标签
git tag -l 列举所有标签  
git tag [tag name] [commit id] -m [message] 为指定commit生成tag，包含描述信息  
git tag -m [message] [tag name -f]  为指定tag更新描述信息  
git tag -d [tag name] 删除指定tag  
git tag -v [tag name] 显示指定tag详情  
git push [repository name] [tag name] 提交指定tag
git push [repository name] :refs/tags/[del tag name] 向远程仓库提交“删除tag”操作  
git push [repository name] --tags 提交所有tag

git describe [commit id] --tags 输出格式：\<tag>_\<numCommits>_g\<hash>，返回距离给出的提交最近的tag，
如果该提交与tag指向的提交不同，则注明两次提交的距离，以及给出的提交的部分hash；否则只返回tag名称

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

git reset [commit id] 恢复为指定提交  
git revert [commit id] “撤销为该提交的上一次提交”，并将本次撤销操作生成一次提交，提交到当前分支上  

---
### 日志
git log --online --graph -n/--all 使用图显示*最近n个/所有*commit，每个commit显示为一行  
git log [file name] 查看涉及指定文件的所有commit  
git blame [file name] 查看指定文件各行最后修改所在的commit及作者

---
### 整理提交记录
git cherry-pick [commit id(s)] 在当前分支上新增指定的提交记录

git rebase -interactive/-i [commit id] 修改指定提交之后的提交，包括删除和调整顺序

git commit --amend 在当前提交上，追加修改

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
### git stash clear误删除与恢复
#### 核心命令：  
$ git log --graph --oneline --decorate $(git fsck --no-reflog | awk '/dangling commit/ {print $3}')  
$ git stash apply [commit id, e4a07d4 in this time] (WIP record)  
$ git show [commit id]  
#### 部分执行结果：
```text  
Checking object directories: 100% (256/256), done.  
Checking objects: 100% (5645/5645), done.  
* 201284c index on master: 5c1fd9f background picture and npm lint format
| *   e4a07d4 WIP on master: 5c1fd9f background picture and npm lint format
| |\
|/ /
| * fded4cd index on master: 5c1fd9f background picture and npm lint format
|/
* 5c1fd9f (HEAD -> master, origin/master, origin/HEAD) background picture and npm lint format
* 2d2ba64 Change the data representation
| *   5a0304b WIP on master: 0c0309b blockchain format as ui-cut
... ...（折叠了历史记录）
```

---
### 总结
1. git的核心是一条条的提交记录（commit record），而分支则像是一个指向提交记录的指针。  
   - 想一想提交记录的相对引用（“^”、“~<num>”），分支是不是很像一个链表的指针？
1. git checkout实际上移动的是HEAD“指针”，HEAD可以指向一个分支，也可以直接指向一个提交。  
1. git为每一个提交生成了一个hash，基于sha-1算法，拥有40位16进制数。  

---
###### Mario
I've been pretending to work hard, but you're really growing up.
