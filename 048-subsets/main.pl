#!/usr/bin/env perl
use 5.40.0;
use Test::More;
use Data::Dumper;

sub subsets {
    my ($nums) = @_;
    my @result;
    my $total = 2 ** @$nums;
    for my $num (0..$total-1) {
        my $idx = 0;
        my @subset;
        while ($num > 0) {
            if ($num % 2) {
                push @subset, $nums->[$idx];
            }
            $idx++;
            $num /= 2;
        }
        push @result, \@subset;
    }
    return \@result;
}


is_deeply(subsets([1,2,3]), [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]], 'Example 1');
is_deeply(subsets([0]), [[],[0]], 'Example 2');

done_testing();
