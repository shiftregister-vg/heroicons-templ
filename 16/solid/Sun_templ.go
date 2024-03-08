// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Sun(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M8 1C8.41421 1 8.75 1.33579 8.75 1.75V3.25C8.75 3.66421 8.41421 4 8 4C7.58579 4 7.25 3.66421 7.25 3.25V1.75C7.25 1.33579 7.58579 1 8 1Z\" fill=\"black\"></path> <path d=\"M10.5 8C10.5 9.38071 9.38071 10.5 8 10.5C6.61929 10.5 5.5 9.38071 5.5 8C5.5 6.61929 6.61929 5.5 8 5.5C9.38071 5.5 10.5 6.61929 10.5 8Z\" fill=\"black\"></path> <path d=\"M12.9498 4.11089C13.2427 3.81799 13.2427 3.34312 12.9498 3.05023C12.6569 2.75733 12.182 2.75733 11.8891 3.05023L10.8284 4.11089C10.5356 4.40378 10.5356 4.87865 10.8284 5.17155C11.1213 5.46444 11.5962 5.46444 11.8891 5.17155L12.9498 4.11089Z\" fill=\"black\"></path> <path d=\"M15 8C15 8.41421 14.6642 8.75 14.25 8.75H12.75C12.3358 8.75 12 8.41421 12 8C12 7.58579 12.3358 7.25 12.75 7.25H14.25C14.6642 7.25 15 7.58579 15 8Z\" fill=\"black\"></path> <path d=\"M11.8891 12.9498C12.182 13.2427 12.6569 13.2427 12.9498 12.9498C13.2427 12.6569 13.2427 12.182 12.9498 11.8891L11.8891 10.8284C11.5962 10.5356 11.1213 10.5356 10.8285 10.8284C10.5356 11.1213 10.5356 11.5962 10.8285 11.8891L11.8891 12.9498Z\" fill=\"black\"></path> <path d=\"M8 12C8.41421 12 8.75 12.3358 8.75 12.75V14.25C8.75 14.6642 8.41421 15 8 15C7.58579 15 7.25 14.6642 7.25 14.25V12.75C7.25 12.3358 7.58579 12 8 12Z\" fill=\"black\"></path> <path d=\"M5.17157 11.8891C5.46446 11.5962 5.46446 11.1213 5.17157 10.8284C4.87867 10.5355 4.4038 10.5355 4.11091 10.8284L3.05025 11.8891C2.75735 12.182 2.75735 12.6569 3.05025 12.9497C3.34314 13.2426 3.81801 13.2426 4.11091 12.9497L5.17157 11.8891Z\" fill=\"black\"></path> <path d=\"M4 8C4 8.41421 3.66421 8.75 3.25 8.75H1.75C1.33579 8.75 1 8.41421 1 8C1 7.58579 1.33579 7.25 1.75 7.25H3.25C3.66421 7.25 4 7.58579 4 8Z\" fill=\"black\"></path> <path d=\"M4.11091 5.17157C4.40381 5.46446 4.87868 5.46446 5.17157 5.17157C5.46447 4.87867 5.46447 4.4038 5.17157 4.11091L4.11091 3.05025C3.81802 2.75735 3.34315 2.75735 3.05025 3.05024C2.75736 3.34314 2.75736 3.81801 3.05025 4.11091L4.11091 5.17157Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
