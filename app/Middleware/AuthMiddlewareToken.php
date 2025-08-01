<?php

declare(strict_types=1);

namespace App\Middleware;

use App\Application\Services\Authservice;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface as HttpResponse;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\MiddlewareInterface;
use Psr\Http\Server\RequestHandlerInterface;

class AuthMiddlewareToken implements MiddlewareInterface
{
   /**
     * @var ContainerInterface
     */
    protected $container;

    /**
     * @var RequestInterface
     */
    protected $request;

    /**
     * @var HttpResponse
     */
    protected $response;

    private Authservice $authService;
    public function __construct(ContainerInterface $container, HttpResponse $response, RequestInterface $request, Authservice $authService)
    {
        $this->container = $container;
        $this->response = $response;
        $this->request = $request;
        $this->authService =$authService;
    }

    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        // According to the specific business judgment logic, it is assumed that the token carried by the user is valid here.

        $headerToke = $this->request->getHeaderLine('authorization');

        $isValidToken = $this->authService->valideToken(str_replace('Bearer ', '',$headerToke));

        
        if ($isValidToken) {
            return $handler->handle($request);
        }

        return $this->response->json(
            [
                'code' => -1,
                'data' => [
                    'error' => 'The token is invalid, preventing further execution.',
                ],
            ]
        )->withStatus(409);
    }
   
}
 