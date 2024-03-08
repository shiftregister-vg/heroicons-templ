// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SignalSlash(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M2.21967 2.21967C2.51256 1.92678 2.98744 1.92678 3.28033 2.21967L10.0626 9.00193C10.565 9.03295 10.967 9.43498 10.9981 9.93741L17.7803 16.7197C18.0732 17.0126 18.0732 17.4874 17.7803 17.7803C17.4874 18.0732 17.0126 18.0732 16.7197 17.7803L9.93741 10.9981C9.43498 10.967 9.03295 10.565 9.00193 10.0626L2.21967 3.28033C1.92678 2.98744 1.92678 2.51256 2.21967 2.21967Z\" fill=\"#0F172A\"></path> <path d=\"M3.63604 16.364C0.670686 13.3986 0.207189 8.87882 2.24555 5.42753L3.3485 6.53048C1.87267 9.35787 2.32207 12.9287 4.6967 15.3033C4.98959 15.5962 4.98959 16.0711 4.6967 16.364C4.40381 16.6569 3.92893 16.6569 3.63604 16.364Z\" fill=\"#0F172A\"></path> <path d=\"M6.46447 13.5356C5.08411 12.1552 4.67956 10.1685 5.25082 8.43281L6.51322 9.6952C6.4268 10.6896 6.76411 11.7139 7.52513 12.4749C7.81802 12.7678 7.81802 13.2427 7.52513 13.5356C7.23223 13.8284 6.75736 13.8284 6.46447 13.5356Z\" fill=\"#0F172A\"></path> <path d=\"M16.364 3.63605C19.3293 6.6014 19.7928 11.1212 17.7545 14.5725L16.6515 13.4695C18.1273 10.6421 17.6779 7.07134 15.3033 4.69671C15.0104 4.40382 15.0104 3.92895 15.3033 3.63605C15.5962 3.34316 16.0711 3.34316 16.364 3.63605Z\" fill=\"#0F172A\"></path> <path d=\"M13.5355 6.46449C14.9159 7.84484 15.3204 9.83147 14.7492 11.5672L13.4868 10.3048C13.5732 9.31043 13.2359 8.28616 12.4749 7.52515C12.182 7.23225 12.182 6.75738 12.4749 6.46449C12.7678 6.17159 13.2426 6.17159 13.5355 6.46449Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
