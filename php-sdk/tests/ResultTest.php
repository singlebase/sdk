<?php

use PHPUnit\Framework\TestCase;
use Singlebase\ResultOK;
use Singlebase\ResultError;

class ResultTest extends TestCase {
    public function testGetDataSuccessAndDefault() {
        $r = new ResultOK([
            "data" => [
                "address" => [
                    "city" => [
                        "city_fullname" => "San Francisco",
                        "zipcode" => 94107
                    ]
                ]
            ]
        ]);

        // Full data
        $this->assertIsArray($r->getData());

        // Dot notation
        $this->assertEquals("San Francisco", $r->getData("address.city.city_fullname"));
        $this->assertEquals(94107, $r->getData("address.city.zipcode"));

        // Missing key returns default
        $this->assertEquals("USA", $r->getData("address.country", "USA"));
    }

    public function testGetDataThrowsTypeError() {
        $this->expectException(\TypeError::class);

        $r = new ResultOK([
            "data" => [ "user" => [ "id" => 123 ] ]
        ]);

        $r->getData("user.id.value");
    }
}
