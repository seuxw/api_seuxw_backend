# -*- coding: utf-8 -*-
# 总调度器 + 多线程控制器
from do.test import Test
from util.config import ConfigParser
from constant.logger import *

import time

class Do(Test):
    
    @staticmethod
    def do(use_normal=True):
        '''默认运行方法
        - use_normal:使用默认值执行程序
        '''
        pass