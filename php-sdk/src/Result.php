<?php

namespace Singlebase;

/**
 * Represents the result of an API operation.
 */
class Result
{
    /** @var array Response data */
    public array $data;

    /** @var array Metadata from API */
    public array $meta;

    /** @var bool Whether operation succeeded */
    public bool $ok;

    /** @var string|null Error message if failed */
    public ?string $error;

    /** @var int HTTP status code */
    public int $statusCode;

    /**
     * Construct a Result.
     *
     * @param array $data Response data
     * @param array $meta Metadata
     * @param bool $ok Whether operation succeeded
     * @param string|null $error Error message
     * @param int $statusCode HTTP status code
     */
    public function __construct(
        array $data = [],
        array $meta = [],
        bool $ok = true,
        ?string $error = null,
        int $statusCode = 200
    ) {
        $this->data = $data;
        $this->meta = $meta;
        $this->ok = $ok;
        $this->error = $error;
        $this->statusCode = $statusCode;
    }

    /**
     * Convert to associative array.
     *
     * @return array
     */
    public function toArray(): array
    {
        return [
            "data" => $this->data,
            "meta" => $this->meta,
            "ok" => $this->ok,
            "error" => $this->error,
            "statusCode" => $this->statusCode,
        ];
    }

    public function __toString(): string
    {
        return sprintf(
            "<Result ok=%s status=%d error=%s>",
            $this->ok ? "true" : "false",
            $this->statusCode,
            $this->error ?? "null"
        );
    }
}

/**
 * Represents a successful API result.
 */
class ResultOK extends Result
{
    public function __construct(array $data = [], array $meta = [], int $statusCode = 200)
    {
        parent::__construct($data, $meta, true, null, $statusCode);
    }
}

/**
 * Represents a failed API result.
 */
class ResultError extends Result
{
    public function __construct(?string $error, int $statusCode = 400)
    {
        parent::__construct([], [], false, $error, $statusCode);
    }
}
