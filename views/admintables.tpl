<!DOCTYPE html>
<html>
<head>
    <title>Магазин техники</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" type="text/css" media="screen" href="/static/lib/jqGrid/css/ui.jqgrid.css"/>
    <link rel="stylesheet" type="text/css" media="screen" href="/static/lib/jquery-ui/themes/ui-lightness/jquery-ui.css"/>
    <link rel="stylesheet" type="text/css" media="screen" href="/static/css/mystyles.css"/>
    <script src="/static/lib/jquery/dist/jquery.js" type="text/javascript"></script>
    <script src="/static/lib/jquery/dist/jquery.min.js" type="text/javascript"></script>
    <script src="/static/lib/jqGrid/js/i18n/grid.locale-ru.js" type="text/javascript"></script>
    <script src="/static/lib/jqGrid/js/jquery.jqGrid.min.js" type="text/javascript"></script>
</head>
<body>
    <form method="POST" action='{{urlfor "AdminController.ApplyChanges"}}'>
        <table class="globalTable">
            <tr>
                <div class="ui-widget ui-widget-header header">
                    <div class="headerBoxLeft"><a href="/webshop">Магазин виртуальной техники</a></div>
                    <div class="headerEmptyBlock"></div>
                    <div class="headerBoxCenter"><a href="/webshop/purchases">Товар: {{.PurchaseCount}}шт. Сумма: {{.Sum}}</a></div>
                    {{if .Authorized}}
                    <div class="headerBoxRight"><a href="/webshop/logout">Выйти</a></div>
                    <div class="headerBoxRight">{{.Username}}</div>
                    {{else}}
                    <div class="headerBoxRight"><a href="/webshop/signup">Регистрация</a></div>
                    <div class="headerBoxRight"><a href="/webshop/login">Войти</a></div>
                    {{end}}
                    {{if .IsAdmin}}
                    <div class="headerBoxRight"><a href="/webshop/admin">Admin</a></div>
                    {{else}}
                    {{if .Authorized}}
                    <div class="headerBoxRight"><a href="/webshop/user">Личный кабинет</a></div>
                    {{end}}
                    {{end}}
                </div>
            </tr>
            <tr>
                <td class="catalog">
                    <table>
                        <tbody>
                            <tr>
                                <td class="catalog" id="catalog">
                                </td>
                            </tr>
                            {{range .TableNames}}
                            <tr><td><a href="/webshop/admin/edittable/{{.}}">{{.}}</a></td></tr>
                            {{end}}
                            <tr><td><input type="submit" value="Сохранить"/></td></tr>
                        </tbody>

                    </table>
                </td>

                <td class="grid">
                    {{$isGoodsPage := .IsGoodsPage}}
                    <table class="equipTable admingrid">
                        <thead>
                            <tr>
                                <th>Изменить</th>
                                <th>Удалить</th>
                                {{if $isGoodsPage}}
                                    <th>Страница товара</th>
                                {{end}}
                                {{range .Headers}}
                                    <th>{{.Namei18n}}</th>
                                {{end}}
                            </tr>
                        </thead>
                        <tbody>
                            {{$headers := .Headers}}
                            {{range .Data}}
                            <tr>
                                {{$id := index . 0}}
                                {{$equipId := index . 2}}
                                <td><input type="checkbox" name="checkedit{{$id}}" id="checkedit{{$id}}" /></td>
                                <td><input type="checkbox" name="checkremove{{$id}}" id="checkremove{{$id}}" /></td>
                                {{if $isGoodsPage}}
                                    <td><a href="/webshop/goodsinfo/{{$equipId}}">/webshop/goodsinfo/{{$equipId}}</a></td>
                                {{end}}
                                {{range $idx, $value := .}}
                                    {{$h := index $headers $idx}}
                                    {{$itsId := compare $idx 0}}
                                    {{if $itsId}}
                                        <td>{{.}}</td>
                                    {{else}}
                                        {{if $h.IsTextarea}}
                                        <td><textarea name="{{$h.Name}}{{$id}}" onchange="setCheckBox('{{$id}}')" pattern="{{$h.Pattern}}">{{.}}</textarea></td>
                                        {{else}}
                                        <td><input value="{{.}}" name="{{$h.Name}}{{$id}}" onchange="setCheckBox('{{$id}}')" pattern="{{$h.Pattern}}"></td>
                                        {{end}}
                                    {{end}}
                                {{end}}
                            </tr>
                            {{end}}
                            <tr>
                                <td></td>
                                <td><b>Добавить</b></td>
                                {{range $idx, $value := .EmptyData}}
                                    {{$v := index $headers $idx}}
                                    <td></td>
                                {{end}}
                            </tr>
                            <tr>
                                {{$id := index . 0}}
                                {{$equipId := index . 2}}
                                <td></td>
                                <td><input type="checkbox" name="checkedit0" id="checkedit0" /></td>
                                {{if $isGoodsPage}}
                                    <td></td>
                                {{end}}
                                {{range $idx, $value := .EmptyData}}
                                    {{$h := index $headers $idx}}
                                    {{$itsId := compare $idx 0}}
                                    {{if $itsId}}
                                        <td>{{.}}</td>
                                    {{else}}
                                        {{if $h.IsTextarea}}
                                        <td><textarea name="{{$h.Name}}0" onchange="setCheckBox('0')" pattern="{{$h.Pattern}}">{{.}}</textarea></td>
                                        {{else}}
                                        <td><input value="{{.}}" name="{{$h.Name}}0" onchange="setCheckBox('0')" pattern="{{$h.Pattern}}"></td>
                                        {{end}}
                                    {{end}}
                                {{end}}
                            </tr>
                        </tbody>
                    </table>
                </td>
            </tr>
        </table>
    </form>
    <script type="text/javascript">
        $("#catalog").html("{{.Catalog}}");

        function setCheckBox(id) {
            document.getElementById("checkedit" + id).checked = true
        }
    </script>

</body>
</html>
