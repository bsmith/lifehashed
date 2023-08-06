#!perl

use 5.036;
use strict;
use warnings;

sub xy_to_index {
	my ($x, $y) = @_;
	my $index = 0;
	my $cur = 1;
	while ($x > 0 or $y > 0) {
		$index |= ($x & 1) * $cur;
		$index |= ($y & 1) * ($cur << 1);
		$cur <<= 2;
		$x >>= 1;
		$y >>= 1;
	}
	return $index;
}

sub index_to_xy {
	my ($index) = @_;
	my $cur = 1;
	my ($x, $y) = (0, 0);
	while ($index > 0) {
		$x |= ($index & 1) * $cur;
		$y |= (($index & 2) >> 1) * $cur;
		$index >>= 2;
		$cur <<= 1;
	}
	return ($x, $y);
}

my $size = 8;

my @array_1d = ("???") x ($size * $size);
my @array_2d = map [(-1) x $size], 0..$size-1;

for my $column (0..$size-1) {
	printf "%06b ", xy_to_index($column, 0);
}
print "\n";

for my $row (0..$size-1) {
	for my $column (0..$size-1) {
		my $index = xy_to_index($column, $row);
		printf "%2d ", $index;
		$array_1d[$index] = "$row-$column";
		$array_2d[$row][$column] = $index;
	}
	print "\n";
}

for my $index (0..$size*$size-1) {
	printf "(%d,%d) ", index_to_xy($index);
}
print "\n";

