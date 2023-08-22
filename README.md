# StarRocks License

## Support

```
yaml配置文件

starrocks:
  #集群名称: [manager登录地址,登录账号,登录密码]
  cluster1: [http://ip:port:19321,user,password]
  cluster2: [http://ip:port:19321,user,password]
  cluster3: [http://ip:port:19321,user,password]
  cluster4: [http://ip:port:19321,user,password]

```

```
根据配置文件信息，获取license过期时间


2023-08-21T23:52:30.129+0800	info	log format is success.
2023-08-21T23:52:30.129+0800	info	license 检测
2023-08-21T23:52:30.372+0800	info	{"code":20000,"data":{"roles":["Node_priv","Admin_priv","","(false)"]}}
2023-08-21T23:52:30.401+0800	info	{"code":20000,"data":{"roles":["NODE","ADMIN"]}}
2023-08-21T23:52:30.475+0800	info	{"code":20000,"data":{"roles":["Node_priv","Admin_priv","","(false)"]}}
2023-08-21T23:52:30.541+0800	info	{"code":20000,"data":{"roles":["Node_priv","Admin_priv"]}}
2023-08-21T23:52:30.611+0800	info	{"code":20000,"list":[{"cores":288,"expire_at":1694102400000,"hosts":9}],"total":1}
2023-08-21T23:52:30.611+0800	info	现在日期:1692633150611,到期日期:1694102400000
2023-08-21T23:52:30.612+0800	info	集群：cluster5,核数：288,license过期时间：2023-09-08 00:00:00,节点数：9, license 18天后将过期!
2023-08-21T23:52:30.612+0800	info	cluster5 现在day的值是：18
2023-08-21T23:52:30.661+0800	info	{"code":20000,"list":[{"cores":120,"expire_at":1694102400000,"hosts":9}],"total":1}
2023-08-21T23:52:30.661+0800	info	现在日期:1692633150661,到期日期:1694102400000
2023-08-21T23:52:30.661+0800	info	集群：cluster6,核数：120,license过期时间：2023-09-08 00:00:00,节点数：9, license 18天后将过期!
2023-08-21T23:52:30.661+0800	info	cluster6 现在day的值是：18
2023-08-21T23:52:30.737+0800	info	{"code":20000,"data":{"roles":["Node_priv","Admin_priv","","(false)"]}}
2023-08-21T23:52:30.822+0800	info	{"code":20000,"list":[{"cores":304,"expire_at":1694102400000,"hosts":11}],"total":1}
2023-08-21T23:52:30.822+0800	info	现在日期:1692633150822,到期日期:1694102400000
2023-08-21T23:52:30.822+0800	info	集群：cluster3,核数：304,license过期时间：2023-09-08 00:00:00,节点数：11, license 18天后将过期!
2023-08-21T23:52:30.822+0800	info	cluster3 现在day的值是：18
2023-08-21T23:52:30.938+0800	info	{"code":20000,"list":[{"cores":480,"expire_at":1694102400000,"hosts":15}],"total":1}
2023-08-21T23:52:30.938+0800	info	现在日期:1692633150938,到期日期:1694102400000
2023-08-21T23:52:30.938+0800	info	集群：cluster4,核数：480,license过期时间：2023-09-08 00:00:00,节点数：15, license 18天后将过期!
2023-08-21T23:52:30.939+0800	info	cluster4 现在day的值是：18
2023-08-21T23:52:31.220+0800	info	{"code":20000,"list":[{"cores":576,"expire_at":1694102400000,"hosts":9}],"total":1}
2023-08-21T23:52:31.220+0800	info	现在日期:1692633151220,到期日期:1694102400000
2023-08-21T23:52:31.220+0800	info	集群：cluster2,核数：576,license过期时间：2023-09-08 00:00:00,节点数：9, license 18天后将过期!
2023-08-21T23:52:31.220+0800	info	cluster2 现在day的值是：18
2023-08-21T23:52:31.365+0800	info	{"code":20000,"data":{"roles":["Node_priv","Admin_priv","","(false)"]}}
2023-08-21T23:52:32.537+0800	info	{"code":20000,"list":[{"cores":1336,"expire_at":1694102400000,"hosts":22},{"cores":1336,"expire_at":1694102400000,"hosts":22}],"total":2}
2023-08-21T23:52:32.537+0800	info	现在日期:1692633152537,到期日期:1694102400000
2023-08-21T23:52:32.537+0800	info	集群：cluster1,核数：1336,license过期时间：2023-09-08 00:00:00,节点数：22, license 18天后将过期!
2023-08-21T23:52:32.537+0800	info	cluster1 现在day的值是：18
2023-08-21T23:52:32.537+0800	info	done!
```

## Authors

* **Author**  - **_ChengKen_**

  ###### **The wind blows away the thoughts, rolled up unruly time**


## Best Regards

* Hat tip to anyone whose code was used
* Inspiration
* etc
