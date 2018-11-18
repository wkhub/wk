package shell

import (
	"io"
	"os"
	"text/template"
)

// Session stores every thing necessary to launch a shell session
type Session struct {
	Cwd  string
	Env  Env
	Init []string          // A list of command to launch on start
	Dirs map[string]string // A list directory shortcuts
}

func NewSession(fresh bool) Session {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var env Env
	if fresh {
		env = Env{}
	} else {
		env = GetEnv()
	}
	return Session{
		Cwd:  cwd,
		Env:  env,
		Init: []string{},
		Dirs: map[string]string{},
	}
}

func (s *Session) Update(other Session) *Session {
	s.Env.Update(other.Env)
	s.Init = append(s.Init, other.Init...)
	return s
}

func (s Session) Render(tpl string, w io.Writer) {
	tmpl, err := template.New("session").Parse(tpl)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, s)
	if err != nil {
		panic(err)
	}
}
