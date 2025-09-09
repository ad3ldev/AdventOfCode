<?php
$count = 0;
$inputFile = fopen("input.txt", "r") or die("Unable to open file!");
$fileSize = fileSize("input.txt");
$content = fread($inputFile, $fileSize);
foreach (str_split($content) as $char) {
    if ($char == '(') {
        $count += 1;
    } elseif ($char == ')') {
        $count -= 1;
    }
}
fclose($inputFile);
echo $count;
