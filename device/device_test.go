package device

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentSetAndGetDeviceStatus(t *testing.T) {
	var wg sync.WaitGroup
	deviceCount := 10    // 設置10個設備
	concurrentUsers := 5 // 每個設備有5個並發用戶

	for i := 0; i < deviceCount; i++ {
		deviceID := fmt.Sprintf("device%d", i)

		for j := 0; j < concurrentUsers; j++ {
			wg.Add(1)
			go func(userID int) {
				defer wg.Done()
				status := fmt.Sprintf("user%d_status", userID)

				// 設定設備狀態
				SetDeviceStatus(deviceID, status)

				// 獲取並驗證設備狀態
				got, exists := GetDeviceStatus(deviceID)
				if !exists || got.Status != status {
					t.Errorf("用戶 %d 存取設備 %s 獲取狀態錯誤, 獲取到: %v, 預期: %v", userID, deviceID, got, status)
				}
			}(j)
		}
	}

	wg.Wait()
}
