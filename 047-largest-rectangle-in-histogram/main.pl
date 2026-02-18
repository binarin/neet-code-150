#!/usr/bin/env perl
use 5.40.0;
use Test::More;
use Data::Dumper;
use List::Util qw/min max/;

sub dump_stack(@stack) {
    join ", ", map "$_->[0]:$_->[1]", @stack;
}

sub largest_rectangle_area {
    my ($heights) = @_;

    my $largest = 0;

    my $improve_guess = sub ($left_val, $left_pos, $right_pos) {
        my $area = ($right_pos-$left_pos)*$left_val;
        my $rect_str = "$left_val:$left_pos→$right_pos";
        if ($area > $largest) {
            $largest = $area;
            say "New largest $area: $rect_str";
        } else {
            say "Discard: $rect_str";
        }
    };

    my @stack;
  outer:
    for my $i (0..$#$heights) {
        say "At $i: " . dump_stack(@stack);
        my $cur = $heights->[$i];
        unless (@stack) {
            push @stack, [$cur, $i];
            next
        }
        my $last_val = $stack[$#stack][0];
        if ($last_val == $cur) {
            next
        } elsif ($cur > $last_val) {
            push @stack, [$cur, $i];
        } else {
            my $best_fallback = $cur;
            while (@stack) {
                my ($last_val, $last_pos) = (pop @stack)->@*;
                if ($last_val > $cur) {
                    $improve_guess->($last_val, $last_pos, $i);
                    $best_fallback = min($best_fallback, $last_pos);
                } elsif ($last_val == $cur) {
                    push @stack, [$last_val, $last_pos];
                    next outer
                } else {
                    push @stack, [$last_val, $last_pos];
                    push @stack, [$cur, $best_fallback];
                    next outer
                }
            }
            push @stack, [$cur, $best_fallback];
        }
    }

    my $max_pos = scalar(@$heights);
    while (@stack) {
        my ($last_val, $last_pos) = (pop @stack)->@*;
        $improve_guess->($last_val, $last_pos, $max_pos);
    }

    return $largest;
}

# Tests
is(largest_rectangle_area([2, 1, 5, 6, 2, 3]), 10, 'Example 1');
is(largest_rectangle_area([2, 4]), 4, 'Example 2');

done_testing();
