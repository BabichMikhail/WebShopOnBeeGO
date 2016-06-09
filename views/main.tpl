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
            <td class="equipments" id="equipments">
            </td>
        </tr>
    </table>
    <script type="text/javascript">
        $("#catalog").html("{{.Catalog}}");

        function setEquipments() {
            var width = 4
            var equipments = [
                {{range .Equipment}}
                { id: {{.Equip_id}}, name: {{.Name}}, price: {{.Price}}, image: {{.Small_image}}, isCount: {{.IsCount}}},
                {{end}}
            ]
            var html = "<table><tbody>"
            for (var i = 0; i < equipments.length;) {
                html += '<tr>'
                for (var j = 0; j < width && i < equipments.length; ++i, ++j) {
                    var href = '/webshop/goodsinfo/' + equipments[i].id
                    html += '<td><table class="equipment"><tr><td><a href="' + href + '">' +
                        '<img src="' + equipments[i].image + '"></img></a></td></tr>' +
                        '<tr><td><a href="' + href + '">' + equipments[i].name + '</a></td></tr>' +
                        '<tr><td>' + equipments[i].price + '</td></tr>'
                    if (equipments[i].isCount) {
                        html += '<tr><td><a href="/webshop/purchases/">В корзину</a></td></tr></td></table>'
                    } else {
                        html += '<tr><td><a href="/webshop/card/add/' + equipments[i].id + '">Купить</a></td></tr></td></table>'
                    }
                }
                html += '</tr>'
            }
            html += "</tbody></table>"
            return html
        }

        $("#equipments").html(setEquipments());

    </script>
</body>
</html>
