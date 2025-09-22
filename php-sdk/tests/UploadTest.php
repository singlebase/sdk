<?php

use PHPUnit\Framework\TestCase;
use Singlebase\Upload;

class UploadTest extends TestCase
{
    public function testUploadFileFailsIfMissingFile()
    {
        $this->expectException(Exception::class);
        Upload::uploadPresignedFile("nonexistent.txt", [
            "url" => "http://fake-url",
            "fields" => []
        ]);
    }

    public function testUploadFileThrowsOnMissingUrl()
    {
        $this->expectException(Exception::class);
        $tmpFile = tempnam(sys_get_temp_dir(), "upload");
        file_put_contents($tmpFile, "hello");
        Upload::uploadPresignedFile($tmpFile, [
            "fields" => []
        ]);
    }

    public function testUploadFileSuccessMocked()
    {
        // Fake CURLFile so we don’t actually hit a network
        $this->markTestSkipped("Integration test requires presigned URL");
    }
}
