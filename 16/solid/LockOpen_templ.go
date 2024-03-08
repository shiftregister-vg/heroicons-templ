// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func LockOpen() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M11.5 1C9.567 1 8 2.567 8 4.5V7H2.5C1.67157 7 1 7.67157 1 8.5V13.5C1 14.3284 1.67157 15 2.5 15H9.5C10.3284 15 11 14.3284 11 13.5V8.5C11 7.67157 10.3284 7 9.5 7V4.5C9.5 3.39543 10.3954 2.5 11.5 2.5C12.6046 2.5 13.5 3.39543 13.5 4.5V6.25C13.5 6.66421 13.8358 7 14.25 7C14.6642 7 15 6.66421 15 6.25V4.5C15 2.567 13.433 1 11.5 1Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}