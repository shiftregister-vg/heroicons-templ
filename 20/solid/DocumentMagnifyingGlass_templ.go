// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func DocumentMagnifyingGlass(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M8 10C8 9.17157 8.67157 8.5 9.5 8.5C10.3284 8.5 11 9.17157 11 10C11 10.4144 10.8329 10.7884 10.5607 11.0607C10.2884 11.3329 9.91442 11.5 9.5 11.5C8.67157 11.5 8 10.8284 8 10Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4.5 2C3.67157 2 3 2.67157 3 3.5V16.5C3 17.3284 3.67157 18 4.5 18H15.5C16.3284 18 17 17.3284 17 16.5V7.62132C17 7.2235 16.842 6.84197 16.5607 6.56066L12.4393 2.43934C12.158 2.15804 11.7765 2 11.3787 2H4.5ZM9.5 7C7.84315 7 6.5 8.34315 6.5 10C6.5 11.6569 7.84315 13 9.5 13C10.056 13 10.5773 12.8483 11.0239 12.5845L12.2197 13.7803C12.5126 14.0732 12.9874 14.0732 13.2803 13.7803C13.5732 13.4874 13.5732 13.0126 13.2803 12.7197L12.0845 11.5239C12.3483 11.0773 12.5 10.556 12.5 10C12.5 8.34315 11.1569 7 9.5 7Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
