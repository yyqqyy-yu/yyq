package dao

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Orm struct {
	*gorm.DB
}
var DB *gorm.DB
var err error
type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}
func Minit(config string)*Orm {
	if err := Init(config); err != nil {
		panic(err)
	}

	user := viper.GetString("username")
	password := viper.GetString("password")
	addr := viper.GetString("addr")
	dbName := viper.GetString("dbName")
	charset := viper.GetString("charset")
	parseTime := viper.GetString("parseTime")
	loc := viper.GetString("loc")
	fmt.Println("Viper get name:", user, user+":"+password+"@"+"("+addr+")"+"/"+dbName+"?charset="+charset+"&parseTime="+parseTime+"&loc="+loc)
	DB,err= gorm.Open("mysql", user+":"+password+"@"+"("+addr+")"+"/"+dbName+"?charset="+charset+"&parseTime="+parseTime+"&loc="+loc)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return &Orm{DB}


}

/*type Config struct {
	username    string
	password    string
	dsn         string
	DBName      string
	Host        string
	Port        string
	DbCharset   string
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
}

func (c *Config) SetPassword(password string) {
	c.password = password
}

func (c *Config) SetUsername(username string) {
	c.username = username
}

func (c *Config) DSN() string {
	return c.dsn
}

func (c *Config) Init() *Config {
	// check
	if c.DBName == "" || c.Host == "" || c.Port == "" {
		panic(fmt.Sprintf("必要参数缺失, 请检查!  db name: %s, db host: %s, db port: %s", c.DBName, c.Host, c.Port))
	}

	if c.username == "" {
		c.username = "root"
		userNameData, err := ioutil.ReadFile("/usr/local/.db/mysql.uname")
		if err == nil {
			c.username = strings.TrimSpace(string(userNameData))
		}
	}

	if c.password == "" {
		c.password = "root"
		pwdData, err := ioutil.ReadFile("/usr/local/.db/mysql.pas")
		if err == nil {
			c.password = strings.TrimSpace(string(pwdData))
		}
	}

	if c.DbCharset == "" {
		c.DbCharset = "utf8mb4,utf8"
	}

	if c.Active == 0 {
		c.Active = 20
	}

	if c.Idle == 0 {
		c.Idle = 10
	}

	if c.IdleTimeout == time.Duration(0*xtime.Second) {
		c.IdleTimeout = 4 * time.Duration(4*xtime.Hour)
	}

	// todo maybe need add timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local
	c.dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.username, c.password, c.Host, c.Port, c.DBName, c.DbCharset)

	return c
}

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	log.Logger.Sugar().Info(fmt.Sprintf(strings.Repeat("%v ", len(v)), v...))
}

type Orm struct {
	*gorm.DB
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) *Orm {
	c.Init()
	var db *gorm.DB
	db, err := gorm.Open("mysql", c.DSN())
	if err != nil {
		log.Logger.Sugar().Info(fmt.Sprintf("db dsn(%s) error: %v", c.DSN(), err))
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(xtime.Duration(c.IdleTimeout) / xtime.Second)
	db.SetLogger(ormLog{})

	return &Orm{
		DB: db,
	}
}*/
