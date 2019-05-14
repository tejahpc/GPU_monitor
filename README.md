# GPU_monitor
Monitors the usage of GPU for an HPC cluster.

## Getting Started 
The gpu_mon binary monitors the GPU and memory utilization of a node. In production setting, the stats are logged in every 30 seconds for 12 hours. This monitors whether users are using the GPUs efficiently and how the HPC staff may help improve efficiency. 

### Installation

The app uses go. For installation in Biowulf, first download the Nvidia NVML Go bindings
```
module load golang
export GOPATH=~/go
cd ~/go
go get github.com/NVIDIA/gpu-monitoring-tools
```

Once NVML is installed, this repo can be cloned and built 
```
go build . 
```


