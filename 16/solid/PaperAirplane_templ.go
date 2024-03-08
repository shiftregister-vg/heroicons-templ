// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PaperAirplane(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M2.86908 2.29797C2.60248 2.25655 2.33418 2.3615 2.16646 2.57282C1.99873 2.78414 1.95744 3.06927 2.05831 3.31949L3.39039 6.62389C3.54291 7.00223 3.90994 7.25 4.31787 7.25H8.25C8.66421 7.25 9 7.58579 9 8C9 8.41421 8.66421 8.75 8.25 8.75H4.31812C3.9102 8.75 3.54317 8.99777 3.39065 9.37611L2.05831 12.6811C1.95744 12.9314 1.99873 13.2165 2.16646 13.4278C2.33418 13.6391 2.60248 13.7441 2.86908 13.7027C7.23525 13.0242 11.2273 11.2161 14.5368 8.58762C14.7159 8.44534 14.8204 8.22908 14.8204 8.00032C14.8204 7.77155 14.7159 7.55529 14.5368 7.41302C11.2273 4.78455 7.23525 2.97644 2.86908 2.29797Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
