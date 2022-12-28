#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys
import json

try:
    form_data = json.loads(sys.argv[1])
except Exception as e:
    print("waring: ", e)


# ------------- 在下面编写您的业务逻辑代码 -------------

# 格式如下
# dict = {"id": 13, "title": "test", "priority": 1, "form_data": [{"project_name": "Nginx"}]}
# 将form_data中的数据写入到文件中
with open("form_data.txt", "w") as f:
    f.write(json.dumps(form_data["form_data"][0]["project_name"])) #form_data["form_data"][0]["project_name"] 可以跟返回值如"Nginx"进行jenkins发布的交互

