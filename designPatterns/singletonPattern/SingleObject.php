<?php

class SingleObject
{

    private static $singleObject = null;

    private function __construct()
    {
        return self::getInstance();
    }

    public static function getInstance()
    {
        return self::$singleObject ?? new self();
    }

}