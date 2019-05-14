package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
)

// Modify nvml processInfo sample to include gpu utilization and uid/username 
const (
	PINFOHEADER = `# gpu   pid   type  mem  pwr  temp    sm  Command
# Idx     #   C/G   MiB  W    C       %   name`
)

func main() {
	nvml.Init()
	defer nvml.Shutdown()

	count, err := nvml.GetDeviceCount()
	if err != nil {
		log.Panicln("Error getting device count:", err)
	}

	var devices []*nvml.Device
	for i := uint(0); i < count; i++ {
		device, err := nvml.NewDevice(i)
		if err != nil {
			log.Panicf("Error getting device %d: %v\n", i, err)
		}
		devices = append(devices, device)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

/* 
Ticker set to 2 seconds for testing. Will display output every two seconds.
Adjust ticket time to every 30 seconds (production).
Timer set to 60 seconds for testing. 
Adjust timer to 43200 (12 hours) for production
*/
        ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()
        timer := time.NewTimer(60 * time.Second)
        defer timer.Stop()

	fmt.Println(PINFOHEADER)
	for {
		select {
		case <-ticker.C:
			for i, device := range devices {
				pInfo, err := device.GetAllRunningProcesses()
                                st, err := device.Status()
				if err != nil {
					log.Panicf("Error getting device %d processes: %v\n", i, err)
				}
				if len(pInfo) == 0 {
//					fmt.Printf("%5v %5s %5s %5s %-5s\n", i, "-", "-", "-", "-")
                                        continue
				}
				for j := range pInfo {
					fmt.Printf("%5v %5v %5v %5v %5d %5d %5d %-5v\n",
						i, pInfo[j].PID, pInfo[j].Type, pInfo[j].MemoryUsed,*st.Power, *st.Temperature, *st.Utilization.GPU, pInfo[j].Name)
				}
			}
		case <-sigs:
			return
                case <-timer.C:
//                      fmt.Println("timeout 1")   //Uncomment for testing purposes.
                        return
		}
	}
}
