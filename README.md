# ncmlu

NCM 任务执行脚本，Go 版本。

## Implements

- ✅ 登录 
- ✅ 签到
- ✅ 获取推荐歌单
- ✅ 获取歌单内所有歌曲
- ✅ 刷播放量

## Usage

1. 安装 Go 环境

2. 执行`go build`得到可执行文件

3. 将可执行文件（如*ncmlu.exe*）与*config.yaml*文件放在同一目录下

4. 按如下方式配置*config.yaml*文件：

   单账号：

   ```
   accounts:
     - phone: 1111111111 	// 修改为账号的手机号
       passwd: xxxxxxxxx	// 修改为对应的密码
       expired: 2021-09-05	// 到期时间，如设置为2021-09-05则当天及之后不会再执行该账号的任务
       only_sign: false	// 是否只执行签到，设置为true则仅执行签到任务
   ```

   多账号（规则与单账号相同）：

   ```
   accounts:
     - phone: 1111111111
       passwd: xxxxxxxxx
       expired: 2021-09-05
       only_sign: false
     - phone: 1111111111
       passwd: xxxxxxxxx
       expired: 2021-09-05
       only_sign: false
     - phone: 1111111111
       passwd: xxxxxxxxx
       expired: 2021-09-05
       only_sign: false
     - phone: 1111111111
       passwd: xxxxxxxxx
       expired: 2021-09-05
       only_sign: false
   ```

5.   双击可执行文件（如*ncmlu.exe*）执行脚本
6.   查看当前目录下的新文件*ncmlu.log*，可以得到输出结果

## TODO

- 路由
- 日志文件