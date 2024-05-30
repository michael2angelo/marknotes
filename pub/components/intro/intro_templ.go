// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package pub_intro

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Intro() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col prose-sm prose-a:font-extrabold prose-a:underline\"><p class=\"text-center mt-16 text-lg mx-4\" hx-boost=\"true\"><span class=\"text-4xl\">Hello! 🖖 My name is <strong>Wyndham </strong></span><br><br><span class=\"text-xl\">And I'm a Software Engineer that really into Pragmatic, Practical, and Beautiful Code.</span><br><br>My stack mostly consist of <strong>Android Development </strong> (Kotlin, Java), <strong>Backend Development</strong> (Golang, Rails), and <strong>Frontend Development</strong> (Plain HTML, CSS, JS).<br><br>More on my resume <a href=\"/resume\">here </a>.<br><br>Hit me up via  <a href=\"https://x.com/muhwyndham\">Twitter</a>, <a href=\"https://www.linkedin.com/in/m-wyndham-haryata-permana-b43ab2134/\">LinkedIn</a> or  <a href=\"mailto:business@mwyndham.dev\">Email</a> whenever you need help or just want to have some chit-chat!</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
