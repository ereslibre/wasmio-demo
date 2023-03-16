<?php

header("Content-type: image/png");

$string = "Click me!";
$button = imagecreatefrompng("/button.png");
$orange = imagecolorallocate($button, 220, 210, 60);
$px     = (imagesx($button) - 7.5 * strlen($string)) / 2;
imagestring($button, 3, $px, 9, $string, $orange);
imagepng($button);
imagedestroy($button);

?>