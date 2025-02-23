package pluginstorage

import "plugin"

type PluginStorage struct {
	plugins map[string]*plugin.Plugin
}

func InitPluginStorage() *PluginStorage {
	ps := new(PluginStorage)
	ps.plugins = make(map[string]*plugin.Plugin)
	return ps
}

func (ps *PluginStorage) LoadPlugin(providerName, pluginPath string) error {
	plugin, err := plugin.Open(pluginPath)
	if err != nil {
		return nil
	}
	ps.plugins[providerName] = plugin

	return nil
}

func (ps *PluginStorage) GetPlugin(providerName string) *plugin.Plugin {
	return ps.plugins[providerName]
}
