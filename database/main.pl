# script for insert data into database
use strict;
use warnings;

use DBD::SQLite;
use HTTP::Tiny;
use JSON;
use Cwd;
use File::Spec;

use Data::Dumper;

my $wd = Cwd::cwd();
my $big_image_dir = File::Spec->catfile('static', 'img', 'big');
my $small_image_dir = File::Spec->catfile('static', 'img', 'small');
chdir(File::Spec->catfile($wd, 'database'));
my $dbh = DBI->connect("dbi:SQLite:dbSqlite3Shop.db","","");
chdir($wd);

sub LoadImage {
    my ($img_dir, $url) = @_;
    my ($img_name) = $url =~ /([^\/]*$)/;
    #print $img_name . "\n";
    my $fname = File::Spec->catfile($img_dir, $img_name);
    if (-e $fname) {
        $fname =~ s/\\/\//g;
        return "/$fname";
    }
    my $response =  HTTP::Tiny->new->get($url);
    open my $fh, ">", $fname or die "Can't load image $url and save to $fname";
    binmode $fh;
    print $fh $response->{content};
    $fname =~ s/\\/\//g;
    "/$fname";
}

sub Tanks {
    my $response =  HTTP::Tiny->new->get('https://api.worldoftanks.ru/wot/encyclopedia/vehicles/?application_id=demo');
    my $tanks = decode_json($response->{content});
    for (values %{$tanks->{data}}) {
        defined $_->{price_gold} or defined $_->{price_credit} or next;
        $_->{price_gold} = 0 if !defined $_->{price_gold};
        $_->{price_credit} = 0 if !defined $_->{price_credit};
        $_->{price} = 400*$_->{price_gold} + $_->{price_credit};
        $_->{big_icon} = LoadImage($big_image_dir, $_->{images}{big_icon});
        $_->{small_icon} = LoadImage($small_image_dir, $_->{images}{small_icon});
        #$_->{small_icon} = $_->{images}{small_icon};
        #$_->{big_icon} = $_->{images}{big_icon};
        my $is_premium = $_->{is_premium} eq "false" ? "0" : "1";
        my $is_gift = $_->{is_gift} eq "false" ? "0" : "1";
        $_->{weight} = $_->{default_profile}{weight};
        $_->{max_weight} = $_->{default_profile}{max_weight};
        $_->{armor} = $_->{default_profile}{armor}{hull}{front} .
            "/" . $_->{default_profile}{armor}{hull}{sides} .
            "/" . $_->{default_profile}{armor}{hull}{rear};
        $_->{hp} = $_->{default_profile}{hp};
        $_->{speed_forward} = $_->{default_profile}{speed_forward};
        $_->{speed_backward} = $_->{default_profile}{speed_backward};
        $dbh->do(qq ~INSERT INTO
            equipments (
                description, equip_id, equip_type, image,
                is_gift, is_premium, level, name, nation,
                price, short_name, small_image, type
            )
            VALUES (
                "$_->{description}", $_->{tank_id}, "tanks", "$_->{big_icon}",
                $is_gift, $is_premium, $_->{tier}, "$_->{name}", "$_->{nation}",
                $_->{price}, "$_->{short_name}", "$_->{small_icon}", "$_->{type}"
            )~);
        $dbh->do(qq ~INSERT INTO
            tanks (
                description, equip_id,
                is_premium, level, name, nation, price,
                type, weight, max_weight, armor,
                hp, speed_forward, speed_backward
            )
            VALUES (
                "$_->{description}", $_->{tank_id},
                $is_premium, $_->{tier}, "$_->{name}", "$_->{nation}", $_->{price},
                "$_->{type}", $_->{weight}, $_->{max_weight}, "$_->{armor}",
                $_->{hp}, $_->{speed_forward}, $_->{speed_backward}
            )~)
    }
}

sub Warplanes {
    my $response = HTTP::Tiny->new->get('https://api.worldofwarplanes.ru/wowp/encyclopedia/planes/?application_id=demo&fields=plane_id');
    my $query = join '%2C', keys %{decode_json($response->{content})->{data}};
    $query = 'https://api.worldofwarplanes.ru/wowp/encyclopedia/planeinfo/?application_id=demo&plane_id=' . $query;
    $response = HTTP::Tiny->new->get($query);
    my $warplanes = decode_json($response->{content});
    for (values %{$warplanes->{data}}) {
        defined $_->{price_gold} or defined $_->{price_credit} or next;
        $_->{price_gold} = 0 if !defined $_->{price_gold};
        $_->{price_credit} = 0 if !defined $_->{price_credit};
        $_->{price} = 400*$_->{price_gold}  + $_->{price_credit};
        $_->{large_image} = LoadImage($big_image_dir, $_->{images}{large});
        $_->{small_image} = LoadImage($small_image_dir, $_->{images}{small});
        my $is_premium = $_->{is_premium} eq "false" ? "0" : "1";
        my $is_gift = $_->{is_gift} eq "false" ? "0" : "1";
        my $f = $_->{features};
        #for (keys %{$f}) {
        #    print $_ . "\n";
        #}
        $dbh->do(qq ~INSERT INTO
            equipments (
                description, equip_id, equip_type, image,
                is_gift, is_premium, level, name, nation,
                price, short_name, small_image, type
            )
            VALUES (
                "$_->{description}", $_->{plane_id}, "warplanes", "$_->{large_image}",
                $is_gift, $is_premium, $_->{level}, "$_->{name_i18n}", "$_->{nation}",
                $_->{price}, "$_->{short_name_i18n}", "$_->{small_image}", "$_->{type}"
            )~);
        $dbh->do(qq ~INSERT INTO
            warplanes (
                description, equip_id, is_premium, level,
                name, nation, price, type,
                weight, hp, speed_ground, maneuverability,
                max_speed, stall_speed, optimal_height,
                roll_maneuver, dive_speed, opt_maneuver_speed
            )
            VALUES (
                "$_->{description}", $_->{plane_id}, $is_premium, $_->{level},
                "$_->{name_i18n}", "$_->{nation}", $_->{price}, "$_->{type}",
                $f->{mass}, $f->{hp}, $f->{speed_at_the_ground}, $f->{maneuverability},
                $f->{max_speed}, $f->{stall_speed}, $f->{optimal_height},
                $f->{roll_maneuverability}, $f->{dive_speed}, $f->{optimal_maneuver_speed}
            )~);
    }
}

