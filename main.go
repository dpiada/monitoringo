package main

import ( 
	"log"
    "time"
	"syscall"
	"github.com/gin-gonic/gin"
) 

func monitoring_perfomance() {
	var sysinfo syscall.Sysinfo_t

	
	for {
		time.Sleep(10 * time.Second)

		err := syscall.Sysinfo(&sysinfo)
		if err != nil {
			log.Printf("Error retrieving system memory info: %v\n", err)
			continue
		}

		totalRAM := sysinfo.Totalram * uint64(sysinfo.Unit)
		freeRAM := sysinfo.Freeram * uint64(sysinfo.Unit)
		usedRAM := totalRAM - freeRAM

		log.Printf("System memory: Total: %d bytes | Used: %d bytes | Free: %d bytes\n", totalRAM, usedRAM, freeRAM)
	}
}


func main() {
	go monitoring_perfomance()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}