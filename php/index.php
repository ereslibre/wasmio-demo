<?php

header("Content-type: image/png");

$string = "Click me!";
$button = imagecreatefrompng("/templates/button.png");
$black = imagecolorallocate($button, 0, 0, 0);
$px     = (imagesx($button) - 7.5 * strlen($string)) / 2;
imagestring($button, 5, $px, 9, $string, $black);
imagepng($button);
imagedestroy($button);

?>