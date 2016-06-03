<!DOCTYPE html>
<html>
<head>
    <title>Магазин техники</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    {{range .Links}}
    <link rel="stylesheet" type="text/css" media="screen" href="{{.}}"/>
    {{end}}
    {{range .Scripts}}
    <script src="{{.}}" type="text/javascript"></script>
    {{end}}
</head>
<body>
    <table class="globalTable">
        <tr>
            <div class="ui-widget ui-widget-header header">
                <div class="headerBoxLeft"><a href="/webshop">Магазин виртуальной техники</a></div>
                {{if .Authorized}}
                <div class="headerBoxRight"><a href="/webshop/logout">Выйти</a></div>
                <div class="headerBoxRight">{{.Username}}</div>
                {{else}}
                <div class="headerBoxRight"><a href="/webshop/signup">Регистрация</a></div>
                <div class="headerBoxRight"><a href="/webshop/login">Войти</a></div>
                {{end}}
                {{if .IsAdmin}}
                <div class="headerBoxRight"><a href="/webshop/admin">Admin</a></div>
                {{end}}
            </div>
        </tr>
        <tr>
            <td class="catalog" id="catalog">
                {{.Catalogue}}
            </td>
            <td>
                {{.Auth}}
            </td>
            <td>
                {{.CenterPage}}
            </td>
        </tr>
    </table>
    <script type="text/javascript">
        $("#catalog").html("{{.Catalog}}");
    </script>
</body>
</html>
