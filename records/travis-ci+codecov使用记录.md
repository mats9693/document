# travis-ci + codecov使用记录  
> 本文记录了笔者使用[travis-ci](https://travis-ci.org/)和[codecov](https://codecov.io/)的过程。  
> 本文主要记录配置文件和脚本的一些细节，至于如何在网站上操作，请参考其他资料，我懒得截图了。  

---
### travis-ci
> 关于持续集成，我使用它只是为了在每次提交代码时，自动把代码覆盖率报告提交到codecov。  
> 所以如果你的重点在持续集成(continuous integration)、持续交付(continuous delivery)和持续部署(continuous deployment)，请自行参考[travis-ci的文档](https://docs.travis-ci.com/user/tutorial/)

```yaml  
language: go

before_install: chmod +x ./go.test.sh

script: ./go.test.sh

after_success: bash <(curl -s https://codecov.io/bash)
```
这就是最简单的travis-ci配置文件了：
 - 文件名：".travis.yml"，必须是这样，一点也不能错，包括使用.yaml扩展名都不行
 - language：指定[编程语言](https://docs.travis-ci.com/user/languages)（nodejs需要指定版本，java需要指定jdk版本）
 - before_install：这一行是给脚本文件添加执行权限，如果script没有单独写成文件，可以去掉这一行
 - script：你想要travis-ci帮你执行的脚本，成功执行即视为build passing
 - after_success：这一行是把代码覆盖率报告提交到codecov  
 
关于script执行的脚本文件：
```shell script  
#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic "$d"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
```
笔者不熟悉shell，从结果来看，这个脚本的功能是遍历脚本所在目录下的文件夹，执行go test并把所有测试报告整合为相同目录下的coverage.txt文件。  
 
travis-ci默认不会持续集成任何代码仓库，需要在它的网站上，为需要持续集成的代码仓库主动开启持续集成功能。

---
### codecov
> 值得一提的是，codecov不支持中文路径，如果你的项目里有中文路径，对应部分将被屏蔽。

codecov的使用则是需要一个token，可以在github的代码仓库中设置，也可以放在持续集成脚本中。

上面介绍了两个脚本的基本情况，并保证其可用([demo-leetcode](https://github.com/mats9693/leetcode), [demo-utils](https://github.com/mats9693/utils))。  

---
### summary
以上是自己的repo使用travis-ci和codecov的方法；  
对于组织的public repo，需要申请组织授权；  
其中，travis-ci在授权之后，还需要使用管理员账号启动指定repo的持续集成，详情请参考travis-ci的权限问题。    

---
###### Mario
I've been pretending to work hard, but you're really growing up.
