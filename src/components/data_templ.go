// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"

func percDeadVisualizer(dead int, total int) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_percDeadVisualizer_3f91`,
		Function: `function __templ_percDeadVisualizer_3f91(dead, total){const svg = d3.select("#pdv").select("svg")
    let width = svg.attr("width"),
        height = svg.attr("height"),
        radius = Math.min(width, height) / 2,
        g = svg.append("g").attr("transform", "translate(" + width / 2 + "," + height / 2 + ")")

  const pie = d3.pie()
  pie.padAngle(0.03)

  const arc = d3.arc().innerRadius(radius * 0.8).outerRadius(radius)

var color = d3.scaleOrdinal(['#4daf4a','#377eb8','#ff7f00','#984ea3','#e41a1c']);

  const arcs = g.selectAll("arc")
    .data(pie([dead, total-dead]))
    .enter()
    .append("g")
    .attr("inactive", "active")

  arcs.append("path")
    .attr("fill", function(d, i) {
      return color(i)
      })
    .attr("d", arc)

  arcs.append("text")
    .attr("transform", function(d){
        let pos = arc.centroid(d);
        let midangle = d.startAngle + (d.endAngle - d.startAngle) / 2
        pos[0] = radius * 0.4 * (midangle < Math.PI ? 1 : -1);
        return 'translate(' + pos + ')';
    })
    .attr("text-anchor", "middle")
    .text( function(d, i) {
      return [dead + " inactive", (total-dead) + " active"][i]}
    );







}`,
		Call:       templ.SafeScript(`__templ_percDeadVisualizer_3f91`, dead, total),
		CallInline: templ.SafeScriptInline(`__templ_percDeadVisualizer_3f91`, dead, total),
	}
}

func Data(resCount int, users int, deads int, complete int, incorrects templ.Component, allEntries templ.Component) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><title>eels AAA data</title><script src=\"https://cdn.jsdelivr.net/npm/d3@7\"></script><script src=\"https://cdn.tailwindcss.com\"></script><script src=\"https://code.jquery.com/jquery-3.7.1.slim.min.js\" integrity=\"sha256-kmHvs0B+OpCW5GVHUNjv9rOmY0IvSIRcf7zGUDTDQM8=\" crossorigin=\"anonymous\"></script></head><body class=\"bg-gray-100\"><div class=\"flex justify-between\"><div class=\"flex\"><div id=\"pdv\" class=\"rounded-2xl shadow-md m-2 w-min h-min p-3 bg-white\"><span class=\"text-xl text-center font-bold\">Users</span> <svg width=\"300\" height=\"300\"></svg></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = percDeadVisualizer(deads, users).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><div id=\"responses\" class=\"rounded-2xl shadow-md m-2 p-3 bg-white\"><span class=\"text-xl text-center font-bold justify-left\">Responses</span><div class=\"flex items-center flex-col\"><span class=\"text-xl\"><b>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprint(resCount))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/data.templ`, Line: 71, Col: 34}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</b> responses<br>from <b>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprint(users - deads))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/data.templ`, Line: 73, Col: 42}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</b> users.<br><br>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprint(complete))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/data.templ`, Line: 76, Col: 31}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" users finished the test.</span></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = incorrects.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = allEntries.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
