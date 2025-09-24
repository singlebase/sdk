<?php

require 'vendor/autoload.php';

use Singlebase\Client;
use Singlebase\Upload;

$client = new Client(apiKey: "my-api-key", endpointKey: "vector-db");
$result = $client->dispatch([ "op" => "ping" ]);

if ($result->ok) {
    echo "✅ Success: ";
    print_r($result->data);
} else {
    echo "❌ Error: " . $result->error;
}

// Upload a file
Upload::uploadPresignedFile("myfile.txt", [
    "url" => "https://s3.amazonaws.com/bucket",
    "fields" => [ "key" => "uploads/myfile.txt" ]
]);
