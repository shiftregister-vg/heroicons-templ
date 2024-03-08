// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GlobeAmericas(classNames string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M18 10C18 14.4183 14.4183 18 10 18C5.58172 18 2 14.4183 2 10C2 5.58172 5.58172 2 10 2C14.4183 2 18 5.58172 18 10ZM16.5 10C16.5 13.5899 13.5899 16.5 10 16.5C6.41015 16.5 3.5 13.5899 3.5 10C3.5 8.15664 4.26733 6.4925 5.5 5.30957V5.75736C5.5 6.68562 5.86875 7.57586 6.52513 8.23223L8.29289 10L8 10.2929C7.60947 10.6834 7.60948 11.3166 8 11.7071L9.06066 12.7678C9.34197 13.0491 9.5 13.4306 9.5 13.8284V14.191C9.5 14.5698 9.714 14.916 10.0528 15.0854L10.3292 15.2236C10.8232 15.4706 11.4238 15.2704 11.6708 14.7764L13.1249 11.8683C13.4136 11.2908 13.3004 10.5933 12.8439 10.1368L12.0721 9.36495C11.8042 9.09712 11.4081 9.0036 11.0487 9.12338L10.6647 9.2514C10.4286 9.33008 10.1706 9.22322 10.0593 9.00066L9.76344 8.40885C9.67084 8.22365 9.70714 7.99997 9.85355 7.85355C9.99997 7.70714 10.2236 7.67084 10.4088 7.76344L10.6708 7.89443C10.8097 7.96385 10.9628 8 11.118 8H11.3063C11.9888 8 12.4708 7.3313 12.255 6.68377L12.187 6.47978C12.1227 6.28692 12.1816 6.07434 12.3359 5.94204L13.7754 4.70821C15.4248 5.8871 16.5 7.81799 16.5 10Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
