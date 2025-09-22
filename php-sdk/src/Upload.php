<?php

namespace Singlebase;

use Exception;

/**
 * Upload a file to a presigned URL.
 */
class Upload
{
    /**
     * Upload a file using presigned URL data.
     *
     * @param string $filepath Path to the file
     * @param array $data Presigned URL data: [ 'url' => string, 'fields' => array ]
     * @return bool True if upload succeeded
     * @throws Exception If upload fails
     */
    public static function uploadPresignedFile(string $filepath, array $data): bool
    {
        if (!file_exists($filepath)) {
            throw new Exception("File not found: $filepath");
        }

        $fields = $data["fields"] ?? [];
        $url = $data["url"] ?? null;
        if (!$url) {
            throw new Exception("Missing upload URL");
        }

        $postData = [];
        foreach ($fields as $key => $value) {
            $postData[$key] = $value;
        }
        $postData["file"] = new \CURLFile($filepath);

        $ch = curl_init($url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_POST, true);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $postData);

        $response = curl_exec($ch);
        $status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
        if ($response === false) {
            throw new Exception(curl_error($ch));
        }
        curl_close($ch);

        if ($status >= 200 && $status < 300) {
            return true;
        }
        throw new Exception("Upload failed: HTTP $status");
    }
}
