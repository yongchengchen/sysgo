package sysstat

import (
	"fmt"
	"strconv"
	"strings"

	"os/exec"
	"time"

	"github.com/yongchengchen/sysgo/contract"
	"github.com/yongchengchen/sysgo/services/container"
	"github.com/mitchellh/mapstructure"

	"github.com/yongchengchen/sysgo/services/dynamosvc"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

type MEMINFO struct {
	Total  uint64 `json:"total"`
	Used   uint64 `json:"used"`
	Cached uint64 `json:"cached"`
	Free   uint64 `json:"free"`
}
type CPUINFO struct {
	User   uint64 `json:"user"`
	System uint64 `json:"system"`
	Idle   uint64 `json:"idle"`
}

type SYSTEMINFO struct {
	Cpu  CPUINFO  `json:"cpu"`
	Mem  MEMINFO  `json:"mem"`
	Disk DISKINFO `json:"disk"`
	Time int64    `json:"time"`
}

type DISKINFO struct {
	Used  uint64 `json:"used"`
	Total uint64 `json:"total"`
}

type TABEL_RECORD struct {
	Name    string     `json:"name"`
	Sysinfo SYSTEMINFO `json:"sysinfo"`
}

type SYSSTAT_NODE_CONFIG struct {
	Name      string `json:"name"`
	Diskpath  string `json:"diskpath"`
	Dynotable string `json:"dynotable"`
}

var nodeConfig SYSSTAT_NODE_CONFIG

func loadConfig() {
	if conf, ok := container.Get("config").(contract.IConfig); ok {
		configs := conf.Get("app.sysstat")

		if err := mapstructure.Decode(configs, &nodeConfig); err != nil {
			fmt.Printf("app.sysstat config is not correct %#v\n", configs)
		}
	}
}

func DiskUsage(path string) (uint64, uint64, error) {
	out, _ := exec.Command("df", "-P").Output()
	outlines := strings.Split(string(out), "\n")
	l := len(outlines)
	var total, used uint64 = 0, 0
	for _, line := range outlines[1 : l-1] {
		parsedLine := strings.Fields(line)
		if path == parsedLine[0] {
			t, err := strconv.ParseUint(parsedLine[1], 0, 64)
			if err != nil {
				return 0, 0, err
			}
			u, err := strconv.ParseUint(parsedLine[2], 0, 64)
			if err != nil {
				return 0, 0, err
			}
			total += t
			used += u
		}
	}
	return used, total, nil
}

func Post() {
	loadConfig()
	diskUsed, diskTotal, err := DiskUsage(nodeConfig.Diskpath)
	if err != nil {
		return
	}

	memory, err := memory.Get()
	if err != nil {
		return
	}

	cpu, err := cpu.Get()
	if err != nil {
		return
	}

	record := TABEL_RECORD{
		Name: nodeConfig.Name,
		Sysinfo: SYSTEMINFO{
			Cpu: CPUINFO{
				User:   cpu.User * 100 / cpu.Total,
				System: cpu.System * 100 / cpu.Total,
				Idle:   cpu.Idle * 100 / cpu.Total,
			},
			Mem: MEMINFO{
				Total:  memory.Total / 1024 / 1024,
				Used:   memory.Used / 1024 / 1024,
				Cached: memory.Cached / 1024 / 1024,
				Free:   memory.Free / 1024 / 1024,
			},
			Disk: DISKINFO{
				Used:  diskUsed,
				Total: diskTotal,
			},

			Time: time.Now().Unix(),
		},
	}
	dynamosvc.PutRecord(nodeConfig.Dynotable, record)
}
