<?php

define('TWO_LETTERS', 2);
define('THREE_LETTERS', 3);

$boxes = explode(PHP_EOL, file_get_contents('../input.txt'));
$two   = 0;
$three = 0;

foreach ($boxes as $index => $id) {
    $counts = array_count_values(str_split($id));
    $groupedCounts = array_unique(array_values($counts));

    if (in_array(TWO_LETTERS, $groupedCounts)) { $two++; }
    if (in_array(THREE_LETTERS, $groupedCounts)) { $three++; }
}

echo $two * $three;