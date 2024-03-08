// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Trophy(classNames string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10 1C8.17155 1 6.37729 1.14878 4.62882 1.43504C4.26621 1.4944 4 1.80774 4 2.17518V2.56172C3.17339 2.71855 2.35799 2.90717 1.55514 3.12628C1.23821 3.21277 1.01446 3.49546 1.00306 3.82378C1.00102 3.88231 1 3.94105 1 4C1 6.59485 2.97645 8.72783 5.50636 8.97591C6.27572 9.84484 7.29439 10.4898 8.45156 10.7981C8.35539 11.5844 8.11892 12.3268 7.76796 13H7.5C6.67157 13 6 13.6716 6 14.5V17H5.25C4.55964 17 4 17.5596 4 18.25C4 18.6642 4.33579 19 4.75 19H15.25C15.6642 19 16 18.6642 16 18.25C16 17.5596 15.4404 17 14.75 17H14V14.5C14 13.6716 13.3284 13 12.5 13H12.232C11.8811 12.3268 11.6446 11.5844 11.5484 10.7981C12.7056 10.4898 13.7243 9.84484 14.4936 8.97591C17.0235 8.72783 19 6.59485 19 4C19 3.94103 18.999 3.88229 18.9969 3.82378C18.9855 3.49546 18.7618 3.21277 18.4449 3.12628C17.642 2.90717 16.8266 2.71855 16 2.56172V2.17518C16 1.80774 15.7338 1.4944 15.3712 1.43504C13.6227 1.14878 11.8285 1 10 1ZM2.52524 4.42244C3.01226 4.29976 3.50395 4.18878 4 4.08984V5C4 5.73949 4.13404 6.44825 4.37906 7.10288C3.38067 6.58021 2.66567 5.58968 2.52524 4.42244ZM17.4748 4.42244C17.3343 5.58968 16.6193 6.58021 15.6209 7.10288C15.866 6.44825 16 5.73949 16 5V4.08984C16.496 4.18878 16.9877 4.29976 17.4748 4.42244Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
