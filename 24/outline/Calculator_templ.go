// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package outline

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Calculator(classNames string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M15.75 15.75V18M8.25 11.25H8.2575V11.2575H8.25V11.25ZM8.25 13.5H8.2575V13.5075H8.25V13.5ZM8.25 15.75H8.2575V15.7575H8.25V15.75ZM8.25 18H8.2575V18.0075H8.25V18ZM10.7476 11.25H10.7551V11.2575H10.7476V11.25ZM10.7476 13.5H10.7551V13.5075H10.7476V13.5ZM10.7476 15.75H10.7551V15.7575H10.7476V15.75ZM10.7476 18H10.7551V18.0075H10.7476V18ZM13.2524 11.25H13.2599V11.2575H13.2524V11.25ZM13.2524 13.5H13.2599V13.5075H13.2524V13.5ZM13.2524 15.75H13.2599V15.7575H13.2524V15.75ZM13.2524 18H13.2599V18.0075H13.2524V18ZM15.75 11.25H15.7575V11.2575H15.75V11.25ZM15.75 13.5H15.7575V13.5075H15.75V13.5ZM8.25 6H15.75V8.25H8.25V6ZM12 2.25C10.108 2.25 8.24156 2.35947 6.40668 2.57241C5.30608 2.70014 4.5 3.649 4.5 4.75699V19.5C4.5 20.7426 5.50736 21.75 6.75 21.75H17.25C18.4926 21.75 19.5 20.7426 19.5 19.5V4.75699C19.5 3.649 18.6939 2.70014 17.5933 2.57241C15.7584 2.35947 13.892 2.25 12 2.25Z\" stroke=\"#0F172A\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
