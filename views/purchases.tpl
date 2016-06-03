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
                <div class="headerBoxCenter"><a href="/webshop/purchases">Товар: <span id="purchaseCount">{{.PurchaseCount}}</span>шт. Сумма: <span id="purchaseSum">{{.Sum}}<span></a></div>
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
            </td>

                <td class="grid">
                    <form method="POST" action='{{urlfor "PurchaseController.Get"}}'>
                    <table>
                        <thead>
                            <tr>
                                <th>Код</th>
                                <th>Фото</th>
                                <th>Наименование товара</th>
                                <th>Производитель</th>
                                <th>Цена</th>
                                <th>Количество</th>
                            </tr>
                        </thead>
                        <tbody>
                            {{range .Equipment}}
                            <tr>
                                <td>{{.Equip_id}}</td>
                                <td><img src="{{.Small_image}}"></img></td>
                                <td><a href="/webshop/goodsinfo/{{.Equip_id}}">{{.Name}}</a></td>
                                <td>{{.Nation}}</td>
                                <td>{{.Price}}</td>
                                <td><input name="count{{.Equip_id}}" id="count{{.Equip_id}}" value="{{.Count}}" onChange="setSum()"/></td>
                            </tr>
                            {{end}}
                            <tr>
                                <td colspan="5"></td>
                                <td><input type="submit" value='"Купить"'/></td>
                        </tbody>
                    </table>
                    </form>
                </td>

        </tr>
    </table>
    <script type="text/javascript">
        $("#catalog").html("{{.Catalog}}");

        function addValueIfInt(value) {
            if (value == parseInt(value, 10)) {
                return value
            } else {
                alert(value + ' не число')
                return 0
            }
        }

        function setSum() {
            var sum = 0
            var count = 0
            var value = ''
            {{range .Equipment}}
            value = document.getElementById('count{{.Equip_id}}').value
            if (value == parseInt(value, 10)) {
                count += parseInt(value, 10)
                sum += {{.Price}}*count
            } else {
                alert(count + ' не число')
            }
            {{end}}
            document.getElementById("purchaseCount").textContent = count;
            document.getElementById("purchaseSum").textContent = sum;
        }
    </script>

</body>
</html>
