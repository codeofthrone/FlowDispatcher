package device

import "sync"

// DeviceStatus 表示設備的狀態
type DeviceStatus struct {
	Status string
}

// DeviceStatusMap 存儲所有設備的狀態
var DeviceStatusMap = make(map[string]*DeviceStatus)
var mapMutex sync.RWMutex

// SetDeviceStatus 設定設備的狀態
func SetDeviceStatus(deviceID, status string) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	if deviceStatus, exists := DeviceStatusMap[deviceID]; exists {
		deviceStatus.Status = status
	} else {
		DeviceStatusMap[deviceID] = &DeviceStatus{
			Status: status,
		}
	}
}

func GetDeviceStatus(deviceID string) (*DeviceStatus, bool) {
	mapMutex.RLock()
	defer mapMutex.RUnlock()

	deviceStatus, exists := DeviceStatusMap[deviceID]
	return deviceStatus, exists
}
