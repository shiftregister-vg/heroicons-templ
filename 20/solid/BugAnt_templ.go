// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BugAnt(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M6.56061 1.13947C6.8982 1.37948 6.9773 1.84772 6.73729 2.18531C6.54915 2.44994 6.39257 2.73816 6.27282 3.04456C6.45824 3.21553 6.6554 3.37388 6.86291 3.51827C7.59496 2.59415 8.72765 2 10 2C11.2723 2 12.405 2.59414 13.1371 3.51826C13.3446 3.37388 13.5418 3.21552 13.7272 3.04455C13.6074 2.73816 13.4509 2.44994 13.2627 2.18531C13.0227 1.84772 13.1018 1.37948 13.4394 1.13947C13.777 0.899458 14.2452 0.97856 14.4852 1.31615C14.8539 1.83471 15.1353 2.42054 15.3072 3.05195C15.3752 3.30183 15.3095 3.56913 15.1334 3.75901C14.7474 4.17528 14.3111 4.5446 13.8342 4.85737C13.9421 5.2199 14 5.6036 14 6C14 6.52014 13.6988 6.96316 13.277 7.18697C12.9093 7.38211 12.5221 7.54541 12.1193 7.67324C12.2496 7.88095 12.3502 8.10918 12.4152 8.35198C13.8282 8.17751 15.194 7.85193 16.4961 7.39199C16.4987 7.26165 16.5 7.13098 16.5 7C16.5 6.34959 16.4682 5.70683 16.4061 5.07318C16.3656 4.66094 16.6671 4.29399 17.0793 4.25358C17.4915 4.21316 17.8585 4.51458 17.8989 4.92682C17.9658 5.60904 18 6.30064 18 7C18 7.31999 17.9928 7.63836 17.9787 7.95496C17.9653 8.25331 17.7762 8.51531 17.4972 8.62195C15.867 9.24519 14.1395 9.67109 12.3441 9.87057C12.3126 9.95523 12.2767 10.0377 12.2367 10.1178C12.8138 10.1789 13.384 10.2634 13.9464 10.3704C15.1681 10.6027 16.3525 10.9409 17.4895 11.3751C17.7934 11.4911 17.9879 11.79 17.9709 12.1149C17.8783 13.8835 17.5668 15.5948 17.0628 17.2219C16.9402 17.6176 16.5201 17.839 16.1244 17.7164C15.7287 17.5938 15.5074 17.1737 15.6299 16.7781C16.0452 15.4376 16.3198 14.0349 16.4364 12.587C15.941 12.4139 15.4366 12.2602 14.9239 12.1268C14.9739 12.4106 15 12.7024 15 13C15 14.8137 14.4833 16.312 13.5738 17.3693C12.6592 18.4326 11.3874 19 10 19C8.61265 19 7.34085 18.4326 6.42621 17.3693C5.51672 16.312 5 14.8137 5 13C5 12.7024 5.02608 12.4106 5.07611 12.1268C4.56343 12.2602 4.05896 12.4139 3.56361 12.587C3.68017 14.0349 3.95479 15.4376 4.37006 16.7781C4.49264 17.1737 4.27125 17.5938 3.87559 17.7164C3.47993 17.839 3.05982 17.6176 2.93724 17.2219C2.43318 15.5948 2.12168 13.8835 2.02911 12.1149C2.0121 11.79 2.20656 11.4911 2.51054 11.3751C3.6475 10.9409 4.83186 10.6027 6.05361 10.3704C6.61601 10.2634 7.18625 10.1789 7.76335 10.1178C7.72328 10.0377 7.68738 9.95524 7.65593 9.87058C5.86054 9.6711 4.13298 9.2452 2.50277 8.62196C2.2238 8.51531 2.0347 8.25332 2.02134 7.95496C2.00716 7.63837 2 7.31999 2 7C2 6.30064 2.03421 5.60904 2.1011 4.92682C2.14151 4.51458 2.50846 4.21316 2.9207 4.25358C3.33293 4.29399 3.63436 4.66094 3.59394 5.07318C3.53182 5.70683 3.5 6.34959 3.5 7C3.5 7.13098 3.50129 7.26165 3.50386 7.392C4.80604 7.85194 6.17184 8.17751 7.5848 8.35198C7.64979 8.10918 7.75041 7.88095 7.88072 7.67324C7.47789 7.54541 7.09072 7.38211 6.72298 7.18697C6.30121 6.96315 6 6.52014 6 6C6 5.6036 6.05787 5.2199 6.16579 4.85736C5.68885 4.5446 5.25263 4.17528 4.86656 3.75901C4.69045 3.56913 4.62477 3.30183 4.69281 3.05195C4.86474 2.42054 5.14609 1.83472 5.51477 1.31615C5.75478 0.97856 6.22302 0.899458 6.56061 1.13947Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
