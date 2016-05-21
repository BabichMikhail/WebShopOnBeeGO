# script for insert data into database
# todo download images
use strict;
use warnings;

use DBD::SQLite;
use HTTP::Tiny;
use JSON;

use Data::Dumper;

my $dbh = DBI->connect("dbi:SQLite:orm_test.db","","");

sub Tanks {
    my $response =  HTTP::Tiny->new->get('https://api.worldoftanks.ru/wot/encyclopedia/vehicles/?application_id=demo');
    my $tanks = decode_json($response->{content});
    for (values %{$tanks->{data}}) {
        defined $_->{price_gold} or defined $_->{price_credit} or next;
        $_->{price_gold} = 0 if !defined $_->{price_gold};
        $_->{price_credit} = 0 if !defined $_->{price_credit};
        $_->{price} = 400*$_->{price_gold}  + $_->{price_credit};
        $_->{small_icon} =$_->{images}{small_icon};
        $_->{contour_icon} =$_->{images}{contour_icon};
        $_->{big_icon} =$_->{images}{big_icon};
        my $is_premium = $_->{is_premium} eq "false" ? "0" : "1";
        my $is_gift = $_->{is_gift} eq "false" ? "0" : "1";
        $dbh->do( qq ~ INSERT INTO
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
        $_->{small_image} =$_->{images}{small};
        $_->{medium_image} =$_->{images}{medium};
        $_->{large_image} =$_->{images}{large};
        my $is_premium = $_->{is_premium} eq "false" ? "0" : "1";
        my $is_gift = $_->{is_gift} eq "false" ? "0" : "1";
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

Tanks;
Warplanes;
ClosureTable;
Types;
Nations;
Levels;
