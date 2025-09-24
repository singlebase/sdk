<?php

namespace Singlebase;

use Exception;

/**
 * Singlebase API Client
 */
class Client
{
    private const BASE_API_URL = "https://cloud.singlebaseapis.com/api";

    private string $apiKey;
    private string $apiUrl;
    private array $headers;

    /**
     * Construct the API client.
     *
     * @param string $apiKey Your Singlebase API key
     * @param string|null $apiUrl Full API URL (optional)
     * @param string|null $endpointKey Endpoint key appended to base URL if apiUrl not given
     * @param array $headers Additional headers
     *
     * @throws Exception If required parameters are missing
     */
    public function __construct(string $apiKey, ?string $apiUrl = null, ?string $endpointKey = null, array $headers = [])
    {
        if (!$apiKey) {
            throw new Exception("MISSING_API_KEY");
        }
        if (!$apiUrl && !$endpointKey) {
            throw new Exception("MISSING_ENDPOINT_KEY");
        }

        $this->apiKey = $apiKey;
        $this->headers = $headers;
        $this->apiUrl = $apiUrl ?? self::BASE_API_URL . "/" . $endpointKey;
    }

    /**
     * Validate payload.
     *
     * @param array $payload
     * @throws Exception if invalid
     */
    private function validatePayload(array $payload): void
    {
        if (!isset($payload["op"]) || !is_string($payload["op"])) {
            throw new Exception("INVALID_PAYLOAD: missing 'op'");
        }
    }

    /**
     * Dispatch an api call
     *
     * @param array $payload Request payload
     * @param array $headers Optional additional headers
     * @param string|null $bearerToken Optional bearer token
     * @return Result
     */
    public function dispatch(array $payload, array $headers = [], ?string $bearerToken = null): Result
    {
        try {
            $this->validatePayload($payload);

            $reqHeaders = array_merge(
                $this->headers,
                $headers,
                [
                    "x-api-key: " . $this->apiKey,
                    "x-sbc-sdk-client: singlebase-php",
                    "Content-Type: application/json",
                ]
            );

            if ($bearerToken) {
                $reqHeaders[] = "Authorization: Bearer $bearerToken";
            }

            $ch = curl_init($this->apiUrl);
            curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
            curl_setopt($ch, CURLOPT_HTTPHEADER, $reqHeaders);
            curl_setopt($ch, CURLOPT_POST, true);
            curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));

            $response = curl_exec($ch);
            $status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
            if ($response === false) {
                throw new Exception(curl_error($ch));
            }
            curl_close($ch);

            $parsed = json_decode($response, true);

            if ($status >= 200 && $status < 300) {
                return new ResultOK($parsed["data"] ?? [], $parsed["meta"] ?? [], $status);
            } else {
                return new ResultError($parsed["error"] ?? "Unknown Error", $status);
            }
        } catch (Exception $e) {
            return new ResultError("EXCEPTION: " . $e->getMessage(), 500);
        }
    }
}
