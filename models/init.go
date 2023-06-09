package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据连接
func init() {

	beego.LoadAppConfig("ini", "../conf/app.conf")
	//读取配置文件，设置数据库参数
	//数据库类别
	dbType := beego.AppConfig.String("db_type")
	//连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")

	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")

	//数据库IP（域名）
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	//数据库端口
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	switch dbType {
	case "sqlite3":
		orm.RegisterDataBase(dbAlias, dbType, dbName)
	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "::db_charset")

		orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
			dbPort+")/"+dbName+"?charset="+dbCharset, 200, 30)

	}
	//注册model
	initRegister()
	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	isDev = true
	//自动建表
	orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}

func initRegister() {
	orm.RegisterModel(
		new(Inscript),
		new(Ens),
	)
}

// TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("mysql::db_prefix")
	return prefix + name
}

func InscriptTBName() string {
	return TableName("inscript")
}

func EnsTBName() string {
	return TableName("ens")
}
