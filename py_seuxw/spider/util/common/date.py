# -*- coding: utf-8 -*-

import time
import datetime

class Time(object):

    def __init__(self):
        pass

    @staticmethod
    def now_date_str():
        '''返回形如 20180301 的日期字符串'''
        return time.strftime("%Y%m%d", time.localtime())

    @staticmethod
    def now_time_str():
        '''返回形如 120102 的时间字符串'''
        return time.strftime("%H%M%S", time.localtime())

    @staticmethod
    def now_str():
        '''返回形如 20180101_120102 的时间字符串'''
        return time.strftime("%Y%m%d_%H%M%S", time.localtime())

    @staticmethod
    def now_datetime_str():
        '''返回形如 20180101_120102 的时间字符串'''
        return Time.now_str()

    @staticmethod
    def ISO_str():
        '''返回标准时间字符串 2018-12-12 12:13:14'''
        return time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())

    @staticmethod
    def ISO_datetime_str():
        '''返回标准时间字符串 2018-12-12 12:13:14'''
        return Time.ISO_str()

    @staticmethod
    def ISO_date_str():
        '''返回标准日期字符串 2018-12-12'''
        return time.strftime("%Y-%m-%d", time.localtime())

    @staticmethod
    def ISO_time_str():
        '''返回标准时间字符串 12:13:14'''
        return time.strftime("%H:%M:%S", time.localtime())

class DateTime(object):

    @staticmethod
    def str_to_datetime(date_str):
        try:
            str_timestamp = time.mktime(time.strptime(date_str, "%Y-%m-%d"))
        except Exception:
            str_timestamp = time.mktime(time.strptime(date_str, "%Y-%m-%d %H:%M:%S"))
        return datetime.datetime.fromtimestamp(str_timestamp)

    @property
    def date_str(self):
        return self.ISO_date_str.replace('-', '')

    @property
    def ISO_date_str(self):
        return str(self.t.date())

    def __init__(self, date):
        if isinstance(date, (str)):
            self.t = DateTime.str_to_datetime(date)
        elif isinstance(date, (datetime.datetime)):
            self.t = date
        else:
            raise TypeError("类型错误，仅接收 str/datetime 类型的时间数据")
    
    def __sub__(self, data):
        '''重载 - 计算时间差/减去指定时间'''
        if isinstance(data, DateTime):
            return (self.t-data.t).days
        elif isinstance(data, datetime.datetime):
            return (self.t-data).days
        elif isinstance(data, str):
            return (self.t - DateTime(data).t).days
        elif isinstance(data, int):
            return self+(-data)
        else:
            raise TypeError("类型错误，仅接收 DateTime/str/datetime/int 类型的时间数据")

    def __add__(self, data):
        '''重载 + '''
        if isinstance(data, int):
            self.add_days(data)
            return self
        else:
            raise TypeError("类型错误，仅接收 int 类型的添加时间数据")

    def __lt__(self, data):
        '''重载 <'''
        if isinstance(data, DateTime):
            if self - data < 0:
                return True
            else:
                return False
        else:
            raise TypeError("类型错误，仅接收 DateTime 类型的时间数据")

    def __gt__(self, data):
        '''重载 >'''
        if isinstance(data, DateTime):
            if self - data > 0:
                return True
            else:
                return False
        else:
            raise TypeError("类型错误，仅接收 DateTime 类型的时间数据")

    def __le__(self, data):
        '''重载 <='''
        if isinstance(data, DateTime):
            if self - data <= 0:
                return True
            else:
                return False
        else:
            raise TypeError("类型错误，仅接收 DateTime 类型的时间数据")

    def __ge__(self, data):
        '''重载 >='''
        if isinstance(data, DateTime):
            if self - data >= 0:
                return True
            else:
                return False
        else:
            raise TypeError("类型错误，仅接收 DateTime 类型的时间数据")
    
    def __eq__(self, data):
        '''重载 =='''
        if isinstance(data, DateTime):
            if self - data == 0:
                return True
            else:
                return False
        else:
            raise TypeError("类型错误，仅接收 DateTime 类型的时间数据")

    def __ne__(self, data):
        '''重载 !='''
        if isinstance(data, DateTime):
            if self - data != 0:
                return True
            else:
                return False
        else:
            raise TypeError("类型错误，仅接收 DateTime 类型的时间数据")


    def add_days(self, days):
        self.t = self.t + datetime.timedelta(days=days)

if __name__ == "__main__":
    print (Time.now_date_str())
    print (Time.now_datetime_str())