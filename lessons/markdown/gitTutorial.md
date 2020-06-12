# Git 入门教程

## 安装 Git

- 如果你想在 Linux 上用二进制安装程序来安装基本的 Git 工具，可以使用发行版包含的基础软件包管理工具来安装。以 Fedora 为例，如果你在使用它（或与之紧密相关的基于 RPM 的发行版，如 RHEL 或 CentOS），你可以使用 dnf：

    ```shell
    sudo dnf install git-all
    ```

- 如果你在基于 Debian 的发行版上，如 Ubuntu，请使用 apt：

    ```shell
    sudo apt install git-all
    ```

- 要了解更多选择，Git 官方网站上有在各种 Unix 发行版的系统上安装步骤，网址为 [点击这里](https://git-scm.com/download/linux) 。

## 初次运行 Git 前的配置

- 安装完 Git 之后，要做的第一件事就是设置你的用户名和邮件地址。这一点很重要，因为每一个 Git 提交都会使用这些信息，它们会写入到你的每一次提交中，不可更改：

    ```shell
    git config --global user.name "John Doe"
    git config --global user.email johndoe@example.com
    ```

- 既然用户信息已经设置完毕，你可以配置默认文本编辑器了，当 Git 需要你输入信息时会调用它。如果未配置，Git 会使用操作系统默认的文本编辑器。如果你想使用不同的文本编辑器，例如 Emacs，可以这样做：

    ```shell
    git config --global core.editor emacs
    ```

## 获取 Git 仓库

- 在已存在目录中初始化仓库。

    ```shell
    cd /home/user/my_project
    git init
    ```

- 克隆现有的仓库，克隆仓库的命令是 git clone 。比如，要克隆 Git 的链接库 libgit2，可以用下面的命令：

    ```shell
    git clone https://github.com/libgit2/libgit2
    ```

## 记录每次更新到仓库

- 可以用 git status 命令查看哪些文件处于什么状态：

    ```shell
    git status
    ```

- 在项目下创建一个新的 README 文件。如果之前并不存在这个文件，使用 git status 命令，你将看到一个新的未跟踪文件。
- 使用命令 git add 开始跟踪一个文件。所以，要跟踪 README 文件，运行：

    ```shell
    git add README
    ```

- 修改一个文件，再使用 git add 命令可以暂存修改的文件。
- 忽略文件，我们可以创建一个名为 .gitignore 的文件，列出要忽略的文件的模式。来看一个实际的 .gitignore 例子：

    ```shell
    cat .gitignore
    *.[oa]
    *~
    ```

- 要查看尚未暂存的文件更新了哪些部分，不加参数直接输入 git diff：

    ```shell
    git diff
    ```

- 若要查看已暂存的将要添加到下次提交里的内容，可以用如下命令：

    ```shell
    git diff --staged
    ```

- 现在的暂存区已经准备就绪，可以提交了。在此之前，请务必确认还有什么已修改或新建的文件还没有 git add 过，否则提交的时候不会记录这些尚未暂存的变化。这些已修改但未暂存的文件只会保留在本地磁盘。所以，每次准备提交前，先用 git status 看下，你所需要的文件是不是都已暂存起来了，然后再运行提交命令 git commit，这样会启动你选择的文本编辑器来输入提交说明：

    ```shell
    git commit
    ```

- Git 提供了一个跳过使用暂存区域的方式，只要在提交的时候，给 git commit 加上 -a 选项，Git 就会自动把所有已经跟踪过的文件暂存起来一并提交，从而跳过 git add 步骤。
- 要从 Git 中移除某个文件，就必须要从已跟踪文件清单中移除（确切地说，是从暂存区域移除），然后提交。可以用 git rm 命令完成此项工作，并连带从工作目录中删除指定的文件，这样以后就不会出现在未跟踪文件清单中了。
- 另外一种情况是，我们想把文件从 Git 仓库中删除（亦即从暂存区域移除），但仍然希望保留在当前工作目录中。换句话说，你想让文件保留在磁盘，但是并不想让 Git 继续跟踪。当你忘记添加 .gitignore 文件，不小心把一个很大的日志文件或一堆 .a 这样的编译生成文件添加到暂存区时，这一做法尤其有用。为达到这一目的，使用 --cached 选项：

    ```shell
    git rm --cached README
    ```

- 在提交了若干更新，又或者克隆了某个项目之后，你也许想回顾下提交历史。完成这个任务最简单而又有效的工具是 git log 命令：

    ```shell
    git log
    ```

## 撤消操作

- 有时候我们提交完了才发现漏掉了几个文件没有添加，或者提交信息写错了。此时，可以运行带有 --amend 选项的提交命令来重新提交：

    ```shell
    git commit --amend
    ```

- 你提交后发现忘记了暂存某些需要的修改，可以像下面这样操作：

    ```shell
    git commit -m 'initial commit'
    git add forgotten_file
    git commit --amend
    ```

- 使用 git reset HEAD < file >... 来取消暂存。
- 撤消对文件的修改：

    ```shell
    git checkout -- CONTRIBUTING.md
    ```

## 远程仓库的使用

- 如果想查看你已经配置的远程仓库服务器，可以运行 git remote 命令。它会列出你指定的每一个远程服务器的简写。如果你已经克隆了自己的仓库，那么至少应该能看到 origin ——这是 Git 给你克隆的仓库服务器的默认名字：

    ```shell
    git clone https://github.com/schacon/ticgit
    ```

- 添加远程仓库：

    ```shell
    git remote add pb https://github.com/paulboone/ticgit
    ```

- 从远程仓库中抓取与拉取：

    ```shell
    git fetch <remote>
    ```

- 推送到远程仓库：

    ```shell
    git push origin master
    ```

- 查看某个远程仓库：

    ```shell
    git remote show origin
    ```

- 远程仓库的重命名：

    ```shell
    git remote rename pb paul
    ```

- 远程仓库删除：

    ```shell
    git remote remove paul
    ```

## Git 分支

- 分支创建:

    ```shell
    git branch testing
    ```

- 分支切换：

    ```shell
    git checkout testing
    ```

- 分支合并，假设你已经修正了 #53 问题，并且打算将你的工作合并入 master 分支。为此，你需要合并 iss53 分支到 master 分支，你只需要检出到你想合并入的分支，然后运行 git merge 命令：

    ```shell
    git checkout master
    git merge iss53
    ```

- 分支变基，你可以检出 experiment 分支，然后将它变基到 master 分支上：

    ```shell
    git checkout experiment
    git rebase master
    ```
