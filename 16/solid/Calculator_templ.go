// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Calculator(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5 1C3.89543 1 3 1.89543 3 3V13C3 14.1046 3.89543 15 5 15H11C12.1046 15 13 14.1046 13 13V3C13 1.89543 12.1046 1 11 1H5ZM5.75 7C5.33579 7 5 7.33579 5 7.75C5 8.16421 5.33579 8.5 5.75 8.5C6.16421 8.5 6.5 8.16421 6.5 7.75C6.5 7.33579 6.16421 7 5.75 7ZM5 3.75C5 3.33579 5.33579 3 5.75 3H10.25C10.6642 3 11 3.33579 11 3.75C11 4.16421 10.6642 4.5 10.25 4.5H5.75C5.33579 4.5 5 4.16421 5 3.75ZM5.75 11.5C5.33579 11.5 5 11.8358 5 12.25C5 12.6642 5.33579 13 5.75 13C6.16421 13 6.5 12.6642 6.5 12.25C6.5 11.8358 6.16421 11.5 5.75 11.5ZM5 10C5 9.58579 5.33579 9.25 5.75 9.25C6.16421 9.25 6.5 9.58579 6.5 10C6.5 10.4142 6.16421 10.75 5.75 10.75C5.33579 10.75 5 10.4142 5 10ZM10.25 7C9.83579 7 9.5 7.33579 9.5 7.75C9.5 8.16421 9.83579 8.5 10.25 8.5C10.6642 8.5 11 8.16421 11 7.75C11 7.33579 10.6642 7 10.25 7ZM9.5 10C9.5 9.58579 9.83579 9.25 10.25 9.25C10.6642 9.25 11 9.58579 11 10V12.25C11 12.6642 10.6642 13 10.25 13C9.83579 13 9.5 12.6642 9.5 12.25V10ZM8 7C7.58579 7 7.25 7.33579 7.25 7.75C7.25 8.16421 7.58579 8.5 8 8.5C8.41421 8.5 8.75 8.16421 8.75 7.75C8.75 7.33579 8.41421 7 8 7ZM7.25 12.25C7.25 11.8358 7.58579 11.5 8 11.5C8.41421 11.5 8.75 11.8358 8.75 12.25C8.75 12.6642 8.41421 13 8 13C7.58579 13 7.25 12.6642 7.25 12.25ZM8 9.25C7.58579 9.25 7.25 9.58579 7.25 10C7.25 10.4142 7.58579 10.75 8 10.75C8.41421 10.75 8.75 10.4142 8.75 10C8.75 9.58579 8.41421 9.25 8 9.25Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
