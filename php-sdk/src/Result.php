<?php

namespace Singlebase;

class Result {
    public array $data;
    public array $meta;
    public bool $ok;
    public ?string $error;
    public int $statusCode;

    public function __construct(array $props = []) {
        $this->data = $props['data'] ?? [];
        $this->meta = $props['meta'] ?? [];
        $this->ok = $props['ok'] ?? true;
        $this->error = $props['error'] ?? null;
        $this->statusCode = $props['statusCode'] ?? 200;
    }

    public function toArray(): array {
        return [
            "data" => $this->data,
            "meta" => $this->meta,
            "ok" => $this->ok,
            "error" => $this->error,
            "statusCode" => $this->statusCode
        ];
    }

    public function __toString(): string {
        return "<Result ok=" . ($this->ok ? "true" : "false") .
               " status=" . $this->statusCode .
               " error=" . ($this->error ?? "null") . ">";
    }

    /**
     * Retrieve value from data using dot notation.
     *
     * @param string|null $path Dot-notation path, e.g. "address.city.zip"
     * @param mixed $default Default value if path not found
     * @return mixed
     * @throws \TypeError if traversal encounters a non-array type
     */
    public function getData(?string $path = null, $default = null) {
        if ($path === null || $path === "") {
            return $this->data;
        }

        $current = $this->data;
        foreach (explode('.', $path) as $part) {
            if (!is_array($current)) {
                throw new \TypeError("Cannot traverse '$part' — expected array, got " . gettype($current));
            }
            if (!array_key_exists($part, $current)) {
                return $default;
            }
            $current = $current[$part];
        }
        return $current;
    }
}

class ResultOK extends Result {
    public function __construct(array $props = []) {
        parent::__construct(array_merge($props, ["ok" => true]));
    }
}

class ResultError extends Result {
    public function __construct(array $props = []) {
        parent::__construct(array_merge($props, ["ok" => false]));
    }
}
