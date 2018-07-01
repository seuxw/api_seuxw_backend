# -*- coding: utf-8 -*-
# 图片合成器
from PIL import Image, ImageFont, ImageDraw
from io import BytesIO

def set_paragraph(text, count):
    text = "\t\t"+text
    rtn_text = str()
    for idx in range(0, int(len(text)/count)+1):
        try:
            rtn_text += text[idx*(count):(idx+1)*count]
            rtn_text += '\n'
        except Exception:
            rtn_text += text[idx*(count-1):]
    return rtn_text

class ImgBlender():

    def __init__(self, img_path, contents):
        '''
        img_path : 背景图片地址
        contents : 指定类型的文本字典
        '''
        self.img = Image.open(img_path)             # 图片对象
        self.img_draw = ImageDraw.Draw(self.img)    # 绘制图片对象
        
        self.font_dict = dict()
        self.posi_dict = dict()
        self.fill_dict = dict()
        
        self.contents = contents
        self.__set_default_config__

    @property
    def __set_default_config__(self):
        # 文本换行设置
        self.contents['content_send'] = set_paragraph(self.contents['content_send'], 15)
        self.contents['content_reply'] = set_paragraph(self.contents['content_reply'], 15)

        # 字体字典配置
        self.font_dict["title"] = ImageFont.truetype(font='./bin/ttf/weibeijian.ttf', size=100, encoding="utf-8")
        self.font_dict["content_reply"] = self.font_dict["content_send"] = ImageFont.truetype(font='./bin/ttf/xihei.ttf', size=75, encoding='utf-8')
        self.font_dict["sender"] = self.font_dict["replier"] = self.font_dict["date_send"] = self.font_dict["date_reply"] = ImageFont.truetype(font='./bin/ttf/xihei.ttf', size=60, encoding='utf-8')

        # 文字位置配置
        self.posi_dict["title"] = (150,100)
        self.posi_dict["content_send"] = (150,270)
        self.posi_dict["sender"] = (950,700)
        self.posi_dict["date_send"] = (950,800)
        self.posi_dict["content_reply"] = (150,1000)
        self.posi_dict["replier"] = (950,2000)
        self.posi_dict["date_reply"] = (950,2100)

        # 文字颜色配置
        self.fill_dict["title"] = "#ffffff"
        self.fill_dict["content_send"] = (80,80,80)
        self.fill_dict["sender"] = (80,80,80)
        self.fill_dict["date_send"] = (80,80,80)
        self.fill_dict["content_reply"] = (20,80,30)
        self.fill_dict["replier"] = (80,80,80)
        self.fill_dict["date_reply"] = (80,80,80)

    def blend_type_treehole(self):
        for k in self.contents.keys():
            self.img_draw.text(self.posi_dict[k], self.contents[k], font=self.font_dict[k], fill=self.fill_dict[k])
        
        self.img.save("./output/treehole.jpg")

if __name__ == "__main__":
    pass
    