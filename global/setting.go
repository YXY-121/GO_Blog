package global

import (
	"awesomeProject/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine *gorm.DB
)
