/*
	Copyright 2017 by GoWeb author: gdccmcm14@live.com.
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License
*/
package home

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare() {
	this.Lang = ""

	l := this.Ctx.GetCookie("lang")
	if l == "" {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5]
			if i18n.IsExist(al) {
				this.Lang = al
			}
		}
	} else {
		switch l {
		case "en":
			this.Lang = "en-US"
		default:
			this.Lang = "zh-CN"

		}
	}
	if this.Lang == "" {
		this.Lang = "zh-CN"
	}

	this.Data["Lang"] = this.Lang
}

// Get the address of template
func (this *baseController) GetTemplate() string {
	templatetype := beego.AppConfig.String("home_template")
	if templatetype == "" {
		templatetype = "default"
	}
	return templatetype
}

func (this *baseController) Rsp(status bool, str string) {
	if status {
		this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
		this.ServeJSON()
	}
	this.Ctx.WriteString(`
	<html>
<head>
    <title>404-帮帮宝贝回家</title>
    <script>
        document.body.classList.add('page-fullscreen');
    </script>
    <script type="text/javascript" src="http://www.qq.com/404/search_children.js" charset="utf-8" homePageUrl="https://www.github.com/hunterhug" homePageName="更多精彩：https://www.github.com/hunterhug"></script>
</head>
</html>
	`)
	this.StopRun()
}
