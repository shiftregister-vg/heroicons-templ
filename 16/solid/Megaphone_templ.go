// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Megaphone() templ.Component {
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
		templ_7745c5c3_Err = MegaphoneWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MegaphoneWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M13.407 2.58984C13.3168 2.18556 12.916 1.93091 12.5117 2.02107C12.1074 2.11124 11.8528 2.51206 11.9429 2.91635C12.3077 4.55194 12.5003 6.25315 12.5003 8.00023C12.5003 9.74731 12.3077 11.4485 11.9429 13.0841C11.8528 13.4884 12.1074 13.8892 12.5117 13.9794C12.916 14.0696 13.3168 13.8149 13.407 13.4106C13.6713 12.2255 13.8506 11.0086 13.9382 9.76675C14.5699 9.43059 15 8.76552 15 8C15 7.23447 14.5699 6.56939 13.9382 6.23324C13.8505 4.99157 13.6713 3.77483 13.407 2.58984Z\" fill=\"black\"></path> <path d=\"M4.34757 11H4C2.34315 11 1 9.65685 1 8C1 6.34315 2.34315 5 4 5H6C7.64721 5 9.21691 4.66794 10.6459 4.06738C10.8783 5.34142 11 6.6554 11 7.99943C11 9.34372 10.8783 10.658 10.6457 11.9322C9.33751 11.3824 7.91134 11.0578 6.41572 11.0069C6.61941 11.7252 6.89425 12.4137 7.23216 13.0641C7.35143 13.2937 7.28883 13.5787 7.07691 13.7271L6.24892 14.3068C6.01046 14.4738 5.67952 14.4029 5.54168 14.1465C5.01475 13.1663 4.60992 12.1107 4.34757 11Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
