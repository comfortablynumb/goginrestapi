package module

// Structs

type ModuleManager struct {
	Modules []Module
}

func (m *ModuleManager) GetModules() []Module {
	return m.Modules
}

func (m *ModuleManager) AddModule(module Module) *ModuleManager {
	m.Modules = append(m.Modules, module)

	return m
}

// Static Functions

func NewModuleManager() *ModuleManager {
	return &ModuleManager{
		Modules: make([]Module, 0),
	}
}
