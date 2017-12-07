# NeteaseCloudMusicApi
golang版网易云音乐api，使用beego api框架搭建

# 编译和运行环境说明
    docker 17.04.0-ce
    golang 1.9
    ubuntu 14.04
    运行脚本 ./newBuildDocker 可以一键编译并运行

# 自动化生成文档
    配置文件中设置：EnableDocs = true
    bee run -gendoc=true -downdoc=true

# 文件外挂说明
    配置文件 /neteasecloudmusicapi/webData/conf
    日志文件 /neteasecloudmusicapi/log

# 参照了以下前辈的代码
    https://github.com/yitimo/api-163-go/blob
    https://github.com/Binaryify/NeteaseCloudMusicApi