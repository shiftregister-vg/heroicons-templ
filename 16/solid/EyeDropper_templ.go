// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func EyeDropper(classNames string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15 4C15 5.39788 14.0439 6.57245 12.75 6.90549V8.5C12.75 8.69891 12.671 8.88968 12.5303 9.03033L12.0303 9.53033C11.7374 9.82322 11.2626 9.82322 10.9697 9.53033L10.25 8.81069L5.57322 13.4875C5.24503 13.8157 4.79992 14.0001 4.33579 14.0001H3.66421C3.59791 14.0001 3.53432 14.0264 3.48744 14.0733L2.78033 14.7804C2.63968 14.921 2.44891 15.0001 2.25 15.0001C2.05109 15.0001 1.86032 14.921 1.71967 14.7804L1.21967 14.2804C0.926777 13.9875 0.926777 13.5126 1.21967 13.2197L1.92678 12.5126C1.97366 12.4657 2 12.4021 2 12.3358V11.6643C2 11.2001 2.18437 10.755 2.51256 10.4268L7.18937 5.75003L6.46967 5.03033C6.17678 4.73744 6.17678 4.26256 6.46967 3.96967L6.96967 3.46967C7.11032 3.32902 7.30109 3.25 7.5 3.25H9.09451C9.42755 1.95608 10.6021 1 12 1C13.6569 1 15 2.34315 15 4ZM9.18937 7.75003L8.25003 6.81069L3.57322 11.4875C3.52634 11.5344 3.5 11.598 3.5 11.6643V12.3358C3.5 12.3938 3.49713 12.4514 3.49146 12.5086C3.54862 12.5029 3.60627 12.5001 3.66421 12.5001H4.33579C4.40209 12.5001 4.46568 12.4737 4.51256 12.4268L9.18937 7.75003Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
