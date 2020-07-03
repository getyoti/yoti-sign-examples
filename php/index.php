<?php

require __DIR__ . '/vendor/autoload.php';

use GuzzleHttp\Client;
use GuzzleHttp\Psr7\MultipartStream;
use GuzzleHttp\Psr7\Request;
use Symfony\Component\Dotenv\Dotenv;

use function GuzzleHttp\Psr7\stream_for;

$dotenv = new Dotenv();
$dotenv->load(__DIR__.'/.env');

$options = (object) [
    'file_name' => 'test.pdf',
    'name' => 'Sign your request for..',
    'emails' => (object) [
        'invitation' => (object) [
            'body' => (object) [
                'message' => 'Please sign the example document',
            ],
        ],
    ],
    'recipients' => [
        (object) [
            'name' => 'John Smith',
            'email' => 'example@example.com',
            'role' => 'Signee',
            'auth_type' => 'no-auth',
            'sign_group' => 1,
            'tags' => [
                (object) [
                    'page_number' => 1,
                    'x' => 0.3,
                    'y' => 0.4,
                    'type' => 'signature',
                    'optional' => false,
                    'file_name' => 'test.pdf',
                ],
            ],
        ],
    ]
];

$request = new Request(
    'POST',
    $_ENV['YOTI_SIGN_BASE_URL'] . '/envelopes',
    [
        'Authorization' => 'Bearer ' . $_ENV['YOTI_AUTHENTICATION_TOKEN'],
    ],
    new MultipartStream([
        [
            'name' => 'file',
            'contents' => stream_for(fopen('./test.pdf', 'r')),
        ],
        [
            'name' => 'options',
            'contents' => stream_for(json_encode($options)),
        ],
    ])
);

$client = new Client();
$response = $client->send($request);

$json = json_decode($response->getBody());
$envelope_id = $json->envelope_id;

echo json_encode($json, JSON_PRETTY_PRINT) . "\n";
