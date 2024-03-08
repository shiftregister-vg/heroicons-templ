// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CurrencyPound(classNames string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10 18C14.4183 18 18 14.4183 18 10C18 5.58172 14.4183 2 10 2C5.58172 2 2 5.58172 2 10C2 14.4183 5.58172 18 10 18ZM8.73223 6.23223C9.70854 5.25592 11.2915 5.25592 12.2678 6.23223C12.5607 6.52513 13.0355 6.52513 13.3284 6.23223C13.6213 5.93934 13.6213 5.46447 13.3284 5.17157C11.7663 3.60948 9.23367 3.60948 7.67157 5.17157C6.89067 5.95247 6.5 6.97747 6.5 8V8.16481C6.5 8.52945 6.53382 8.89272 6.60072 9.25H6.25C5.83579 9.25 5.5 9.58579 5.5 10C5.5 10.4142 5.83579 10.75 6.25 10.75H6.98724C7.175 11.782 7.05016 12.853 6.61952 13.8219L6.56464 13.9454C6.44963 14.2042 6.49018 14.5057 6.66948 14.7249C6.84878 14.9441 7.13625 15.0436 7.4127 14.9821L8.68531 14.6993C9.21441 14.5818 9.76361 14.589 10.2894 14.7204C11.0842 14.9191 11.9158 14.9191 12.7106 14.7204L13.6819 14.4776C14.0837 14.3771 14.3281 13.9699 14.2276 13.5681C14.1271 13.1663 13.7199 12.9219 13.3181 13.0224L12.3468 13.2652C11.7908 13.4042 11.2092 13.4042 10.6532 13.2652C9.91296 13.0802 9.14026 13.0673 8.39452 13.2275C8.58441 12.4159 8.62237 11.5757 8.5063 10.75H9.75C10.1642 10.75 10.5 10.4142 10.5 10C10.5 9.58579 10.1642 9.25 9.75 9.25H8.13603C8.0458 8.89574 8 8.53121 8 8.16481V8C8 7.35903 8.24393 6.72054 8.73223 6.23223Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
