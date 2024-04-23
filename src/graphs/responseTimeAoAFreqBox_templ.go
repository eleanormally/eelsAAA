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
)

func ResponseTimeAoAFreqBox(db *pgxpool.Pool) templ.Component {
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
			"Figure 5",
			"Blue box surrounds the center 2 quartiles. Whiskers surround the 5th and 9th percentiles. All data the falls outside of this is plotted",
			displayResponseTimeAoAFreqBox(getResponseTimeAoAFreqData(db)),
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

type responseTimeAoAFreqData struct {
	EarlyLow  responseEntryData
	EarlyHigh responseEntryData
	LateLow   responseEntryData
	LateHigh  responseEntryData
}

func getResponseTimeAoAFreqData(db *pgxpool.Pool) (*responseTimeAoAFreqData, error) {
	var out responseTimeAoAFreqData
	sVal, err := db.Query(context.Background(),
		`select 
	aoa,
  freq,
	min(CASE WHEN quartile = 2 and word = true then time end) as q2, 
	max(CASE WHEN quartile = 3 and word = true then time end) as q3,
  min(CASE WHEN percentile >= 0.05 and word = true then time end) as p5,
  max(CASE WHEN percentile <= 0.95 and word = true then time end) as p95,
  avg(CASE WHEN word = true then time end) mean,
  percentile_cont(0.5) within group (order by time) filter(WHERE word = true) as med,
  stddev(CASE WHEN word = true then time end) std
from (select
	r.time,
  wp.aoa,
  wp.freq,
  ntile(4) over (partition by r.word order by r.time) as quartile,
  PERCENT_RANK() over (partition by r.word order by r.time) as percentile,
 	r.word

from results as r join "wordPairs" as wp on r.pair_id = wp.id
order by time
) as s group by aoa, freq`)
	defer sVal.Close()
	if err != nil {
		return nil, err
	}
	for i := 0; i < 4; i++ {
		if sVal.Next() == false {
			return nil, &GraphError{}
		}
		var d responseEntryData
		var s1 string
		var s2 string
		sVal.Scan(
			&s1,
			&s2,
			&d.SecondQuartile,
			&d.ThirdQuartile,
			&d.FifthPc,
			&d.NineFifthPc,
			&d.Mean,
			&d.Median,
			&d.StdDev,
		)
		switch s1 + s2 {
		case "earlyhigh":
			out.EarlyHigh = d
		case "earlylow":
			out.EarlyLow = d
		case "latehigh":
			out.LateHigh = d
		case "latelow":
			out.LateLow = d
		}
	}

	eVal, err := db.Query(context.Background(),
		`
select 
			wp.aoa,
      wp.freq,
      r.time, 
      r.word, 
      wp.word 
    from results as r join "wordPairs" as wp on r.pair_id = wp.id 
    where r.word = true and 
    	wp.aoa = 'early' and
      wp.freq = 'high' and 
          (r.time < $1
            or
           r.time > $2
          )
    union all
 select 
			wp.aoa,
      wp.freq,
      r.time, 
      r.word, 
      wp.word 
    from results as r join "wordPairs" as wp on r.pair_id = wp.id 
    where r.word = true and 
    	wp.aoa = 'late' and
      wp.freq = 'high' and
          (r.time < $3
            or
           r.time > $4
          )
    union all
    select 
			wp.aoa,
      wp.freq,
      r.time, 
      r.word, 
      wp.word 
    from results as r join "wordPairs" as wp on r.pair_id = wp.id 
    where r.word = true and 
    	wp.aoa = 'early' and
      wp.freq = 'low' and
          (r.time < $5
            or
           r.time > $6
          )
    union all
    select 
			wp.aoa,
      wp.freq,
      r.time, 
      r.word, 
      wp.word 
    from results as r join "wordPairs" as wp on r.pair_id = wp.id 
    where r.word = true and 
    	wp.aoa = 'late' and
      wp.freq = 'low' and
          (r.time < $7
            or
           r.time > $8
          )
`,
		out.EarlyHigh.FifthPc,
		out.EarlyHigh.NineFifthPc,
		out.LateHigh.FifthPc,
		out.LateHigh.NineFifthPc,
		out.EarlyLow.FifthPc,
		out.EarlyLow.NineFifthPc,
		out.LateLow.FifthPc,
		out.LateLow.NineFifthPc,
	)
	defer eVal.Close()
	if err != nil {
		return nil, err
	}
	for eVal.Next() {
		var aoa string
		var freq string
		var time int
		var isWord bool
		var word string
		eVal.Scan(&aoa, &freq, &time, &isWord, &word)
		if !isWord {
			continue
		}
		switch aoa + freq {
		case "earlyhigh":
			out.EarlyHigh.Times = append(out.EarlyHigh.Times, response{
				Time: time,
				Word: word,
			})
		case "earlylow":
			out.EarlyLow.Times = append(out.EarlyLow.Times, response{
				Time: time,
				Word: word,
			})
		case "latehigh":
			out.LateHigh.Times = append(out.LateHigh.Times, response{
				Time: time,
				Word: word,
			})
		case "latelow":
			out.LateLow.Times = append(out.LateLow.Times, response{
				Time: time,
				Word: word,
			})
		}
	}
	return &out, nil
}

func renderResponseTimeAoAFreqBox(d *responseTimeAoAFreqData) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_renderResponseTimeAoAFreqBox_a3cf`,
		Function: `function __templ_renderResponseTimeAoAFreqBox_a3cf(d){let svg = d3.select("#response-time-aoa-freq-box"),
    margin = 200,
    height = 1200-margin,
    width = 800-margin

  let data = [
    ["Early Low", d.EarlyLow],
    ["Early High", d.EarlyHigh],
    ["Late Low", d.LateLow],
    ["Late High", d.LateHigh],
  ]
  let outlierData = [
    ...d.LateHigh.Times.map(d => ["Early Low", d]),
    ...d.EarlyHigh.Times.map(d => ["Early High", d]),
    ...d.LateLow.Times.map(d => ["Late Low", d]),
    ...d.LateHigh.Times.map(d => ["Late High", d]),
    ]

  let xScale = d3.scaleBand().range([0, width]).padding(0.4),
      yScale = d3.scaleLinear().range([height, 0])

  let g = svg.append("g")
            .attr("transform", "translate(100, 100)")

  let tt = d3.select("body")
        .append("div")
  tt.attr("id", "timeaftt")
    .attr("class", "drop-shadow-md text-sm rounded-md p-1 border-1 border-grey bg-white flex flex-col")
    .style("position", "absolute")
    .style("visibility", "hidden")
  svg.on('mousemove', function(event) {
      tt.style("left", event.pageX - 10)
         .style("top", event.pageY + 10)
    });


  // generate scale
  xScale.domain(["Early Low", "Early High", "Late Low", "Late High"])
  const outlierVals = outlierData.flat().filter(d => typeof d != "string").map(d => d.Time)
  const outlierMax = Math.max(...outlierVals)
  yScale.domain([0, Math.ceil(outlierMax/1000)*1000])
  g.append("g")
    .attr("transform", "translate(0,"+height+")")
    .call(d3.axisBottom(xScale))
  g.append("g")
    .call(d3.axisLeft(yScale).tickFormat(d => d).ticks(20))


  //apply median box
  g.selectAll(".bar").data(data).enter()
    .append("rect")
    .attr("class", "bar")
    .attr("x", d => xScale(d[0]))
    .attr("y", d => yScale(d[1].ThirdQuartile))
    .attr("width", xScale.bandwidth())
    .attr("height", d => yScale(d[1].SecondQuartile)-yScale(d[1].ThirdQuartile))
    .on("mouseover", function(event, d) {
      d3.select(this).transition().duration(100).style("opacity", 0.85)
      tt.html(` + "`" + `
        <div class="flex justify-between">
          <span class="mr-1">Mean:</span><span>${d[1].Mean.toFixed(1)}ms</span>
        </div>
        <div class="flex justify-between">
          <span class="mr-1">Median:</span><span>${d[1].Median.toFixed(1)}ms</span>
        </div>
        <div class="flex justify-between">
          <span class="mr-1">Std Dev:</span><span>${d[1].StdDev.toFixed(1)}ms</span>
        </div>
        <div class="flex justify-between overflow-visible">
          <span class="mr-1">Center Quartiles:</span><span>${d[1].SecondQuartile.toFixed(0)}-${d[1].ThirdQuartile.toFixed(0)}ms</span>
        </div>
        <div class="flex justify-between overflow-visible">
          <span class="mr-1">5-95 Percentiles:</span><span>${d[1].FifthPc.toFixed(0)}-${d[1].NineFifthPc.toFixed(0)}ms</span>
        </div>
      ` + "`" + `)
      tt.style("visibility", "visible")
    })
    .on("mouseout", function(event, d) {
      d3.select(this).transition().duration(100).style("opacity", 1)
      tt.style("visibility", "hidden")
    })
  //apply error bars
  let lines = g.append("g").style("pointer-events", "none").selectAll("line.err").data(data)
  lines.enter()
    .append("line")
    .attr("class", "err")
    .attr('x1', d => xScale(d[0])+(xScale.bandwidth()/2))
    .attr('x2', d => xScale(d[0])+(xScale.bandwidth()/2))
    .attr('y1', d => yScale(d[1].FifthPc))
    .attr('y2', d => yScale(d[1].NineFifthPc))
    .style('stroke-width', xScale.bandwidth()/50)

  lines.enter()
    .append("line")
    .attr("class", "err")
    .attr('x1', d => xScale(d[0])+(xScale.bandwidth()/2)-(xScale.bandwidth()/10))
    .attr('x2', d => xScale(d[0])+(xScale.bandwidth()/2)+(xScale.bandwidth()/10))
    .attr('y1', d => yScale(d[1].FifthPc))
    .attr('y2', d => yScale(d[1].FifthPc))
  lines.enter()
    .append("line")
    .attr("class", "err")
    .attr('x1', d => xScale(d[0])+(xScale.bandwidth()/2)-(xScale.bandwidth()/10))
    .attr('x2', d => xScale(d[0])+(xScale.bandwidth()/2)+(xScale.bandwidth()/10))
    .attr('y1', d => yScale(d[1].NineFifthPc))
    .attr('y2', d => yScale(d[1].NineFifthPc))
  //apply median bar
  lines.enter()
    .append("line")
    .attr("class", "err")
    .attr('x1', d => xScale(d[0]))
    .attr('x2', d => xScale(d[0])+xScale.bandwidth())
    .attr('y1', d => yScale(d[1].Median))
    .attr('y2', d => yScale(d[1].Median))
  //apply mean cross
  g.selectAll('.cross').data(data).enter()
    .append("path")
    .style("pointer-events", "none")
    .attr("class", "cross")
    .attr("d", d3.symbol(d3.symbolTimes).size(200))
    .style("stroke", "black")
    .style("stroke-width", "1")
    .attr('transform', d => (
      ` + "`" + `translate(${xScale(d[0])+(xScale.bandwidth()/2)},${yScale(d[1].Mean)})` + "`" + `
    ))


  //display outliers
  let outliers = g.append("g").selectAll(".outlier").data(outlierData)
  outliers.enter()
    .append("circle")
    .attr("class", "outlier")
    .attr("cx", d=> xScale(d[0])+(xScale.bandwidth()/2))
    .attr("cy", d=> yScale(d[1].Time))
    .attr("r", 5)
    .attr("id", d => d[1].Time + d[1].Word)
    .attr("fill", "orange")
    .on("mouseover", function(event, d) {
      d3.select(this).raise()
      d3.select(this).transition().duration(50).attr("r", 10).attr("fill", "#3366ff")

      tt.html(` + "`" + `
        <div class="flex justify-between">
          <span class="mr-1">Time:</span><span>${d[1].Time}ms</span>
        </div>
        <div class="flex justify-between">
          <span class="mr-1">${d[0]}:</span><span>${d[1].Word}</span>
        </div>
        <span>${d[0]}</span>
      ` + "`" + `)
      tt.style("visibility", "visible")
    })
    .on("mouseout", function(event, d) {
      d3.select(this).transition().duration(50).attr("r", 5).attr("fill", "orange")
      tt.style("visibility", "hidden")
    })

}`,
		Call:       templ.SafeScript(`__templ_renderResponseTimeAoAFreqBox_a3cf`, d),
		CallInline: templ.SafeScriptInline(`__templ_renderResponseTimeAoAFreqBox_a3cf`, d),
	}
}

func displayResponseTimeAoAFreqBox(d *responseTimeAoAFreqData, err error) templ.Component {
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
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(err.Error())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `graphs/responseTimeAoAFreqBox.templ`, Line: 354, Col: 44}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 4)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = svgComponent("0 0 800 1200", "response-time-aoa-freq-box").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 5)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderResponseTimeAoAFreqBox(d).Render(ctx, templ_7745c5c3_Buffer)
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
