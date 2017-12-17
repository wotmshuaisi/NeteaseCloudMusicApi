# NeteaseCloudMusicApi
golang版网易云音乐api，请大家多多指教

# 编译和运行环境说明
    docker 1.9.1
    golang 1.9
    ubuntu 14.04
    1)安装包管理
            go get -u github.com/golang/dep/cmd/dep
            dep init
    2)运行脚本 ./newBuildDocker 可以一键编译并运行

# 自动化生成文档
    配置文件中设置：EnableDocs = true
    bee run -gendoc=true -downdoc=true

# 文件外挂说明
    配置文件 /neteasecloudmusicapi/webData/conf
    日志文件 /neteasecloudmusicapi/log

#demo
    http://haibarai.com/music

# 参照了以下前辈的代码
    https://github.com/yitimo/api-163-go/blob
    https://github.com/Binaryify/NeteaseCloudMusicApi
    http://moonlib.com/606.html#more-606