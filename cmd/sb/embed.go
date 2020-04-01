package main

const htmlTemplateString = `<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8">
  <!-- Global site tag (gtag.js) - Google Analytics -->
  <script async src="https://www.googletagmanager.com/gtag/js?id=UA-162001246-1"></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag() { dataLayer.push(arguments); }
    gtag('js', new Date());

    gtag('config', 'UA-162001246-1');
  </script>

  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://fonts.googleapis.com/css?family=Roboto+Condensed&display=swap" rel="stylesheet">
  <style>
    .best {
      font-weight: bold;
      color: green;
    }

    .failed {
      text-decoration: line-through;
    }

    ul,
    ul li {
      display: inline-block
    }

    div,
    ul,
    li {
      padding: 0;
      margin: 2px;
    }

    body {
      font-family: 'Roboto Condensed', sans-serif;
    }

    tr {
      background-color: #fff;
    }

    tr:hover {
      background-color: #eee;
    }

    main {
      height: 100%;
      overflow: auto;
      margin-left: auto;
      margin-right: auto;
    }

    th {
      position: -webkit-sticky;
      position: sticky;
      background: inherit;
    }

    thead th {
      top: 0;
    }

    tbody th {
      left: 0;
    }

    table {
      border-collapse: collapse;
      text-align: right;
    }

    .center {
      text-align: center;
    }

    td,
    th {
      padding: 3px;
    }
  </style>
</head>

<body>
  <div>
    Legend:
    <ul>
      <li class="best">best</li>
      <li class="">accepted</li>
      <li class="failed">failed</li>
    </ul>
  </div>
  <main>
    <table>
      <thead>
        <th style="z-index: 1; left: 0">User</th>
        <th class="center">Rank</th>
        <th class="center">Passed</th>
        <th class="center" colspan="2">Time</th>
        {{range $name := .Homework.Cases}}
        <th class="center">{{$name}}</th>
        {{end}}
      </thead>
      <tbody>
        {{range $row := .Rows}}
        <tr>
          <th>{{$row.Submission.User}}</th>
          <td class="center">{{$row.Rank}}</td>
          <td class="center">{{$row.NumPassed}}</td>
          <td>
            {{$row.TotalTime | printf "%.2f"}}
          </td>
          <td>
            {{if gt $row.PenaltyTime 0.0}}(+{{$row.PenaltyTime | printf "%.0f"}}){{end}}
          </td>
          {{range $cell := $row.Cells}}
          <td class="{{$cell.Class}}" title="{{$cell.Title}}">{{$cell.Value}}</td>
          {{end}}
        </tr>
        {{end}}
      </tbody>
    </table>
  </main>
  <div>
    The scoreboard and judge submissions are not used for grading.
    Only code submitted to iLMS count.
  </div>
</body>

</html>
`