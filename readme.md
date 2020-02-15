# 说明

1. cptool -h
Usage of cptool:
  -e int
        网关服务结束端口
  -h    帮助
  -o int
        需要替换的网关端口
  -s int
        网关服务起始端口
  -src string
        源文件夹
        
2. 启动
cptool -src=./GateServer/ -s=8501 -e=8570 -o=8500