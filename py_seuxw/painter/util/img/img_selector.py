import sys
sys.path.append("../..")

from constant.list import TREEHOLE_TYPE_LIST
from constant.value import TEMPLATE_BASE_PATH
import os

class ImgSelector():

    def __init__(self, treehole_type="treehole"):
        self.treehole_type = treehole_type
        self.path = "%s/%s"%(TEMPLATE_BASE_PATH, treehole_type)
        self.__select_img__
    
    @property
    def __select_img__(self):
        from random import randint

        img_list = list()
        for root, dirs, files in os.walk(self.path):
            for file in files:
                img_list.append(os.path.abspath(os.path.join(os.getcwd(), root, file)))
        try:
            idx = randint(0, len(img_list)-1)
            self.img = img_list[idx]
        except Exception:
            raise Exception

if __name__ == "__main__":
    img = ImgSelector().img
    print(img)