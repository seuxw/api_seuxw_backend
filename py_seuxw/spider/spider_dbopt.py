#!/usr/bin/python3.5
# -*- coding:utf-8 -*-
# 预运行工具 - 数据库操作工具

from util.database import DBController
from util.redis import RedisController
from util.common.date import Time, DateTime
from util.config import ConfigReader
from util.common.logger import use_logger

import sys
import time
import datetime

@use_logger(level="info")
def db_optor_info(msg):
    pass

if __name__ == "__main__":
    pass