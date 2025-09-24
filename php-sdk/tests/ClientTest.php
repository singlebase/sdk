<?php

use PHPUnit\Framework\TestCase;
use Singlebase\Client;
use Singlebase\ResultOK;
use Singlebase\ResultError;

class ClientTest extends TestCase
{
    public function testValidatePayloadValid()
    {
        $client = new Client("abc", null, "test");
        $payload = ["op" => "create", "foo" => 123];
        $this->assertEquals($payload, $this->invokeMethod($client, "validatePayload", [$payload]));
    }

    public function testValidatePayloadInvalid()
    {
        $this->expectException(Exception::class);
        $client = new Client("abc", null, "test");
        $this->invokeMethod($client, "validatePayload", [["foo" => "bar"]]);
    }

    public function testDispatchSuccess()
    {
        $client = $this->getMockBuilder(Client::class)
            ->setConstructorArgs(["abc", null, "test"])
            ->onlyMethods(["dispatch"])
            ->getMock();

        $client->method("dispatch")->willReturn(new ResultOK(["msg" => "ok"]));

        $result = $client->dispatch(["op" => "ping"]);
        $this->assertInstanceOf(ResultOK::class, $result);
        $this->assertEquals("ok", $result->data["msg"]);
    }

    public function testDispatchError()
    {
        $client = $this->getMockBuilder(Client::class)
            ->setConstructorArgs(["abc", null, "test"])
            ->onlyMethods(["dispatch"])
            ->getMock();

        $client->method("dispatch")->willReturn(new ResultError("Bad Request", 400));

        $result = $client->dispatch(["op" => "ping"]);
        $this->assertInstanceOf(ResultError::class, $result);
        $this->assertEquals("Bad Request", $result->error);
    }

    /**
     * Helper to access private/protected methods
     */
    protected function invokeMethod(&$object, $methodName, array $parameters = [])
    {
        $reflection = new ReflectionClass(get_class($object));
        $method = $reflection->getMethod($methodName);
        $method->setAccessible(true);

        return $method->invokeArgs($object, $parameters);
    }
}
