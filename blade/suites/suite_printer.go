package suites

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jedib0t/go-pretty/v6/text"
)

func (se *Suite) printSliceMapStringAny(title string, fn func() ([]map[string]any, *Post, error)) {
	se.printFunc(title, func() (*Post, string, error) {
		var body string
		sm, p, err := fn()
		if err == nil {
			body, err = se.renderSliceMapStringAny(sm, p)
		}

		return p, body, err
	})
}

func (se *Suite) printFunc(title string, fn func() (*Post, string, error)) {
	se.pfIndexIn()

	s := &strings.Builder{}
	s.WriteString(se.stringTitle(title) + "\n")

	p, body, err := fn()

	se.pfIndexExec(func(_f func() string) {
		s.WriteString(_f())
		s.WriteString("\n")
	})

	if err != nil {
		se.PrintlnErr("[Error] ", err)
	} else {
		s.WriteString(body)
	}

	se.Cmd.Print(s.String())
	se.Cmd.Println("")

	if p == nil {
		p = NewPost()
	}

	ok, err := p.ExtraExec()
	if ok {
		if err != nil {
			var nErr validator.ValidationErrors
			switch {
			case errors.As(err, &nErr):
				for k, fieldError := range nErr {
					se.PrintlnErr("extra:", k, fieldError)
				}
			default:
				se.PrintlnErr("extra:", err)
			}
		}

		se.Cmd.Println("")
	}
}

func (se *Suite) PfBody(body string) {
	se.pfFunc[se.pfIndex] = append(se.pfFunc[se.pfIndex], func() string {
		return se.stringBody(body)
	})
}

func (se *Suite) PfLn(a ...any) {
	se.pfFunc[se.pfIndex] = append(se.pfFunc[se.pfIndex], func() string {
		return fmt.Sprint(a...)
	})
}

func (se *Suite) PfErr(a ...any) {
	se.pfFunc[se.pfIndex] = append(se.pfFunc[se.pfIndex], func() string {
		return se.stringErr(a...)
	})
}

func (se *Suite) Println(s string) {
	se.Cmd.Println(s)
}

func (se *Suite) PrintlnAny(i ...any) {
	se.Cmd.Println(i...)
}

func (se *Suite) PrintlnTitle(title string) {
	se.Cmd.Println(se.stringTitle(title))
}

func (se *Suite) PrintlnBody(body string) {
	se.Cmd.Println(se.stringBody(body))
}

func (se *Suite) PrintlnErr(a ...any) {
	se.Cmd.PrintErrln(se.stringErr(a...))
}

func (se *Suite) stringTitle(title string) string {
	return text.FgBlue.Sprint(title)
}

func (se *Suite) stringBody(body string) string {
	return text.FgYellow.Sprint(body)
}

func (se *Suite) stringErr(a ...any) string {
	return text.FgHiRed.Sprint(a...)
}

func (se *Suite) renderSliceMapStringAny(sm []map[string]any, p *Post) (string, error) {
	if p == nil {
		p = NewPost()
	}

	return p.render(sm)
}
