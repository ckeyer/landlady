package scan

type ScannerManger struct {
	modules map[string]interface{}
}

func New() *ScannerManger {
	return &ScannerManger{
		modules: nil,
	}
}

// Start
func (s *ScannerManger) Start() {
	for name, scanner := range s.modules {
		if pc, ok := scanner.(PageCounter); ok {
			pc.PageCount()
		}
	}
}
