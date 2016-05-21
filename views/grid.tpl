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
                <div class="headerBoxCenter"><a href="/webshop">Магазин виртуальной техники</a></div>
                {{if .Authorized}}
                <div class="headerBoxRight"><a href="/webshop/logout">Выйти</a></div>
                <div class="headerBoxRight">{{.Username}}</div>
                {{else}}
                <div class="headerBoxRight"><a href="/webshop/signup">Регистрация</a></div>
                <div class="headerBoxRight"><a href="/webshop/login">Войти</a></div>
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
                    { name: 'id', width: 120, label: {{.Fields.Id}}, align: 'center', key: true },
                    { name: 'name', width: 250, label: {{.Fields.Name}}, align: 'center' },
                    { name: 'price', width: 150, label: {{.Fields.Price}}, align: 'center', sorttype: 'int' },
                    { name: 'nation', width: 130, label: {{.Fields.Nation}}, align: 'center' }
                ],
                pager: '#jqGridPager',
                viewrecords: true,
                rowNum: 50,
                rowList: [50, 100, 200],
                rownumbers: true,
                caption: "Техника",
            });
            $('#grid').jqGrid('setGridParam', {
                datatype: 'local',
                data: [
                    {{range .Equipment}}
                    { id: {{.Equip_id}}, name: {{.Name}}, price: {{.Price}}, nation: {{.Nation}} },
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
