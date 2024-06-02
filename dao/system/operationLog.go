package system

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"new-go-edas-admin/common/global"
	"new-go-edas-admin/models/system"
	"new-go-edas-admin/utils"
	"time"
)

type OperationLogService interface {
	SaveOperationLogChannel(olc <-chan *system.OperationLog)
	GetOperationLogList(limit, page int) (*system.OperationLogList, error)
}

type operationLogService struct{}

func NewOperationLogService() OperationLogService {
	return &operationLogService{}
}

// 处理OperationLogChan将日志记录到数据库
func (s *operationLogService) SaveOperationLogChannel(olc <-chan *system.OperationLog) {
	var url = viper.GetString("IpLocation.siteURL")
	var header = map[string]string{}
	Logs := make([]system.OperationLog, 0)
	//5s 自动同步一次
	duration := 5 * time.Second
	timer := time.NewTimer(duration)
	defer timer.Stop()
	for {
		select {
		case log := <-olc:
			var ipLocation system.IPLocation
			fullURL := fmt.Sprintf("%s?ip=%s", url, log.Ip)
			data, err := utils.DoRequest("GET", fullURL, header, nil)
			if err != nil {
				log.IpLocation = "查找失败"
			}
			_ = json.Unmarshal([]byte(data), &ipLocation)
			if ipLocation.Data.City == "" {
				log.IpLocation = ipLocation.Data.Continent
			} else {
				log.IpLocation = ipLocation.Data.City
			}
			Logs = append(Logs, *log)
			//每五条记录到数据库
			if len(Logs) >= 5 {
				global.GORM.Create(&Logs)
				Logs = make([]system.OperationLog, 0)
				timer.Reset(duration)
			}
		case <-timer.C:
			if len(Logs) > 0 {
				global.GORM.Create(&Logs)
				Logs = make([]system.OperationLog, 0)
			}
			timer.Reset(duration)
		}
	}
}

// 获取操作日志列表
func (s *operationLogService) GetOperationLogList(limit, page int) (*system.OperationLogList, error) {
	//定义分页起始位置
	startSet := (page - 1) * limit
	var (
		operationLogList []system.OperationLog
		total            int64
	)
	if err := global.GORM.Model(&system.OperationLog{}).Count(&total).Limit(limit).Offset(startSet).
		Order("start_time desc").Find(&operationLogList).Error; err != nil {
		return nil, err
	}
	return &system.OperationLogList{
		Items: operationLogList,
		Total: total,
	}, nil
}
