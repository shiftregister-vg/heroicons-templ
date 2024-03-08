// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreditCard(classNames string) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{classNames}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M2.5 3C1.67157 3 1 3.67157 1 4.5V5H15V4.5C15 3.67157 14.3284 3 13.5 3H2.5Z\" fill=\"#333333\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15 7H1V11.5C1 12.3284 1.67157 13 2.5 13H13.5C14.3284 13 15 12.3284 15 11.5V7ZM3 10.25C3 9.83579 3.33579 9.5 3.75 9.5H4.25C4.66421 9.5 5 9.83579 5 10.25C5 10.6642 4.66421 11 4.25 11H3.75C3.33579 11 3 10.6642 3 10.25ZM6.75 9.5C6.33579 9.5 6 9.83579 6 10.25C6 10.6642 6.33579 11 6.75 11H9.25C9.66421 11 10 10.6642 10 10.25C10 9.83579 9.66421 9.5 9.25 9.5H6.75Z\" fill=\"#333333\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
