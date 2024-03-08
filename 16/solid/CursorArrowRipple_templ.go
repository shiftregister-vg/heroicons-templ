// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CursorArrowRipple() templ.Component {
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
		templ_7745c5c3_Err = CursorArrowRippleWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CursorArrowRippleWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M4.03774 4.03769C1.98748 6.08794 1.98748 9.41206 4.03774 11.4623C4.33063 11.7552 4.33063 12.2301 4.03774 12.523C3.74484 12.8159 3.26997 12.8159 2.97708 12.523C0.341036 9.88693 0.341036 5.61307 2.97708 2.97703C5.61311 0.34099 9.88698 0.34099 12.523 2.97703C13.8409 4.2949 14.5 6.02369 14.5 7.75C14.5 8.16421 14.1643 8.5 13.75 8.5C13.3358 8.5 13 8.16421 13 7.75C13 6.40525 12.4876 5.06296 11.4624 4.03769C9.4121 1.98744 6.08799 1.98744 4.03774 4.03769Z\" fill=\"black\"></path> <path d=\"M7.71218 7.13596C8.02003 7.05347 8.34639 7.17455 8.52594 7.43786L11.5104 11.8146C11.6779 12.0602 11.6845 12.3817 11.5271 12.634C11.3697 12.8863 11.0781 13.0218 10.7838 12.9794L10.0249 12.8701L10.3128 13.945C10.4199 14.3451 10.1824 14.7563 9.78232 14.8635C9.3822 14.9706 8.97098 14.7331 8.86384 14.333L8.57607 13.2585L7.97381 13.7324C7.74012 13.9163 7.41984 13.9447 7.15739 13.8049C6.89494 13.6651 6.73988 13.3834 6.76212 13.0869L7.1584 7.8043C7.18224 7.48648 7.40433 7.21844 7.71218 7.13596Z\" fill=\"black\"></path> <path d=\"M5.80549 9.69454C4.73155 8.6206 4.73155 6.8794 5.80549 5.80546C6.87943 4.73151 8.62064 4.73151 9.69458 5.80546C10.2317 6.34257 10.5 7.04505 10.5 7.75C10.5 8.16421 10.8358 8.5 11.25 8.5C11.6642 8.5 12 8.16421 12 7.75C12 6.66349 11.585 5.57451 10.7552 4.7448C9.09551 3.08507 6.40456 3.08507 4.74483 4.7448C3.0851 6.40452 3.0851 9.09548 4.74483 10.7552C5.03772 11.0481 5.5126 11.0481 5.80549 10.7552C6.09838 10.4623 6.09838 9.98744 5.80549 9.69454Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
