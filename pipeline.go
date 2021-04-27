package frunner

import "path/filepath"

type Pipeline struct {
	Steps []*Step `yaml:"steps"`
}

func (p *Pipeline) FillDir(dir string) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		absDir = dir
	}
	for i := range p.Steps {
		step := p.Steps[i]
		step.Dir = absDir
	}
}
