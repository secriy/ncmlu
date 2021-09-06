# ncmlu

NCM 任务执行脚本，Go 版本。

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/secriy/ncmlu/Go)
![GitHub](https://img.shields.io/github/license/secriy/ncmlu)

## Implements

- ✅ 登录 
- ✅ 签到
- ✅ 获取推荐歌单
- ✅ 获取歌单内所有歌曲
- ✅ 刷播放量
- ✅ 日志文件

## Usage

1. 跳转 [Release](https://github.com/secriy/ncmlu/releases) 页面下载对应操作系统的压缩包，如 Windows 环境下载 ncmlu_xxx_Windows_x86_64.tar.gz
2. 解压压缩包，得到可执行文件，如 Windows 下的*ncmlu.exe*文件
3. 创建配置文件*config.yaml*
4. 将可执行文件（如*ncmlu.exe*）与*config.yaml*文件放在同一目录下
5. 按如下方式配置*config.yaml*文件：

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

6.   双击可执行文件（如*ncmlu.exe*）执行脚本
7.   查看当前目录下的新文件*ncmlu.log*，可以得到输出结果

## TODO

- 路由
