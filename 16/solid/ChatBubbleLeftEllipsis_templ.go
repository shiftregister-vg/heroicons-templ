// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChatBubbleLeftEllipsis() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1 8.74067C1 9.72338 1.71344 10.5655 2.68911 10.6828C3.45355 10.7748 4.22402 10.8472 5 10.8994V13.25C5 13.5533 5.18273 13.8268 5.46299 13.9429C5.74324 14.059 6.06583 13.9948 6.28033 13.7803L8.7905 11.2702C8.97217 11.0885 9.21682 10.9842 9.47361 10.9758C10.7676 10.9332 12.0475 10.8347 13.311 10.6826C14.2866 10.5653 15 9.72316 15 8.74048V4.25947C15 3.27678 14.2866 2.43469 13.3109 2.3173C11.5698 2.1078 9.79728 2 7.99962 2C6.20224 2 4.43002 2.10777 2.68909 2.31721C1.71343 2.43458 1 3.27668 1 4.25938V8.74067ZM5.5 6.5C5.5 7.05228 5.05228 7.5 4.5 7.5C3.94772 7.5 3.5 7.05228 3.5 6.5C3.5 5.94772 3.94772 5.5 4.5 5.5C5.05228 5.5 5.5 5.94772 5.5 6.5ZM8 7.5C8.55228 7.5 9 7.05228 9 6.5C9 5.94772 8.55228 5.5 8 5.5C7.44772 5.5 7 5.94772 7 6.5C7 7.05228 7.44772 7.5 8 7.5ZM11.5 7.5C12.0523 7.5 12.5 7.05228 12.5 6.5C12.5 5.94772 12.0523 5.5 11.5 5.5C10.9477 5.5 10.5 5.94772 10.5 6.5C10.5 7.05228 10.9477 7.5 11.5 7.5Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}