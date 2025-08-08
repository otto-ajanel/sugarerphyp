<?php

declare(strict_types=1);

namespace App\Middleware;

use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Server\MiddlewareInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class CorsMiddleware implements MiddlewareInterface
{
     public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        // 1. Manejar solicitudes OPTIONS (Preflight CORS)
        if ($request->getMethod() === 'OPTIONS') {
            return $this->handlePreflight($request);
        }

        // 2. Procesar solicitudes normales
        $response = $handler->handle($request);

        // 3. Añadir headers CORS a la respuesta
        return $this->addCorsHeaders($request, $response);
    }

    private function handlePreflight(ServerRequestInterface $request): ResponseInterface
    {
        $res =new \Hyperf\HttpMessage\Server\Response();
        return $res
            ->withHeader('Access-Control-Allow-Origin', '*')
            ->withHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS')
            ->withHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization, X-Requested-With')
            ->withHeader('Access-Control-Max-Age', '86400')
            ->withStatus(204);
    }

    private function addCorsHeaders(ServerRequestInterface $request, ResponseInterface $response): ResponseInterface
    {
        return $response
            ->withHeader('Access-Control-Allow-Origin', '*')
            ->withHeader('Access-Control-Expose-Headers', 'Content-Disposition');
    }
}

