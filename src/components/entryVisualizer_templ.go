// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"strings"
)

type Entry struct {
	Word string  `json:"word"`
	Time float32 `json:"time"`
	Freq string  `json:"freq"`
	Aoa  string  `json:"aoa"`
}

func listEntries(entries []Entry, id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_listEntries_91fe`,
		Function: `function __templ_listEntries_91fe(entries, id){const svg = d3.select("#"+id).select("svg")
  let g = svg.append("g")

  const margin = 25
  const width = svg.attr("width")
  const height = svg.attr("height")

  const maxVal = d3.max(entries, (d) => d.time)

  const y = d3.scaleBand()
    .domain(d3.groupSort(entries, ([d]) => -d.time, (d) => d.word))
    .range([margin, height-margin])
    .padding(.2)

  const x = d3.scaleLinear().domain([-maxVal, maxVal]).range([margin,width-margin])
  
  g.selectAll().data(entries)
     .join("rect")
        .attr("y", (d) => y(d.word))
        .attr("height", y.bandwidth())
        .attr("width", (d) => x(d.time)-x(0))
        .append("svg:title").text((d) => d.word)
  g.selectAll("text").data(entries).enter()
      .append("text")
        .text((d) => d.time)
        .attr("y", (d) => y(d.word)+y.bandwidth()-1)
        .style("fill", "white")
        .style("font-size", y.bandwidth()+"px")

  const update = () => {
    if(d3.select("#" + id + "aoa").property("checked")) {
      g.selectAll("rect").data(entries)
      .attr("fill", (d) => d3.hsl(180, 0.75, d.aoa === "early" ? 0.6 : 0.4).formatHex())
      .attr("x", (d) => {
        if(d.aoa === "early") {
          return x(-d.time)
        } else {
          return x(0)
        }
      })
      g.selectAll("text").data(entries)
        .attr("x", (d) => {
          if(d.aoa === "early") {
            return x(-d.time)+10
          }
          else {
            return x(d.time)-10
          }
          })
        .style("text-anchor", (d) => d.aoa === "early" ? "start" : "end")
    } else {
      g.selectAll("rect").data(entries)
      .attr("fill", (d) => d3.hsl(20, 0.75, d.freq === "high" ? 0.6 : 0.4).formatHex())
      .attr("x", (d) => {
        if(d.freq === "low") {
          return x(-d.time)
        } else {
          return x(0)
        }
      })
      g.selectAll("text").data(entries)
        .attr("x", (d) => {
          if(d.freq === "low") {
            return x(-d.time)+10
          }
          else {
            return x(d.time)-10
          }
          })
        .style("text-anchor", (d) => d.freq === "low" ? "start" : "end")
    }
  }


  d3.select("#" + id + "aoa").on("change", update  )
  update()




}`,
		Call:       templ.SafeScript(`__templ_listEntries_91fe`, entries, id),
		CallInline: templ.SafeScriptInline(`__templ_listEntries_91fe`, entries, id),
	}
}

func titlize(title string) string {
	return strings.Join((strings.Split(title, " ")), "_") + "_listdiv"
}

func EntryVisualizer(entries []Entry, title string, width int) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(titlize(title)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"rounded-2xl shadow-md m-2 p-3 bg-white flex flex-col\"><span class=\"text-xl font-bold w-full text-right \">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/entryVisualizer.templ`, Line: 103, Col: 60}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span><div class=\"flex justify-end w-full\"><label for=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(titlize(title) + "aoa"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"font-bold mr-2\">Distinguish AoA</label> <input type=\"checkbox\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(titlize(title) + "aoa"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><svg width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprint(width)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"700\"></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = listEntries(entries, titlize(title)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
