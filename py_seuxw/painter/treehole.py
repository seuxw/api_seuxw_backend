# -*- coding: utf-8 -*-
# 树洞图片生成基础模块
from util.img.img_blender import ImgBlender
from util.img.img_selector import ImgSelector

def exec():
    contents = {
        'title':'小微树洞·悄悄话',
        'content_send':'小微你好，我本来喜欢一个女生。但那个女生的性格特别的man，一点都不温柔可爱，也不打算找男朋友。我该怎么办？',
        'sender':'测试人员',
        'date_send':'2018-01-02',
        'content_reply':"你可以考虑和她当哥们呀～",
        'replier':"测试人员2",
        'date_reply':'2018-03-21'
    }
    img = ImgBlender(ImgSelector().img, contents)
    img.blend_type_treehole()

if __name__ == "__main__":
    exec()