// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func HomeModern(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M10.5358 3.44371C10.9188 3.28599 11.1014 2.84765 10.9437 2.46463C10.786 2.08162 10.3476 1.89898 9.96463 2.05669L3.5 4.7186V3.75C3.5 3.33579 3.16421 3 2.75 3C2.33579 3 2 3.33579 2 3.75V5.33624L1.46463 5.55669C1.08162 5.7144 0.898977 6.15274 1.05669 6.53576C1.20951 6.9069 1.62585 7.0899 2 6.95739V12.5H1.75C1.33579 12.5 1 12.8358 1 13.25C1 13.6642 1.33579 14 1.75 14H4C4.55228 14 5 13.5523 5 13V12C5 11.4477 5.44772 11 6 11C6.55228 11 7 11.4477 7 12V13C7 13.5523 7.44772 14 8 14H9C9.55228 14 10 13.5523 10 13V3.66431L10.5358 3.44371Z\" fill=\"black\"></path> <path d=\"M11.8285 5.80164C11.6206 5.94087 11.4958 6.1746 11.4958 6.42481V8.5C11.4958 8.52688 11.4972 8.55344 11.5 8.57959V13C11.5 13.5523 11.9477 14 12.5 14H13C13.5523 14 14 13.5523 14 13V7.95717C14.3742 8.08968 14.7905 7.90668 14.9433 7.53554C15.101 7.15252 14.9184 6.71418 14.5354 6.55647L12.5314 5.7313C12.3 5.63603 12.0364 5.66242 11.8285 5.80164Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
