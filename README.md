# 网站自动登录

## 实现

使用selenium+webdriver实现

go语言作为基本的实现，python作为验证码接口识别的实现

## 使用说明
按照confile.yaml中对应的格式填写即可
- target_url 填写目标网站的登录页面
- browser_path 填写chrome浏览器driver的路径,请尽量和本地的chrome版本保持一致,地址为`https://registry.npmmirror.com/binary.html?path=chromedriver/`
- headless 是否使用无头模式(也就是是否显示的显示窗口),默认为false
- captchaon 是否开启验证码识别,默认为false
- captcha_serverurl 验证码识别服务器的地址,对应脚本为`ddddocr.py`
- userinput_jspath 用户名输入框的xpath,可以使用chrome的开发者工具查看,直接右击复制即可
- passinput_xpath 密码输入框的xpath,可以使用chrome的开发者工具查看,直接右击复制即可
- captchainput_xpath 验证码输入框的xpath,可以使用chrome的开发者工具查看,直接右击复制即可
- captchaimg_xpath 验证码图片的xpath,可以使用chrome的开发者工具查看,直接右击复制即可
- loginbutton_xpath 登录按钮的xpath,可以使用chrome的开发者工具查看,直接右击复制即可
- body_exclude_regex 用于判断是否登录成功的字符(尽量填写登录页面的,请勿填写请求数据包的响应),如果匹配则表示登录成功,否则表示登录失败
- timeinterval_ms 每次登录后的等待时间,单位为毫秒
- proxy 代理服务器地址,如果不需要代理则留空

请git clone 本项目后使用`go build`进行编译,然后使用`testgo`运行,

初次使用可先使用项目本身的配置进行测试

## 更新日志

### 2023年4月27日09:46:12

1. 解决了代码上的一些问题导致的错误
2. 修改实现逻辑