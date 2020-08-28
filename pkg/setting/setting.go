package setting

import (
	"github.com/prometheus/common/log"
	"gopkg.in/ini.v1"
	"time"
)

//app配置
type App struct {
	JwtSecret       string        //jwt密钥
	JwtExpireTime   time.Duration //jwt失效时间，单位是小时
	PageSize        int           //分页返回的数据个数
	RuntimeRootPath string        //保存文件的跟路径

	ImagePrefixUrl string   //图片url
	ImageSavePath  string   //图片要保存的路径
	ImageMaxSize   int      //图片最大尺寸
	ImageAllowExts []string //图片允许的格式   jpg jpeg png

	ApkSavePath string //apk文件路径
	ApkAllowExt string //apk文件格式
	AppStoreUrl string //iOS应用在App Store中的地址，用于版本更新

	LogSavePath string //日志文件保存的路径
	LogSaveName string //日志文件名称
	LogFileExt  string //日志文件后缀
	TimeFormat  string //文件的日期名称

	WechatAppID  string //微信的appID
	WechatSecret string //微信的secret
	QQAppID      string //QQ的appID
	QQAppKey     string //QQ的appkey
}

var AppSetting = &App{}

//服务配置
type Server struct {
	RunMode      string        //运行模式
	HttpPort     int           //端口号
	ReadTimeout  time.Duration //读取超时时间
	WriteTimeout time.Duration //写入超时时间
}

var ServerSetting = &Server{}

//数据库配置
type Database struct {
	Type        string //数据库类型
	User        string //数据库用户
	Password    string //数据库密码
	Host        string //数据库地址+端口号
	Name        string //数据库名称
	TablePrefix string //数据库数据表前缀
}

var DatabaseSetting = &Database{}

//es配置
type Elasticsearch struct {
	Host                   string        //连接地址
	SetHealthcheckInterval time.Duration //
	SetGzip                bool          //
}

var ElasticsearchSetting = &Elasticsearch{}

func SetUp() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("获取app.ini配置失败")
	}

	//映射配置       section为取出ini的数组   mapto映射到某个结构体中
	err = cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("cfg配置文件映射 AppSetting 错误:%v", err)
	}
	//设置允许上传的图片最大尺寸
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("cfg配置文件映射 ServerSetting 错误:%v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("cfg配置文件映射 DatabaseSetting 错误:%v", err)
	}

	err = cfg.Section("elasticsearch").MapTo(ElasticsearchSetting)
	if err != nil {
		log.Fatalf("cfg配置文件映射 ElasticsearchSetting 错误:%v", err)
	}
}
