// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Rss(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.75 4.5C3.75 4.08579 4.08579 3.75 4.5 3.75H5.25C13.5343 3.75 20.25 10.4657 20.25 18.75V19.5C20.25 19.9142 19.9142 20.25 19.5 20.25H18.75C18.3358 20.25 18 19.9142 18 19.5V18.75C18 11.7084 12.2916 6 5.25 6H4.5C4.08579 6 3.75 5.66421 3.75 5.25V4.5ZM3.75 11.25C3.75 10.8358 4.08579 10.5 4.5 10.5H5.25C9.80635 10.5 13.5 14.1937 13.5 18.75V19.5C13.5 19.9142 13.1642 20.25 12.75 20.25H12C11.5858 20.25 11.25 19.9142 11.25 19.5V18.75C11.25 15.4363 8.56371 12.75 5.25 12.75H4.5C4.08579 12.75 3.75 12.4142 3.75 12V11.25ZM3.75 18.75C3.75 17.9216 4.42157 17.25 5.25 17.25C6.07843 17.25 6.75 17.9216 6.75 18.75C6.75 19.5784 6.07843 20.25 5.25 20.25C4.42157 20.25 3.75 19.5784 3.75 18.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
