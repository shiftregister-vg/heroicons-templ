// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PlayPause(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M12.75 4C12.3358 4 12 4.33579 12 4.75V15.25C12 15.6642 12.3358 16 12.75 16H13.25C13.6642 16 14 15.6642 14 15.25V4.75C14 4.33579 13.6642 4 13.25 4H12.75Z\" fill=\"#0F172A\"></path> <path d=\"M17.75 4C17.3358 4 17 4.33579 17 4.75V15.25C17 15.6642 17.3358 16 17.75 16H18.25C18.6642 16 19 15.6642 19 15.25V4.75C19 4.33579 18.6642 4 18.25 4H17.75Z\" fill=\"#0F172A\"></path> <path d=\"M3.28824 4.81859C2.28891 4.20135 1 4.9202 1 6.09478V13.9053C1 15.0798 2.2889 15.7987 3.28824 15.1815L9.61101 11.2762C10.56 10.6901 10.56 9.30998 9.61101 8.72382L3.28824 4.81859Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
