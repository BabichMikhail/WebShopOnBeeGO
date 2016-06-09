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
            <td class="catalog" id="catalog">
            </td>
            <form class="ui-form ui-widget" method="POST" action="/webshop/admin/editcard/{{.TableName}}/{{.EquipId}}">
            <td class="goods">
                <table class="goodsTable">
                    <tbody>
                        <tr>
                            <td class="propImg"><img src="{{.Image}}"></img></td>
                            {{if .IsCount}}
                            <td width="500px"></td>
                            <td class="equipPlus goodInfoImg"><a href="{{urlfor "PurchaseController.Get"}}/change/{{.EquipId}}/1">
                                <img src="/static/img/grid/plus.jpg"></img>
                            </a></td>
                            <td class="equipMinus goodInfoImg"><a href="{{urlfor "PurchaseController.Get"}}/change/{{.EquipId}}/0">
                                <img src="/static/img/grid/minus.jpg"></img>
                            </a></td>
                            <td></td>
                            {{else}}
                            <td colspan="3"></td>
                            <td class="equipBuy"><a href="/webshop/card/add/{{.EquipId}}">Купить</a></td>
                            {{end}}
                        </tr>
                        {{$isAdmin := .IsAdmin}}
                        <tr>
                            <td class="characteristics descrTitle" colspan="5">Описание</td>
                        </tr>
                        <tr>
                            <td class="descr" colspan="5">
                                {{if $isAdmin}}
                                <textarea rows="10" cols="100" name="description">{{.Description}}</textarea>
                                {{else}}
                                {{.Description}}
                                {{end}}
                            </td>
                        </tr>
                        <tr><td class="afterLastProp" colspan="5"></td></tr>
                        <tr>
                            <td class="characteristics descrTitle" colspan="5">Характеристики</td>
                        </tr>
                        {{range .Characteristics}}
                        <tr>
                            <td class="propName">{{.Key}}</td>
                            {{if $isAdmin}}
                            <td class="propValue"><input name="{{.PropName}}" value="{{.Value}}"></td>
                            {{else}}
                            <td class="propValue">{{.Value}}</td>
                            {{end}}
                            <td colspan="3"></td>
                        </tr>
                        {{end}}
                        <tr><td class="afterLastProp" colspan="5"></td></tr>
                        <tr>
                            <td>
                                {{if $isAdmin}}
                                    <input type="submit" value="Сохранить"/>
                                {{end}}
                            </td>
                        </tr>
                    </tbody>
                </table>
            </td>
            </form>
        </tr>
    </table>
    <script type="text/javascript">
        $("#catalog").html("{{.Catalog}}");
    </script>

</body>
</html>
