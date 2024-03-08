// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CodeBracket(classNames string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M6.28033 5.21979C6.57322 5.51269 6.57322 5.98756 6.28033 6.28045L2.56066 10.0001L6.28033 13.7198C6.57322 14.0127 6.57322 14.4876 6.28033 14.7805C5.98744 15.0733 5.51256 15.0733 5.21967 14.7805L0.96967 10.5305C0.676777 10.2376 0.676777 9.76269 0.96967 9.46979L5.21967 5.21979C5.51256 4.9269 5.98744 4.9269 6.28033 5.21979ZM13.7197 5.21979C14.0126 4.9269 14.4874 4.9269 14.7803 5.21979L19.0303 9.46979C19.3232 9.76269 19.3232 10.2376 19.0303 10.5305L14.7803 14.7805C14.4874 15.0733 14.0126 15.0733 13.7197 14.7805C13.4268 14.4876 13.4268 14.0127 13.7197 13.7198L17.4393 10.0001L13.7197 6.28045C13.4268 5.98756 13.4268 5.51269 13.7197 5.21979Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.3774 2.01103C11.7856 2.0814 12.0595 2.46936 11.9891 2.87755L9.48909 17.3776C9.41872 17.7857 9.03076 18.0596 8.62257 17.9892C8.21438 17.9188 7.94053 17.5309 8.0109 17.1227L10.5109 2.62269C10.5813 2.2145 10.9692 1.94065 11.3774 2.01103Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
