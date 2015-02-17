package main

type Processor struct {
	processor       int      `json:"processor"`
	vendorId        string   `json:"vendor_id, omitempty"`
	cpuFamily       int      `json:"cpu family, omitempty"`
	model           int      `json:"model, omitempty"`
	modelName       string   `json:"model name, omitempty"`
	stepping        int      `json:"stepping, omitempty"`
	microcode       string   `json:"microcode, omitempty"`
	cpuMhz          float32  `json:"cpu MHz, omitempty"`
	cacheSize       string   `json:"cache size, omitempty"`
	physicalId      int      `json:"physical id, omitempty"`
	siblings        int      `json:"siblings, omitempty"`
	coreId          int      `json:"core id, omitempty"`
	cpuCores        int      `json:"cpu cores, omitempty"`
	apicId          int      `json:"apicid, omitempty"`
	fpu             bool     `json:"fpu, omitempty"`
	fpuException    bool     `json:"fpu_exception, omitempty"`
	cpuidLevel      int      `json:"cpuid level, omitempty"`
	wp              bool     `json:"wp, omitempty"`
	flags           []string `json:"flags, omitempty"`
	bogomips        float32  `json:"bogomips, omitempty"`
	clflushSize     int      `json:"clflush size, omitempty"`
	cacheAlignment  int      `json:"cache_alignment, omitempty"`
	addressSizes    string   `json:"address sizes, omitempty"`
	powerManagement string   `json:"power management, omitempty"`
}
