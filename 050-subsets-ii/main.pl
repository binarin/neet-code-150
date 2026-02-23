#!/usr/bin/env perl
use 5.40.0;
use Data::Dumper;

sub subsets_with_dup($nums) {
    my $len = scalar @$nums;
    my @result;

    my %freqs;
    for my $num (@$nums) {
        $freqs{$num}++;
    }

    my $total_combinations = 1;
    my @freq_nums = map { int } keys %freqs;
    my @freqs = values %freqs;
    for my $freq (@freqs) {
        $total_combinations *= $freq + 1;
    }

    for my $num (0..$total_combinations-1) {
        # say "NUM: $num";
        my @subset;
        my $idx = 0;

        while ($num > 0 && $idx < @freqs) {
            my $power = $freqs[$idx] + 1;
            my $occurences = $num % $power;
            # say "At $idx: p $power, o $occurences";
            push @subset, ($freq_nums[$idx]) x $occurences;
            $num /= $power;
            $idx++;
        }
        # say "subset $num: " . Dumper(\@subset);
        push @result, \@subset;
    }
    return \@result;
}

package tests;
use 5.40.0;
use Test::More;
use Data::Dumper;

sub is_subsets($got, $expected) {
    my %got;
    for my $ss (@$got) {
        $got{join ",", sort @$ss}++;
    }

    my %expected;
    for my $ss (@$expected) {
        $expected{join ",", sort @$ss}++;
    }

    is_deeply(\%got, \%expected, "@{[Data::Dumper->new([$expected])->Terse(1)->Indent(0)->Dump]} == @{[Data::Dumper->new($got)->Terse(1)->Indent(0)->Dump]}");
}

# Tests
sub run_tests {
    is_subsets main::subsets_with_dup([1, 2, 2]), [[], [1], [1, 2], [1, 2, 2], [2], [2, 2]];
    is_subsets main::subsets_with_dup([1]), [[], [1]];
    done_testing();
}


package main;

if ($ENV{RUN_TESTS}) {
    tests::run_tests();
} else {
    say Dumper(subsets_with_dup([1, 2, 2]));
}
