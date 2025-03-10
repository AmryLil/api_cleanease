package helpers

import (
	"fmt"
	"sync"

	"github.com/sony/sonyflake"
)

var (
	sf   *sonyflake.Sonyflake
	once sync.Once
)

func initSonyflake() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		fmt.Println("Sonyflake not created")
	}
}

// GenerateID menggunakan Sonyflake yang sudah diinisialisasi sebelumnya
func GenerateID() uint {
	once.Do(initSonyflake) // Pastikan Sonyflake hanya diinisialisasi sekali

	if sf == nil {
		fmt.Println("Sonyflake not created")
		return 0
	}

	id, err := sf.NextID()
	if err != nil {
		fmt.Println("Failed to generate ID:", err)
		return 0
	}
	return uint(id)
}
