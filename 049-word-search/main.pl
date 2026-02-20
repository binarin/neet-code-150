#!/usr/bin/env perl
use 5.40.0;
use strict;
use warnings;
use Data::Dumper;
use Test::Simple tests => 3;

sub exist {
    my ($board, $word) = @_;
    my $height = @$board;
    my $width = @{$board->[0]};

    my $char_at = sub ($x, $y, $used_cells) {
        if ($x < 0 || $x >= $width || $y < 0 || $y >= $height) {
            return '!';
        }
        if (exists $used_cells->{$y * $width + $x}) {
            return '!';
        }
        return $board->[$y][$x];
    };

    my $dump_state = sub ($cur_x, $cur_y, $used_cells) {
        for my $y (0..$height-1) {
            for my $x (0..$width-1) {
                if ($cur_x == $x && $cur_y == $y) {
                    print "·";
                } elsif ($used_cells->{$y*$width + $x}) {
                    print "!";
                } else {
                    print $board->[$y][$x];
                }
            }
            say "";
        }
        say "";
    };

    my $recur;
    $recur = sub ($x, $y, $subword, $used_cells) {
        # $dump_state->($x, $y, $used_cells);
        if (length $subword == 0) {
            return true
        }

        my $needed_char = substr($subword, 0, 1);
        my $rest_of_word = substr($subword, 1);

        for my $delta ([-1, 0], [1, 0], [0, -1], [0, 1]) {
            my ($dx, $dy) = @$delta;
            my $new_x = $x + $dx;
            my $new_y = $y + $dy;

            my $available_char = $char_at->($new_x, $new_y, $used_cells);
            if ($available_char ne $needed_char) {
                next
            }
            my %new_used = %$used_cells;
            $new_used{$new_y * $width + $new_x} = true;
            if ($recur->($new_x, $new_y, $rest_of_word, \%new_used)) {
                return true;
            }
        }
        return false;
    };

    my $first_char = substr($word, 0, 1);
    my $rest_of_word = substr($word, 1);
    for my $y (0..$height-1) {
        for my $x (0..$width-1) {
            if ($char_at->($x, $y, {}) eq $first_char) {
                if ($recur->($x, $y, $rest_of_word, { ($y*$width + $x) => true })) {
                    return true
                }
            }
        }
    }
    return false;
}

my @board = (
    ['A', 'B', 'C', 'E'],
    ['S', 'F', 'C', 'S'],
    ['A', 'D', 'E', 'E'],
);

ok exist(\@board, "ABCCED");
ok exist(\@board, "SEE");
ok !exist(\@board, "ABCB");
