
echo "开始创建数据库"
goctl model mysql datasource -url="xu756:123456@tcp(127.0.0.1:3306)/mylooklook" -table="user" -c -dir .