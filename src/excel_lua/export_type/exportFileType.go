package export_type

type FileType int

const (
	// 导出Lua
	Lua FileType = iota
	// 导出json
	Json
	CSharp
)
