package service

import (
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
	"time"
)

type DeviceService struct{}

type DeviceOnlineMsg struct {
	Mac    string `json:"mac"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

func (ds *DeviceService) Ping(key string) error {
	d := query.Device
	_, err := d.Where(d.Key.Eq(key)).
		Update(d.OnlineTime, int32(time.Now().Unix()))
	return err
}

func (ds *DeviceService) Online(msg DeviceOnlineMsg) (*model.Device, error) {

	d := query.Device

	// 查询设备
	device, err := d.Where(
		d.Key.Eq(msg.Key),
		d.Secret.Eq(msg.Secret),
	).First()
	if err != nil {
		return nil, err
	}

	// 修改设备为在线
	device.IsOnline = 1
	device.OnlineTime = int32(time.Now().Unix())
	_, err = d.Where(d.ID.Eq(device.ID)).Updates(device)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (ds *DeviceService) Offline(key string) error {
	d := query.Device
	_, err := d.Where(d.Key.Eq(key)).
		Update(d.IsOnline, 0)
	return err
}