sub ClosureTable {
    my @catalogstreepath = (
        '2, "tanks", "Танки"',
        '3, "light_tanks", "Лёгкие танки"',
        '4, "medium_tanks", "Средние танки"',
        '5, "heavy_tanks", "Тяжёлые танки"',
        '6, "warplanes", "Самолёты"',
        '7, "fighter", "Истребители"',
        '8, "multirole_fighter", "Многоцелевые истребители"',
        '10, "heavy_fighter", "Тяжёлые истребители"',
        '11, "aircraft", "Штурмовики"',
        '12, "spg", "САУ"',
        '13, "at_spg", "ПТ-САУ"',
        '14, "low_levels", "Низкого уровня"',
        '15, "medium_levels", "Среднего уровня"',
        '16, "high_levels", "Высокого уровня"',
    );
    my @catalogs = (
        '2, 0, 3',
        '2, 0, 4',
        '2, 0, 5',
        '2, 0, 12',
        '2, 0, 13',
        '6, 0, 7',
        '6, 0, 8',
        '6, 0, 10',
        '6, 0, 11',
    );
    my @s = ({
        ancestor => 2,
        cid => [3, 4, 5, 12, 13],
    }, {
        ancestor => 6,
        cid => [7, 8, 10, 11],
    });

    for my $s (@s) {
        for my $cid (@{$s->{cid}}) {
            for my $descendant (qw[14 15 16]) {
                push @catalogs, "$cid, $s->{ancestor}, $descendant";
                push @catalogs, "$descendant, $cid, 0";
            }
        }
    }

    for (@catalogstreepath) {
        $dbh->do(qq ~INSERT INTO
            catalogstreepath (
                ctpid, name, name_i18n
            )
            VALUES (
                $_
            )~);
    }
    for (@catalogs) {
        $dbh->do(qq ~INSERT INTO
            catalogs (
                cid, ancestor, descendant
            )
            VALUES (
                $_
            )~);
    }
}

sub Levels {
    my @levels = (
        '"low_levels", 1',
        '"low_levels", 2',
        '"low_levels", 3',
        '"low_levels", 4',
        '"medium_levels", 5',
        '"medium_levels", 6',
        '"medium_levels", 7',
        '"high_levels", 8',
        '"high_levels", 9',
        '"high_levels", 10',
    );
    for (@levels) {
        $dbh->do(qq ~INSERT INTO
            levels (
                value, level
            )
            VALUES (
                $_
            )~);
    }
}

sub Types {
    my @types = (
        '"fighter", "fighter"',
        '"multirole_fighter", "navy"',
        '"heavy_fighter", "fighterHeavy"',
        '"aircraft", "assault"',
        '"light_tanks", "lightTank"',
        '"medium_tanks", "mediumTank"',
        '"heavy_tanks", "heavyTank"',
        '"at_spg", "AT-SPG"',
        '"spg", "SPG"',
    );
    for (@types) {
        $dbh->do(qq ~INSERT INTO
            types (
                name_catalog, name
            )
            VALUES (
                $_
            )~);
    }
}

sub Nations {
    my @nations = (
        '"uk", "Великобритания"',
        '"ussr", "СССР"',
        '"germany", "Германия"',
        '"france", "Франция"',
        '"chine", "Китай"',
        '"japan", "Япония"',
        '"usa", "США"',
        '"czech", "Чехословакия"',
    );
    for (@nations) {
        $dbh->do(qq ~INSERT INTO
            nations (
                name, name_i18n
            )
            VALUES (
                $_
            )~);
    }
}

sub Admin {
    $dbh->do(qq ~INSERT INTO
        users (
            login, password, rights
        )
        VALUES (
            "admin", "d033e22ae348aeb5660fc2140aec35850c4da997", 0
        )~);
}

Tanks;
Warplanes;
ClosureTable;
Types;
Nations;
Levels;
Admin;
