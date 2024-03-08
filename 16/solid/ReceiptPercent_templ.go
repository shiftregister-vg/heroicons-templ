// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ReceiptPercent(classNames string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5.25 2C4.00736 2 3 3.00736 3 4.25V13.25C3 13.5302 3.15618 13.787 3.40495 13.9159C3.65373 14.0448 3.95361 14.0243 4.18251 13.8627L5.875 12.668L7.56749 13.8627C7.82678 14.0458 8.17322 14.0458 8.43251 13.8627L10.125 12.668L11.8175 13.8627C12.0464 14.0243 12.3463 14.0448 12.595 13.9159C12.8438 13.787 13 13.5302 13 13.25V4.25C13 3.00736 11.9926 2 10.75 2H5.25ZM10.7803 6.28033C11.0732 5.98744 11.0732 5.51256 10.7803 5.21967C10.4874 4.92678 10.0126 4.92678 9.71967 5.21967L5.21967 9.71967C4.92678 10.0126 4.92678 10.4874 5.21967 10.7803C5.51256 11.0732 5.98744 11.0732 6.28033 10.7803L10.7803 6.28033ZM7 6.25C7 6.66421 6.66421 7 6.25 7C5.83579 7 5.5 6.66421 5.5 6.25C5.5 5.83579 5.83579 5.5 6.25 5.5C6.66421 5.5 7 5.83579 7 6.25ZM9.75 10.5C10.1642 10.5 10.5 10.1642 10.5 9.75C10.5 9.33579 10.1642 9 9.75 9C9.33579 9 9 9.33579 9 9.75C9 10.1642 9.33579 10.5 9.75 10.5Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
