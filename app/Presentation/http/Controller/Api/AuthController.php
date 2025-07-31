<?php
declare(strict_types=1);

namespace App\Presentation\http\Controller\Api;

use App\Application\Services\UserService;
use Hyperf\DbConnection\Db;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface;
use Psr\Http\Message\ResponseInterface as PsrRes;

class AuthController
{

    private UserService $userService;
    private RequestInterface $req;
    private ResponseInterface $res;

    public function __construct(
        UserService $userService,
        RequestInterface $request,
        ResponseInterface $response
    ) {
        $this->userService = $userService;
        $this->req = $request;
        $this->res= $response;
    }

    public function login($email, $password): PsrRes
    {
        $data = Db::connection('pgsql')->select('select * from users');

        return $this->res->json($data);
    }
}
