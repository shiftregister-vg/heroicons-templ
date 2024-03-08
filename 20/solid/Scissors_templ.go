// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Scissors() templ.Component {
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
		templ_7745c5c3_Err = ScissorsWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ScissorsWithAttrs(attrs templ.Attributes) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1.46876 3.74997C0.502265 5.424 1.07583 7.56456 2.74985 8.53106C4.20093 9.36884 6.00256 9.04946 7.08619 7.85921L7.9695 8.36919C7.99431 8.46181 8.11599 8.48608 8.17843 8.41331C8.32967 8.23707 8.49725 8.07489 8.67914 7.92911C8.96545 7.69963 8.979 7.21997 8.66123 7.03651L7.83619 6.56017C8.32514 5.0266 7.70093 3.30666 6.24985 2.46889C4.57583 1.50239 2.43526 2.07595 1.46876 3.74997ZM3.49985 7.23203C2.54327 6.67974 2.21552 5.45656 2.7678 4.49997C3.32009 3.54339 4.54327 3.21564 5.49985 3.76792C6.45644 4.32021 6.78419 5.54339 6.2319 6.49997C5.67962 7.45656 4.45644 7.78431 3.49985 7.23203Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M9.95556 8.32167C9.17367 8.65155 8.58789 9.32394 8.36825 10.1436L7.96977 11.6308L7.08628 12.1409C6.00265 10.9506 4.20097 10.6311 2.74985 11.4689C1.07583 12.4354 0.502265 14.576 1.46876 16.25C2.43526 17.924 4.57583 18.4976 6.24985 17.5311C7.70089 16.6933 8.32511 14.9735 7.83623 13.4399L18.5151 7.27449C18.7778 7.12282 18.9233 6.82783 18.8837 6.52708C18.8441 6.22633 18.6272 5.97904 18.3342 5.90053L17.631 5.71211C17.0403 5.55383 16.4137 5.59696 15.8503 5.83468L9.95556 8.32167ZM2.7678 15.5C2.21552 14.5434 2.54327 13.3203 3.49985 12.768C4.45644 12.2157 5.67962 12.5434 6.2319 13.5C6.78419 14.4566 6.45644 15.6798 5.49985 16.2321C4.54327 16.7844 3.32009 16.4566 2.7678 15.5Z\" fill=\"#0F172A\"></path> <path d=\"M12.5201 11.8902C12.1624 12.0967 12.1952 12.6233 12.5758 12.7839L15.85 14.1653C16.4134 14.403 17.04 14.4461 17.6307 14.2879L18.3339 14.0994C18.6269 14.0209 18.8438 13.7736 18.8834 13.4729C18.923 13.1721 18.7775 12.8772 18.5148 12.7255L15.0441 10.7217C14.8894 10.6323 14.6988 10.6323 14.5441 10.7217L12.5201 11.8902Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
