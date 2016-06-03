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
                {{end}}
            </div>
        </tr>
        <tr>
            <td class="catalog" id="catalog">
            </td>
            <td class="grid">
                <table id="grid"></table>
                <div id="jqGridPager"></div>
            </td>
        </tr>
    </table>
    <script type="text/javascript">
        $(document).ready(function () {
            $('#grid').jqGrid({
                datatype: 'local',
                data: [],
                height: 650,
                width: 700,
                colModel: [
                    { name: 'id', width: 50, label: 'ID', align: 'center', key: true },
                    { name: 'sum', width: 150, label: 'Сумма', align: 'center', sorttype: 'int' },
                    { name: 'count', width: 150, label: 'Количество товаров', align: 'center', sorttype: 'int' },
                    { name: 'date', width: 400, label: 'Дата', align: 'center' }
                ],
                pager: '#jqGridPager',
                viewrecords: true,
                rowNum: 50,
                rowList: [50, 100, 200],
                rownumbers: true,
                caption: "Покупки",
            });
            $('#grid').jqGrid('setGridParam', {
                datatype: 'local',
                data: [
                    {{range .Purchases}}
                    { id: {{.Id}}, sum: {{.Sum}}, count: {{.Count}}, date: {{.Date}} },
                    {{end}}
                ],
            }).trigger('reloadGrid');
        });
    </script>
    <script type="text/javascript">
        $("#catalog").html("{{.Catalog}}");
    </script>

</body>
</html>
