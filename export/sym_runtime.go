package export

// Generated by 'go run gen.go runtime'. Do not edit!

import "runtime"

var sym_runtime = &map[string]interface{}{
	"BlockProfile":            runtime.BlockProfile,
	"BlockProfileRecord":      new(runtime.BlockProfileRecord),
	"Breakpoint":              runtime.Breakpoint,
	"CPUProfile":              runtime.CPUProfile,
	"Caller":                  runtime.Caller,
	"Callers":                 runtime.Callers,
	"CallersFrames":           runtime.CallersFrames,
	"Compiler":                runtime.Compiler,
	"Error":                   new(runtime.Error),
	"Frame":                   new(runtime.Frame),
	"Frames":                  new(runtime.Frames),
	"Func":                    new(runtime.Func),
	"FuncForPC":               runtime.FuncForPC,
	"GC":                      runtime.GC,
	"GOARCH":                  runtime.GOARCH,
	"GOMAXPROCS":              runtime.GOMAXPROCS,
	"GOOS":                    runtime.GOOS,
	"GOROOT":                  runtime.GOROOT,
	"Goexit":                  runtime.Goexit,
	"GoroutineProfile":        runtime.GoroutineProfile,
	"Gosched":                 runtime.Gosched,
	"KeepAlive":               runtime.KeepAlive,
	"LockOSThread":            runtime.LockOSThread,
	"MemProfile":              runtime.MemProfile,
	"MemProfileRate":          runtime.MemProfileRate,
	"MemProfileRecord":        new(runtime.MemProfileRecord),
	"MemStats":                new(runtime.MemStats),
	"MutexProfile":            runtime.MutexProfile,
	"NumCPU":                  runtime.NumCPU,
	"NumCgoCall":              runtime.NumCgoCall,
	"NumGoroutine":            runtime.NumGoroutine,
	"ReadMemStats":            runtime.ReadMemStats,
	"ReadTrace":               runtime.ReadTrace,
	"SetBlockProfileRate":     runtime.SetBlockProfileRate,
	"SetCPUProfileRate":       runtime.SetCPUProfileRate,
	"SetCgoTraceback":         runtime.SetCgoTraceback,
	"SetFinalizer":            runtime.SetFinalizer,
	"SetMutexProfileFraction": runtime.SetMutexProfileFraction,
	"Stack":                   runtime.Stack,
	"StackRecord":             new(runtime.StackRecord),
	"StartTrace":              runtime.StartTrace,
	"StopTrace":               runtime.StopTrace,
	"ThreadCreateProfile":     runtime.ThreadCreateProfile,
	"TypeAssertionError":      new(runtime.TypeAssertionError),
	"UnlockOSThread":          runtime.UnlockOSThread,
	"Version":                 runtime.Version,
}
