package module

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	strings2 "github.com/auho/go-handknife/emergencybox/toolkit/strings"
)

//go:embed tmpl/command.tmpl
var commandTmpl string

//go:embed tmpl/module.tmpl
var moduleTmpl string

type Module struct {
	appName string
}

func NewSub(appName string) *Module {
	return &Module{appName: appName}
}

func (m *Module) Build(module, cmd, subCmd string) error {
	if module == "" {
		return errors.New("module is empty")
	}

	if cmd == "" {
		return errors.New("cmd is empty")
	}

	if module == cmd {
		return errors.New("module is the same as cmd")
	}

	dir, err := m.buildModule(module)
	if err != nil {
		return fmt.Errorf("build module error: %w", err)
	}

	err = m.buildCmd(module, cmd, dir)
	if err != nil {
		return fmt.Errorf("build cmd error: %w", err)
	}

	if subCmd != "" {
		err = m.buildSubCmd(cmd, subCmd, dir)
		if err != nil {
			return fmt.Errorf("build subcmd error: %w", err)
		}
	}

	return nil
}

func (m *Module) buildModule(module string) (string, error) {
	dir, err := m.makeModuleDir(module)
	if err != nil {
		return "", fmt.Errorf("module dir error: %w", err)
	}

	path := m.namingFilePath(dir, module)
	err = m.moduleContent(path, module)
	if err != nil {
		return "", fmt.Errorf("write file error: %w", err)
	}

	return dir, nil
}

func (m *Module) buildCmd(module, name, dir string) error {
	path := m.namingFilePath(dir, name)
	return m.cmdContent(path, module, name)
}

func (m *Module) buildSubCmd(cmd, subCmd, dir string) error {
	dir = m.namingDirPath(dir, cmd+"/cmd")

	err := m.mkdir(dir)
	if err != nil {
		return err
	}

	path := m.namingFilePath(dir, subCmd)

	return m.cmdContent(path, cmd, subCmd)
}

func (m *Module) makeModuleDir(sub string) (string, error) {
	_path := "modules"
	if sub != "" {
		_path = _path + "/" + sub
	}

	err := m.mkdir(_path)
	return _path, err
}

func (m *Module) mkdir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0744)
			if err != nil {
				return fmt.Errorf("make dir[%s] error: %w", dir, err)
			}
		} else {
			return fmt.Errorf("make dir[%s] error: %w", dir, err)
		}
	}

	return nil
}

func (m *Module) namingDirPath(dir, name string) string {
	return dir + "/" + strings2.ToUnderlineNaming(name)
}

func (m *Module) namingFilePath(dir, name string) string {
	return dir + "/" + strings2.ToUnderlineNaming(name) + ".go"
}

func (m *Module) moduleContent(path, module string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open file error: %w", err)
	}

	defer func() { _ = f.Close() }()

	_t, err := template.New("module").Parse(moduleTmpl)
	if err != nil {
		return fmt.Errorf("parse module template error: %w", err)
	}

	_module := strings2.ToHumpNaming(module)
	_moduleName := strings.ToUpper(_module[0:1]) + _module[1:]

	err = _t.Execute(f, map[string]any{
		"package":    strings2.ToUnderlineNaming(module),
		"module":     _module,
		"moduleName": _moduleName,
	})
	if err != nil {
		return fmt.Errorf("execute module template error: %w", err)
	}

	return nil
}

func (m *Module) cmdContent(path, module, cmd string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open file[%s] error: %w", path, err)
	}

	defer func() { _ = f.Close() }()

	_t, err := template.New("cmd").Parse(commandTmpl)
	if err != nil {
		return fmt.Errorf("parse cmd template error: %w", err)
	}

	_cmd := strings2.ToHumpNaming(cmd)
	_cmdName := strings.ToUpper(_cmd[0:1]) + _cmd[1:]

	err = _t.Execute(f, map[string]any{
		"app":     m.appName,
		"package": strings2.ToUnderlineNaming(module),
		"cmd":     _cmd,
		"cmdName": _cmdName,
	})
	if err != nil {
		return fmt.Errorf("execute cmd template error: %w", err)
	}

	return nil
}
