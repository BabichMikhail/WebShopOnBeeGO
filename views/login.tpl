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
            <td class="auth" >
                <div class="ui-dialog-content ui-widget-content form-div logForm">
                    <form class="ui-form ui-widget" method="POST" action='{{urlfor "LoginController.Login"}}'>
                        <div class="ui-dialog-titlebar ui-widget-header ui-corner-all ui-helper-clearfix ui-draggable-handle">
                            Авторизация
                        </div>
                        <div>
                            <label for="Login" class="form-label">Имя пользователя</label>
                            <input class="ui-corner-all ui-state-default" name="Login" type="login" value="username"
                                required pattern="[a-zA-Z0-9_.]+" />
                        </div>
                        <div>
                            <label for="Login" class="form-label">Пароль</label>
                            <input class="ui-corner-all ui-state-default" name="Password" type="password"
                                value="" required pattern="[a-zA-Z0-9]{3,}" title="[a-zA-Z0-9]{3,}" />
                        </div>
                        <div>
                            {{if .Error}}
                            <div class="ui-state-error ui-state-error-text">{{.Err_msg}}</div>
                            {{end}}
                            <input class="" type="submit" value="Войти"/>
                        </div>
                    </form>
                </div>
            </td>
        </tr>
    </table>
    <script type="text/javascript">
        $("#catalog").html("{{.Catalog}}");
    </script>
</body>
</html>
