# 流程

## 发起 Merge Request，assign 给 peer 做 review

    - 代码变动要尽量小且专注于一个任务，不要攒的很大，或者做多个任务，要保证审查者可以较快、较容易的 review。需要一次性提交大量不需要 review 的文件的（比如框架、依赖包等），分两个 commit，不需要 review 的放在第一个 commit，并在 MR 描述中说明。
    - 如果与目标分支有冲突，提交者应该自己使用 git rebase 或 git merge（共享分支的情况）解决。
    - Assign 给别人之前一定要自己先 review 一遍（在 GitLab 或者其它 review 工具上检查最终效果），确保自己提交的每一行变更都是正确的、必要的，对自己的代码负责，不要浪费别人的时间。

## 审查者 review 代码

    - 对《编写整洁的代码》中各项要求进行检查，在任何有疑问或建议的地方留评论。
    - 如果 review 工具是 GitLab 且提交者已针对之前的评论做了修复，审查者需在确认问题已修复后 resolve discussion（解决讨论）。
    - 从中学习一些好的东西。
    - Review 完后，assign 给提交者处理；较小的问题可以不必 assign 回去，提交者也需保证及时响应。

## 提交者响应评论

    - 确实有问题的，修复之。如果该分支未被其他人使用，应使用 git commit --amend 提交以减少不必要的 commit 历史（--amend 选项表示修改上一个 commit 而不是创建一个新的 commit，commit 被修改过后，git push 必须加 -f 强制推送才能 push 成功）。
    - 不同意的，讨论。
    - 完成后，assign 给审查者再次 review（不需要额外留评论）；如果之前并未 assign 回来，那么留评论 fixed （只需在整个 MR 下留一个）通知审查者已修完。回到第二步。
    - 注意，来回 assign 或留评论都是沟通机制，二选一，没道理都做。
    - 审查者确认没有问题之后，将 MR assign 给目标分支的维护者进行二次 review 或合并。
