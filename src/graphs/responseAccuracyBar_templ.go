// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package graphs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"math"
)

type responseAccuracyData struct {
	NonWordAccuracy float64
	NonWordError    float64
	WordAccuracy    float64
	WordError       float64
}

func getResponseAccuracyBarData(db *pgxpool.Pool) (*responseAccuracyData, error) {
	d := responseAccuracyData{}
	val, err := db.Query(context.Background(), `select avg(col), stddev(col), count(col) from (select 1 as col from results where word = true and correct = true
    union all
    select 0 as col from results where word = true and correct = false) as s
    union select avg(col), stddev(col), count(col) from (select 1 as col from results where word = false and correct = true
    union all
    select 0 as col from results where word = false and correct = false) as s`)
	defer val.Close()
	if err != nil {
		return nil, err
	}
	if val.Next() == false {
		return nil, &GraphError{}
	}
	var avg float64
	var stddev float64
	var count int
	val.Scan(&avg, &stddev, &count)
	d.WordAccuracy = avg
	d.WordError = stddev / math.Sqrt(float64(count))
	if val.Next() == false {
		return nil, &GraphError{}
	}
	val.Scan(&avg, &stddev, &count)
	d.NonWordAccuracy = avg
	d.NonWordError = stddev / math.Sqrt(float64(count))
	return &d, nil
}

func ResponseAccuracyBar(db *pgxpool.Pool) templ.Component {
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
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 1)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = graphBase(
			"Figure 1",
			"The average accuracy of users at identifying words and nonwords.",
			displayResponseAccuracyBar(getResponseAccuracyBarData((db))),
		).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func displayResponseAccuracyBar(d *responseAccuracyData, err error) templ.Component {
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
		if err != nil {
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 3)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 4)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = svgComponent("0 0 800 500", "response-accuracy-graph").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 5)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderResponseAccuracyGraph(d.WordAccuracy, d.WordError, d.NonWordAccuracy, d.NonWordError).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func renderResponseAccuracyGraph(wa float64, we float64, na float64, ne float64) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_renderResponseAccuracyGraph_88c8`,
		Function: `function __templ_renderResponseAccuracyGraph_88c8(wa, we, na, ne){let svg = d3.select("#response-accuracy-graph"),
      margin = 200,
      height = 500-margin,
      width = 800-margin

  let xScale = d3.scaleBand().range([0,width]).padding(0.4),
      yScale = d3.scaleLinear().range([height, 0])

  let g = svg.append("g")
            .attr("transform", "translate(100, 100)")


  // applying scale
  xScale.domain(["Word", "Nonword"])
  yScale.domain([0, 1])
  g.append("g")
    .attr("transform", "translate(0,"+height+")")
    .call(d3.axisBottom(xScale))
  g.append("g")
    .call(d3.axisLeft(yScale).tickFormat((d) => {
      return d
    }).ticks(10))

  // generating tooltip
  let div = d3.select("body")
    .append("div")

  div.attr("id", "acctt")
    .attr("class", "drop-shadow-md text-caption rounded-md p-1 border-1 border-grey bg-white absolute flex flex-col")
    .style("visibility", "hidden")
  svg.on('mousemove', function(event) {
      div.style("left", event.pageX + 10)
         .style("top", event.pageY + 10)
    });


  // applying data
  g.selectAll(".bar")
    .data([["Word", wa, we], ["Nonword", na, ne]])
    .enter()
    .append("rect")
    .attr("class", "bar")
    .attr("x", (d) => xScale(d[0]))
    .attr("y", (d) => yScale(d[1]))
    .attr("width", xScale.bandwidth())
    .attr("height", (d) => (height - yScale(d[1])))
    .style("cursor", "default")
    .on('mouseenter', function(event, d) {
      d3.select(this).transition().duration(100).style("opacity", 0.85)
      let avg = d[1].toFixed(3)
      let err = d[2].toFixed(3)
      div.html(` + "`" + `
        <div class="flex justify-between">
        <span class="mr-1">Mean:</span><span>${avg}</span>
        </div>
        <div class="flex justify-between">
        <span class="mr-1">Error:</span><span>${err}</span>
        </div>
      ` + "`" + `)
      div.style("visibility", "visible")
    })
    .on('mouseout', function(event, d) {
      d3.select(this).transition().duration(100).style("opacity", 1)
      div.style("visibility", "hidden")
    })
  // adding error bars
  let lines = g.append("g").selectAll("line.err")
    .data([["Word", wa, we], ["Nonword", na, ne]])
  lines.enter()
    .append("line")
    .attr("class", "err")
    .attr('x1', d => xScale(d[0])+(xScale.bandwidth()/2))
    .attr('x2', d => xScale(d[0])+(xScale.bandwidth()/2))
    .attr('y1', d => yScale(d[1]+d[2]))
    .attr('y2', d => yScale(d[1]-d[2]))
    .style('stroke-width', xScale.bandwidth()/50)
  lines.enter()
    .append("line")
    .attr("class", "err")
    .attr('x1', d => xScale(d[0])+(xScale.bandwidth()/2)-(xScale.bandwidth()/10))
    .attr('x2', d => xScale(d[0])+(xScale.bandwidth()/2)+(xScale.bandwidth()/10))
    .attr('y1', d => yScale(d[1]+d[2]))
    .attr('y2', d => yScale(d[1]+d[2]))
  lines.enter()
    .append("line")
    .attr("class", "err")
    .attr('x1', d => xScale(d[0])+(xScale.bandwidth()/2)-(xScale.bandwidth()/10))
    .attr('x2', d => xScale(d[0])+(xScale.bandwidth()/2)+(xScale.bandwidth()/10))
    .attr('y1', d => yScale(d[1]-d[2]))
    .attr('y2', d => yScale(d[1]-d[2]))

}`,
		Call:       templ.SafeScript(`__templ_renderResponseAccuracyGraph_88c8`, wa, we, na, ne),
		CallInline: templ.SafeScriptInline(`__templ_renderResponseAccuracyGraph_88c8`, wa, we, na, ne),
	}
}