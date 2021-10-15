# ncmlu

NCM 任务执行脚本，Go 版本。

希望能够帮忙点一个 Star ^\_^

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/secriy/ncmlu/Go)
![GitHub](https://img.shields.io/github/license/secriy/ncmlu)

## TODOs

- ✅ 登录
- ✅ 签到
- ✅ 获取推荐歌单
- ✅ 获取歌单内所有歌曲
- ✅ 刷播放量
- ✅ 日志文件
- ✅ 间隔执行
- ✅ 自定义歌单

## Usage

1. 跳转 [Release](https://github.com/secriy/ncmlu/releases) 页面下载对应操作系统的压缩包，如 Windows 环境下载 ncmlu_xxx_Windows_x86_64.tar.gz
2. 解压压缩包，得到可执行文件，如 Windows 下的*ncmlu.exe*文件
3. 创建配置文件*config.yaml*
4. 将可执行文件（如*ncmlu.exe*）与*config.yaml*文件放在同一目录下
5. 按如下方式配置*config.yaml*文件，注意空格和缩进：

   单账号：

   ```yaml
   accounts:
     - phone: 1111111111 	# 修改为账号的手机号
       passwd: 'xxxxxxxxx'	# 修改为对应的密码，为防止解析错误，建议使用半角引号包裹；密码支持 32 位小写 MD5 格式，同样支持明文
       code: 86            # 国家码，如省略本字段则默认使用中国的 86 
       expired: 2021-09-05	# 到期时间，如设置为 2021-09-05 则当天及之后不会再执行该账号的任务
       only_sign: false	# 是否只执行签到，设置为 true 则仅执行签到任务（省略则为 false）
       unstable: false     # 不稳定选项，设置为 true 则可以使用非推荐歌单刷歌，增加刷歌成功数量（可能导致日推风格改变，省略则为 false）
   ```

   多账号（规则与单账号相同）：

   ```yaml
   accounts:
     - phone: 1111111111
       passwd: 'xxxxxxxxx'
       expired: 2021-09-05
     - phone: 1111111111
       passwd: 'xxxxxxxxx'
       expired: 2021-09-05
       only_sign: true
       unstable: true
     - phone: 1111111111
       passwd: 'xxxxxxxxx'
       expired: 2021-09-05
       only_sign: true
     - phone: 1111111111
       passwd: 'xxxxxxxxx'
       expired: 2021-09-05
       unstable: true
   ```

7. 双击可执行文件（如*ncmlu.exe*）执行脚本
8. 查看当前目录下的新文件*ncmlu.log*，可以得到输出结果

### 自定义歌单

注意：使用自定义歌单时，不会获取每日的推荐歌单刷歌，如需取消自定义歌单，删除或注释增加的**playlist**字段即可。

将*config.yaml*配置文件修改为如下格式，多账号同理：

```yaml
playlist:
  - 4234112
  - 4312424
accounts:
  - phone: 1342412432
    passwd: "xxxxxx"
    expired: 2022-09-06
    only_sign: false
```

**playlist**列表填写需要使用的歌单 ID，可指定单个或多个，可以通过下图方式获得：

![image-20210906181506108](README/image-20210906181506108.png)

### 设置睡眠和休眠

由于大量账号的快速执行会导致 IP 被封禁、频繁访问等后果，因此脚本提供自定义执行间隔（interval），以及短时睡眠（catnap）和长时间睡眠（sleep）的选项。

- interval: 每个账号之间的执行间隔，设置为正整数有效（单位：秒）
- catnap: 短期睡眠
    - number: 每执行 number 个账号后就进行睡眠
    - duration: 短时间睡眠的时间（单位：分钟）
- sleep: 长时间睡眠
    - number: 每执行 number 个账号后就进行睡眠
    - duration: 长时间睡眠的时间（单位：分钟）

如不需要睡眠省略字段即可，或者将 number 或 duration 设置为 0 或负数，就不会应用睡眠。

示例：

```yaml
interval: 3 # 单位：秒
catnap:
  number: 20
  duration: 2 # 单位：分钟
sleep:
  number: 500
  duration: 30 # 单位：分钟
accounts:
  - phone: 1342412432
    passwd: "xxxxxx"
    code: 86
    expired: 2022-09-06
    only_sign: false
    unstable: false
```

## Deployment

### Linux 服务器部署

1. 从 [Release](https://github.com/secriy/ncmlu/releases) 页面下载*ncmlu_xxx_Linux_x86_64.tar.gz*压缩包，
2. 将 tar.gz 压缩包中的*ncmlu*可执行文件提取至服务器某一目录，再将项目中的*config.yaml*文件复制到同一目录（也可手动创建文件）
3. 按照前文所述规则填写*config.yaml*文件配置
4. 在同一目录下创建*run.sh*文件，填写如下内容：

   ```sh
   # 这里的 /path_to 需要改为ncmlu文件所在的目录的绝对路径
   cd /path_to
   ./ncmlu
   ```

   如：

   ```sh
   # 该路径仅为举例，请勿直接用于自己的服务器
   cd /home/secriy/task
   ./ncmlu
   ```

5. 输入命令`crontab -e`，在打开的编辑器中填写一行内容（/path_to 同样需要更改）：

   ```
   0 2 * * * bash /path_to/run.sh
   ```

   其中，`0 2 * * *`表示每天的凌晨 2 点 0 分执行，如有需要可以修改，例如：

   ```
   30 5 * * * bash /home/secriy/task/run.sh
   ```

   即每天的 5 点 30 分执行。

6. 保存修改

### Linux 下宝塔面板部署

如果有宝塔面板，首先按照上一条**Linux 服务器部署**的前四步操作，最后在宝塔面板的定时任务里创建一个定时任务，执行的指令填写`bash /path_to/run.sh`。

## 免责声明

本项目仅用于个人用途，完全开源。本人不提供代挂业务，未使用该项目盈利。请勿向他人提供个人账号、密码（包括密码的 MD5 值），向不受信任的第三方提供个人信息导致的问题皆与本人无关。

本项目使用 MIT 开源协议，使用本项目进行盈利、二次开发等导致的风险及法律问题与本项目无关。

所有网易云音乐相关字样版权皆属于网易公司，勿用于商业及非法用途，如产生法律纠纷与本项目无关。

如果该项目侵犯了您的权益，请通过邮箱 secriyal@gmail.com 联系本人及时处理，我们会第一时间为您处理。