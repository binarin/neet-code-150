#!/usr/bin/env perl
use 5.40.0;
use Data::Dumper;

my %cache;
sub num_decodings($s) {
    if (!exists $cache{$s}) {
        if (length $s == 0) {
            return $cache{$s} = 1;
        }
        my $first = substr $s, 0, 1;
        if (int($first) == 0) {
            return $cache{$s} = 0;
        }
        my $sub_decodings = num_decodings(substr $s, 1);
        if (length $s > 1) {
            my $double = substr $s, 0, 2;
            if (int($double) <= 26) {
                $sub_decodings += num_decodings(substr $s, 2);
            }
        }
        return $cache{$s} = $sub_decodings;
    }

    return $cache{$s};
}

package tests;
use 5.40.0;
use Test::More;

sub run_tests {
    is main::num_decodings("12"), 2;
    is main::num_decodings("226"), 3;
    is main::num_decodings("06"), 0;
    is main::num_decodings("11106"), 2;
    done_testing();
}

package main;

if ($ENV{RUN_TESTS}) {
    tests::run_tests();
} else {
    say num_decodings("12");
}
