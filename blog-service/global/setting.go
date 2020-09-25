package global

import(
	"github.com/go-projects/blog-service/pkg/setting"
	"github.com/go-projects/blog-service/pkg/logger"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger *logger.Logger
)