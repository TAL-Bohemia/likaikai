<?php

class Demo
{

    public function test()
    {
        return SingleObject::getInstance();
    }

}

var_dump((new Demo)->test());