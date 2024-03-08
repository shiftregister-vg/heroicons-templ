// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ClipboardDocumentList(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.9862 3H12C13.1045 3 14 3.89543 14 5V11C14 11.9319 13.3626 12.715 12.5 12.937V7C12.5 5.61929 11.3807 4.5 9.99998 4.5H4.06299C4.28501 3.63739 5.06806 3 5.99998 3H6.01371C6.13807 1.87501 7.09184 1 8.24998 1H9.74998C10.9081 1 11.8619 1.87501 11.9862 3ZM10.5 4V3.25C10.5 2.83579 10.1642 2.5 9.74998 2.5H8.24998C7.83577 2.5 7.49998 2.83579 7.49998 3.25V4H10.5Z\" fill=\"black\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3 6C2.44772 6 2 6.44772 2 7V14C2 14.5523 2.44772 15 3 15H10C10.5523 15 11 14.5523 11 14V7C11 6.44772 10.5523 6 10 6H3ZM4.75 8.5C4.33579 8.5 4 8.83579 4 9.25C4 9.66421 4.33579 10 4.75 10H8.25C8.66421 10 9 9.66421 9 9.25C9 8.83579 8.66421 8.5 8.25 8.5H4.75ZM4 11.75C4 11.3358 4.33579 11 4.75 11H8.25C8.66421 11 9 11.3358 9 11.75C9 12.1642 8.66421 12.5 8.25 12.5H4.75C4.33579 12.5 4 12.1642 4 11.75Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
