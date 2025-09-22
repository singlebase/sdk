<?php

use PHPUnit\Framework\TestCase;
use Singlebase\Result;
use Singlebase\ResultOK;
use Singlebase\ResultError;

class ResultTest extends TestCase
{
    public function testResultToArrayAndToString()
    {
        $r = new Result(["foo" => "bar"], [], true, null, 201);
        $arr = $r->toArray();

        $this->assertEquals("bar", $arr["data"]["foo"]);
        $this->assertStringContainsString("Result", (string)$r);
    }

    public function testResultOkAndError()
    {
        $ok = new ResultOK(["success" => true]);
        $err = new ResultError("Something failed", 400);

        $this->assertTrue($ok->ok);
        $this->assertFalse($err->ok);
        $this->assertEquals(400, $err->statusCode);
    }
}
