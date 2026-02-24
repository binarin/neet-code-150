#!/usr/bin/env perl
use 5.40.0;
use Data::Dumper;

package DetectSquares;

sub new($class) {
    my $self = {
    };
    bless $self, $class;
    return $self;
}

sub add($self, $point) {
    die "not implemented";
}

sub count($self, $point) {
    die "not implemented";
}

package tests;
use 5.40.0;
use Test::More;

sub run_tests {
    # Example 1:
    # Input: ["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
    #        [[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
    # Output: [null, null, null, null, 1, 0, null, 2]
    my $ds = DetectSquares->new();
    $ds->add([3, 10]);
    $ds->add([11, 2]);
    $ds->add([3, 2]);
    is($ds->count([11, 10]), 1, "count [11,10] after 3 adds");
    is($ds->count([14, 8]), 0, "count [14,8] returns 0");
    $ds->add([11, 2]); # Adding duplicate points is allowed
    is($ds->count([11, 10]), 2, "count [11,10] after duplicate add");
    done_testing();
}

package main;

if ($ENV{RUN_TESTS}) {
    tests::run_tests();
} else {
    # Example 1
    my $ds = DetectSquares->new();
    $ds->add([3, 10]);
    $ds->add([11, 2]);
    $ds->add([3, 2]);
    say $ds->count([11, 10]); # expected: 1
}
