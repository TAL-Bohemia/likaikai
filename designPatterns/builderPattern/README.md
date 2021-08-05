# 建造者模式

## 概念
    使用多个简单的对象一步一步构建成一个复杂的对象

## 介绍
* 何时使用： 基本组件不变，而其组合经常变化的时候。
* 应用实例： 去肯德基，汉堡、可乐、薯条、炸鸡翅等是不变的，而其组合是经常变化的，生成出所谓的"套餐"。
* 优点：
  * 1、建造者独立，易扩展。
  * 2、便于控制细节风险。
* 缺点：
  * 1、产品必须有共同点，应用范围有限。
  * 2、如内部变化复杂，会有很多的建造类。