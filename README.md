学习网站：https://www.jianshu.com/nb/27703855

##beego + gorm(mysql)##

##创建beego项目
<p>在gopath目录src下，<br></p>
`bee new liteblog`

##运行beego项目
<p>在项目目录下：<br></p>
`bee run`

##git上传使用
一、初始化仓库，生成.git文件夹(只需一次)<br>
`git init`
<br><br>
二、设置(不是必须)<br>
`git config --global user.name '（名称）'`<br>
`git config --global user.email '（名称）'`<br>
<br>
三、添加文件/文件夹<br>
`git add 文件`<br>
`git add 文件夹`<br>
`git add .`<br>
<br>
四、提交<br>
Windows下使用双引号，Linux使用单引号:<br>
`git commit -m "备注"`<br><br>
以下会进入文本编辑(esc切换模式，esc+i插入模式，esc+:+wq保存退出)：<br>
`git commit`<br>
<br>
五、设置远程仓库连接(只需一次)<br>
1.先到github新建仓库，然后复制仓库地址<br>
2.在完成以上所有操作之后，执行：<br>
`git remote add orgin 仓库地址`<br>
(可用`git remote`查看是否已经创建连接)<br>
<br>
六、上传到远程仓库<br>
`git push -u orgin master(或者填分支名)`<br>
`git push`


##git忽略上传文件
1.在本地仓库创建.gitignore文件<br>
windows创建文件/文件夹:<br>
`mkdir 文件夹名`<br>
`echo 内容 > 文件名`<br>
`type 文件名`<br>

2.在.gitignore文件中写需要忽略的文件(包括路径也要写明，#开头)
<br>


##git 分支
作用：在本地仓库commit之后会有master仓库，这个是主仓库，一般是最终版本，而如果平时做修改或者未定版，都是用分支，分支的修改不会影响到主仓库。<br>
1.创建分支，这之前该本地仓库已经commit过，生成master了<br>
`git branch 分支名`<br>
<br>
2.切换分支<br>
`git checkout 分支名`<br>
<br>
3.合并分支<br>
`git merge 分支名`<br>