<html>
  <head>
      <title>all items</title>
  </head>
  <body>

    <h1>List of todo items:</h1>

    <div>
        <ul>
        {{ range $item := .}}
            <li><input type="checkbox" id="{{ $item.Id }}" value="{{ $item.IsFinished }}" /><label for="{{ $item.Id }}">{{ $item.Caption }} </label></li>
        {{ end }}
        </ul>
    </div>

  </body>
</html>
